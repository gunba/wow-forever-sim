package sim

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// The Go sim carries an integer for every ability and nothing else. Every name, icon and
// hover tooltip the site shows is resolved in the browser from Wowhead's Classic data for
// that integer, so an ability Forever changed is described by its Classic ancestor unless
// something says otherwise. That is the same trap the talent trees fell into, where 42
// talents wore another talent's tooltip and Improved Revenge kept Classic's stun long
// after Forever turned it into damage.
//
// ui/core/spells/<class>.json is that something: one entry per spell id the sim registers,
// saying where the numbers came from and, when Forever changed them, what the ability
// actually does. The UI renders those entries instead of the Wowhead tooltip. This test
// keeps the manifest and the sim in step - an ability cannot be registered without saying
// where its numbers came from, and the number still waiting on an answer can only fall.

// Raised by hand when an id is deliberately left undeclared, never by a tool. Every entry
// needs a reason, because an undeclared id is an ability the site describes wrongly.
const unreviewedSpellBudget = 44

// Registration sites whose id the walk cannot read from the source. Each one is an ability
// the manifest cannot cover, so this only ever falls.
const unresolvedSpellSiteBudget = 13

type spellSource struct {
	Ability     string   `json:"ability"`
	File        string   `json:"file"`
	Source      string   `json:"source"`
	ForeverID   int      `json:"foreverId,omitempty"`
	Tooltip     string   `json:"tooltip,omitempty"`
	Note        string   `json:"note,omitempty"`
	Assumptions []string `json:"assumptions,omitempty"`
}

var validSpellSources = map[string]bool{
	// Forever did not change the ability, so Wowhead's Classic tooltip describes it.
	"classic": true,
	// Forever changed it and a published Forever source gave the numbers.
	"forever": true,
	// Forever changed it or it is new, and at least one number is a guess.
	"assumed": true,
	// Not classified yet. Held to unreviewedSpellBudget.
	"unreviewed": true,
}

// Every spell id the sim registers, by id, with the files that use it, plus the
// registration sites whose id could not be read from the source.
//
// Reading the syntax tree rather than grepping keeps ids in comments out of the count, but
// it also has to follow how the ids are actually written. Only a minority are a literal in
// the ActionID; the ranked abilities build theirs from a table, either a local one indexed
// by rank (`spellId := [7]int32{0, 19434, ...}[rank]`), a package level one indexed by
// level (`map[int32]int32{60: 24248, ...}[druid.Level]`), or a rank struct passed in
// (`ActionID{SpellID: ripRank.id}` against `var ripRanks = []RipRankInfo{{id: 1079, ...}}`).
// A walk that only reads literals silently misses every rank of Shred, Rip, Moonfire and
// Tiger's Fury, which is most of what Forever changed about the druid.
func registeredSpellIDs(t *testing.T) (map[int][]string, []string) {
	t.Helper()

	ids := map[int]map[string]bool{}
	var unresolved []string
	fset := token.NewFileSet()

	// One package at a time, because a rank table and the spell that reads it live in the
	// same package but rarely in the same file.
	packages := map[string][]string{}
	err := filepath.WalkDir("..", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			if !strings.HasPrefix(filepath.ToSlash(path), "../sim") && path != ".." {
				return fs.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		dir := filepath.Dir(path)
		packages[dir] = append(packages[dir], path)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, paths := range packages {
		files := map[string]*ast.File{}
		for _, path := range paths {
			file, err := parser.ParseFile(fset, path, nil, 0)
			if err != nil {
				t.Fatalf("%s: %v", path, err)
			}
			files[path] = file
		}

		// Package level tables, which is where the ranked abilities keep their ids.
		globals := map[string]ast.Expr{}
		for _, file := range files {
			for _, decl := range file.Decls {
				gen, ok := decl.(*ast.GenDecl)
				if !ok || (gen.Tok != token.VAR && gen.Tok != token.CONST) {
					continue
				}
				for _, spec := range gen.Specs {
					value, ok := spec.(*ast.ValueSpec)
					if !ok {
						continue
					}
					for i, name := range value.Names {
						if i < len(value.Values) {
							globals[name.Name] = value.Values[i]
						}
					}
				}
			}
		}

		// type name -> field name -> that field's own type, so `rank.judge.spellID` can be
		// followed one hop at a time instead of guessing which struct owns `spellID`.
		structFields := map[string]map[string]string{}
		fieldOrder := map[string][]string{}
		for _, file := range files {
			for _, decl := range file.Decls {
				gen, ok := decl.(*ast.GenDecl)
				if !ok || gen.Tok != token.TYPE {
					continue
				}
				for _, spec := range gen.Specs {
					typeSpec, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}
					structType, ok := typeSpec.Type.(*ast.StructType)
					if !ok || structType.Fields == nil {
						continue
					}
					for _, field := range structType.Fields.List {
						for _, name := range field.Names {
							if structFields[typeSpec.Name.Name] == nil {
								structFields[typeSpec.Name.Name] = map[string]string{}
							}
							structFields[typeSpec.Name.Name][name.Name] = typeName(field.Type)
							fieldOrder[typeSpec.Name.Name] = append(fieldOrder[typeSpec.Name.Name], name.Name)
						}
					}
				}
			}
		}

		// type name -> field name -> the ids that field is ever given in this package. A
		// table of anonymous structs has no type name to file under, so it is filed under
		// the name of the variable holding it, which is what the loop over it reads.
		fields := map[string]map[string][]int{}
		tableNames := map[*ast.CompositeLit]string{}
		for _, file := range files {
			ast.Inspect(file, func(node ast.Node) bool {
				switch stmt := node.(type) {
				case *ast.AssignStmt:
					for i, target := range stmt.Lhs {
						name, ok := target.(*ast.Ident)
						if !ok || i >= len(stmt.Rhs) {
							continue
						}
						if lit := boundComposite(stmt.Rhs[i]); lit != nil {
							tableNames[lit] = name.Name
						}
					}
				case *ast.ValueSpec:
					for i, name := range stmt.Names {
						if i < len(stmt.Values) {
							if lit := boundComposite(stmt.Values[i]); lit != nil {
								tableNames[lit] = name.Name
							}
						}
					}
				}
				return true
			})
			ast.Inspect(file, func(node ast.Node) bool {
				lit, ok := node.(*ast.CompositeLit)
				if !ok {
					return true
				}
				name := typeName(lit.Type)
				if name == "" {
					name = tableNames[lit]
				}
				if name == "" {
					return true
				}
				order := fieldOrder[name]
				if array, ok := lit.Type.(*ast.ArrayType); ok {
					if element, ok := array.Elt.(*ast.StructType); ok {
						for _, field := range element.Fields.List {
							for _, member := range field.Names {
								order = append(order, member.Name)
							}
						}
					}
				}
				// A table of rank structs elides the element type, so the fields sit one
				// level below the only literal that still names the type.
				elements := []*ast.CompositeLit{lit}
				for _, elt := range lit.Elts {
					// A table keyed by level writes its ranks as `60: {spellID: ...}`, so the
					// rank sits behind the key rather than directly in the element list.
					if kv, ok := elt.(*ast.KeyValueExpr); ok {
						elt = kv.Value
					}
					if inner, ok := elt.(*ast.CompositeLit); ok && inner.Type == nil {
						elements = append(elements, inner)
					}
				}
				for _, element := range elements {
					for i, elt := range element.Elts {
						var key string
						value := elt
						if kv, ok := elt.(*ast.KeyValueExpr); ok {
							if member, ok := kv.Key.(*ast.Ident); ok {
								key, value = member.Name, kv.Value
							}
						} else if i < len(order) {
							key = order[i]
						}
						if id, ok := intLiteral(value); ok && key != "" {
							if fields[name] == nil {
								fields[name] = map[string][]int{}
							}
							fields[name][key] = append(fields[name][key], id)
						}
					}
				}
				return true
			})
		}

		// Literal arguments to package-local helpers also name real actions.
		callArguments := map[string]map[int][]ast.Expr{}
		for _, file := range files {
			ast.Inspect(file, func(node ast.Node) bool {
				call, ok := node.(*ast.CallExpr)
				if !ok {
					return true
				}
				fn, ok := call.Fun.(*ast.Ident)
				if !ok {
					return true
				}
				for i, arg := range call.Args {
					if _, ok := intLiteral(arg); ok {
						if callArguments[fn.Name] == nil {
							callArguments[fn.Name] = map[int][]ast.Expr{}
						}
						callArguments[fn.Name][i] = append(callArguments[fn.Name][i], arg)
					}
				}
				return true
			})
		}
		for path, file := range files {
			rel := strings.TrimPrefix(filepath.ToSlash(path), "../")

			// Each function carries its own bindings, so gather them before reading its
			// ActionIDs; anything declared at package level is read with none.
			scopes := []struct {
				body   ast.Node
				locals map[string]ast.Expr
			}{}
			for _, decl := range file.Decls {
				if fn, ok := decl.(*ast.FuncDecl); ok && fn.Body != nil {
					locals := localBindings(fn)
					position := 0
					for _, param := range fn.Type.Params.List {
						for _, name := range param.Names {
							if args := callArguments[fn.Name.Name][position]; len(args) > 0 {
								locals[name.Name] = &ast.CallExpr{Fun: ast.NewIdent("__alternatives"), Args: args}
							}
							position++
						}
					}
					scopes = append(scopes, struct {
						body   ast.Node
						locals map[string]ast.Expr
					}{fn.Body, locals})
					continue
				}
				scopes = append(scopes, struct {
					body   ast.Node
					locals map[string]ast.Expr
				}{decl, map[string]ast.Expr{}})
			}

			for _, scope := range scopes {
				locals := scope.locals

				ast.Inspect(scope.body, func(inner ast.Node) bool {
					lit, ok := inner.(*ast.CompositeLit)
					if !ok || !isActionID(lit.Type) {
						return true
					}
					for _, elt := range lit.Elts {
						kv, ok := elt.(*ast.KeyValueExpr)
						if !ok {
							continue
						}
						key, ok := kv.Key.(*ast.Ident)
						if !ok || key.Name != "SpellID" {
							continue
						}

						found := resolveSpellID(kv.Value, locals, globals, fields, structFields, 0)
						if len(found) == 0 {
							unresolved = append(unresolved, fmt.Sprintf("%s:%d", rel, fset.Position(kv.Pos()).Line))
							continue
						}
						for _, id := range found {
							// A zero id is the empty action, which names nothing.
							if id == 0 {
								continue
							}
							if ids[id] == nil {
								ids[id] = map[string]bool{}
							}
							ids[id][rel] = true
						}
					}
					return true
				})
			}
		}
	}

	out := map[int][]string{}
	for id, files := range ids {
		for file := range files {
			out[id] = append(out[id], file)
		}
		sort.Strings(out[id])
	}
	sort.Strings(unresolved)
	return out, unresolved
}

// The ids a single `SpellID:` value can stand for, in the shapes the sim uses. A name can
// stand for another name, so this follows bindings rather than looking only one step back.
// The composite an expression binds, seen through the index that picks one rank out of it:
// `map[int32]rankInfo{...}[warrior.Level]` binds the whole table to the local, and the
// field lookup needs the table filed under that local's name.
func boundComposite(expr ast.Expr) *ast.CompositeLit {
	for {
		switch node := expr.(type) {
		case *ast.CompositeLit:
			return node
		case *ast.IndexExpr:
			expr = node.X
		case *ast.ParenExpr:
			expr = node.X
		default:
			return nil
		}
	}
}

func resolveSpellID(expr ast.Expr, locals, globals map[string]ast.Expr, fields map[string]map[string][]int, structFields map[string]map[string]string, depth int) []int {
	if depth > 4 {
		return nil
	}
	if id, ok := intLiteral(expr); ok {
		return []int{id}
	}
	if call, ok := expr.(*ast.CallExpr); ok && typeName(call.Fun) == "__alternatives" {
		var found []int
		for _, arg := range call.Args {
			found = append(found, resolveSpellID(arg, locals, globals, fields, structFields, depth+1)...)
		}
		return found
	}

	// A table written in place, either indexed here or handed straight over.
	if found := intsInComposites(expr); len(found) > 0 {
		return found
	}

	// The two arms of a phase switch are both real ids.
	if call, ok := expr.(*ast.CallExpr); ok && strings.Contains(typeName(call.Fun), "Ternary") {
		var found []int
		for _, arg := range call.Args {
			if id, ok := intLiteral(arg); ok {
				found = append(found, id)
			}
		}
		if len(found) > 0 {
			return found
		}
	}

	// A field of a rank struct, matched on the struct's type name so that an unrelated
	// `id` field elsewhere in the package cannot answer for it. The qualifier can itself be
	// a field, as in `rank.judge.spellID`, so the chain is walked from the root down.
	if selector, ok := expr.(*ast.SelectorExpr); ok {
		var chain []string
		for node := ast.Expr(selector); ; {
			sel, ok := node.(*ast.SelectorExpr)
			if !ok {
				break
			}
			chain = append([]string{sel.Sel.Name}, chain...)
			node = sel.X
		}
		if root := rootIdent(selector.X); root != nil && len(chain) > 0 {
			field := chain[len(chain)-1]

			current := localTypeName(locals, root.Name)
			for _, hop := range chain[:len(chain)-1] {
				current = structFields[current][hop]
			}
			if found, ok := fields[current][field]; ok {
				return found
			}

			// A table declared inside the function has no type this can follow, so fall
			// back to the names in the expression itself: the field's own qualifier, then
			// the table the root was ranged over.
			candidates := []string{root.Name}
			if len(chain) > 1 {
				candidates = append([]string{chain[len(chain)-2]}, candidates...)
			}
			for _, bound := range []ast.Expr{locals[root.Name], globals[root.Name]} {
				if ident := rootIdent(bound); ident != nil {
					candidates = append(candidates, ident.Name)
				}
			}
			for _, candidate := range candidates {
				if found, ok := fields[candidate][field]; ok {
					return found
				}
			}
		}
		// A failed field lookup must not read every numeric field in its table.
		return nil
	}

	// A name bound to such a table, in the function or at package level.
	if root := rootIdent(expr); root != nil {
		for _, scope := range []map[string]ast.Expr{locals, globals} {
			bound, ok := scope[root.Name]
			if !ok || bound == expr {
				continue
			}
			if found := resolveSpellID(bound, locals, globals, fields, structFields, depth+1); len(found) > 0 {
				return found
			}
		}
	}
	return nil
}

// Int literals inside any composite the expression carries, restricted to the int typed
// ones so that a damage or mana table sitting beside the id table is not read as ids.
func intsInComposites(expr ast.Expr) []int {
	var found []int
	if lit := boundComposite(expr); lit != nil && isIntTyped(lit.Type) {
		for _, elt := range lit.Elts {
			value := elt
			if kv, ok := elt.(*ast.KeyValueExpr); ok {
				value = kv.Value
			}
			if id, ok := intLiteral(value); ok {
				found = append(found, id)
			}
		}
	}
	return found
}

func intLiteral(expr ast.Expr) (int, bool) {
	switch node := expr.(type) {
	case *ast.ParenExpr:
		return intLiteral(node.X)
	case *ast.CallExpr:
		// `int32(20662)` is as much a literal as `20662` for this purpose.
		if len(node.Args) == 1 && isIntName(typeName(node.Fun)) {
			return intLiteral(node.Args[0])
		}
	case *ast.BasicLit:
		if node.Kind == token.INT {
			id, err := strconv.Atoi(node.Value)
			return id, err == nil
		}
	}
	return 0, false
}

func rootIdent(expr ast.Expr) *ast.Ident {
	for {
		switch node := expr.(type) {
		case *ast.Ident:
			return node
		case *ast.IndexExpr:
			expr = node.X
		case *ast.SelectorExpr:
			expr = node.X
		case *ast.CallExpr:
			expr = node.Fun
		case *ast.ParenExpr:
			expr = node.X
		default:
			return nil
		}
	}
}

// Every name the function binds, whether by parameter or by assignment, so that a
// `SpellID:` naming one of them can be followed back to what it holds.
func localBindings(scope *ast.FuncDecl) map[string]ast.Expr {
	locals := map[string]ast.Expr{}

	if scope.Recv != nil {
		bindNames(locals, scope.Recv.List)
	}
	if scope.Type.Params != nil {
		bindNames(locals, scope.Type.Params.List)
	}

	ast.Inspect(scope.Body, func(node ast.Node) bool {
		switch stmt := node.(type) {
		case *ast.AssignStmt:
			if stmt.Tok != token.DEFINE && stmt.Tok != token.ASSIGN {
				return true
			}
			for i, target := range stmt.Lhs {
				name, ok := target.(*ast.Ident)
				if !ok || i >= len(stmt.Rhs) {
					continue
				}
				value := stmt.Rhs[i]
				if previous, ok := locals[name.Name]; ok {
					if alternatives, ok := previous.(*ast.CallExpr); ok && typeName(alternatives.Fun) == "__alternatives" {
						value = &ast.CallExpr{Fun: ast.NewIdent("__alternatives"), Args: append(alternatives.Args, value)}
					} else {
						value = &ast.CallExpr{Fun: ast.NewIdent("__alternatives"), Args: []ast.Expr{previous, value}}
					}
				}
				locals[name.Name] = value
			}
		case *ast.RangeStmt:
			// The loop variable stands for an element, and an element of a table of ids is
			// an id, so binding it to the table itself resolves to the right set.
			if name, ok := stmt.Value.(*ast.Ident); ok {
				locals[name.Name] = stmt.X
			}
		}
		return true
	})
	return locals
}

// A parameter holds its declared type rather than a value, which is what the rank struct
// lookup needs; the type is stashed under the name so both live in one map.
func bindNames(locals map[string]ast.Expr, list []*ast.Field) {
	for _, field := range list {
		for _, name := range field.Names {
			locals[name.Name] = &ast.CompositeLit{Type: field.Type}
		}
	}
}

func localTypeName(locals map[string]ast.Expr, name string) string {
	bound, ok := locals[name]
	if !ok {
		return ""
	}
	lit, ok := bound.(*ast.CompositeLit)
	if !ok {
		return ""
	}
	return typeName(lit.Type)
}

func typeName(expr ast.Expr) string {
	switch node := expr.(type) {
	case *ast.Ident:
		return node.Name
	case *ast.StarExpr:
		return typeName(node.X)
	case *ast.ArrayType:
		return typeName(node.Elt)
	case *ast.MapType:
		return typeName(node.Value)
	case *ast.SelectorExpr:
		return node.Sel.Name
	}
	return ""
}

func isIntTyped(expr ast.Expr) bool {
	switch node := expr.(type) {
	case *ast.ArrayType:
		return isIntName(typeName(node.Elt))
	case *ast.MapType:
		return isIntName(typeName(node.Value))
	}
	return false
}

func isIntName(name string) bool {
	return name == "int32" || name == "int" || name == "int64"
}

func isActionID(expr ast.Expr) bool {
	switch typ := expr.(type) {
	case *ast.Ident:
		return typ.Name == "ActionID"
	case *ast.SelectorExpr:
		return typ.Sel.Name == "ActionID"
	}
	return false
}

func loadSpellSources(t *testing.T) map[int]spellSource {
	t.Helper()

	dir := filepath.Join("..", "ui", "core", "spells")
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	sources := map[int]spellSource{}
	owner := map[int]string{}
	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			t.Fatal(err)
		}
		var file map[string]spellSource
		if err := json.Unmarshal(data, &file); err != nil {
			t.Fatalf("%s: %v", entry.Name(), err)
		}
		for key, source := range file {
			id, err := strconv.Atoi(key)
			if err != nil {
				t.Errorf("%s: %q is not a spell id", entry.Name(), key)
				continue
			}
			if previous, ok := owner[id]; ok {
				t.Errorf("spell %d is declared in both %s and %s", id, previous, entry.Name())
				continue
			}
			owner[id] = entry.Name()
			sources[id] = source
		}
	}
	return sources
}

func TestSpellSourceScannerKeepsBranchesAndRejectsRankScalars(t *testing.T) {
	ids, _ := registeredSpellIDs(t)
	for _, id := range []int{401502, 1237312, 1237313, 1277324, 1277328, 1259812, 1259813, 1259817, 1259821, 1259823, 1293697, 1293698, 7217, 22841} {
		if len(ids[id]) == 0 {
			t.Errorf("missed rank, conditional or helper ID %d", id)
		}
	}
	for id, files := range ids {
		for _, path := range files {
			if path == "sim/mage/frostfire_bolt.go" && id != 401502 && id != 1237312 && id != 1237313 {
				t.Errorf("Frostfire scalar %d was mistaken for a spell", id)
			}
			if path == "sim/priest/dark_sacrifice.go" && (id < 1277324 || id > 1277328) {
				t.Errorf("Dark Sacrifice scalar %d was mistaken for a spell", id)
			}
		}
	}
}

func TestEveryRegisteredSpellSaysWhereItsNumbersCameFrom(t *testing.T) {
	registered, _ := registeredSpellIDs(t)
	sources := loadSpellSources(t)

	if len(registered) == 0 {
		t.Fatal("found no spell ids in sim/, the walk is broken rather than the manifest")
	}

	var undeclared []int
	for id := range registered {
		if _, ok := sources[id]; !ok {
			undeclared = append(undeclared, id)
		}
	}
	sort.Ints(undeclared)
	for _, id := range undeclared {
		t.Errorf("spell %d (%s) is not in ui/core/spells, so the site describes it with Wowhead's Classic entry",
			id, strings.Join(registered[id], ", "))
	}

	var orphaned []int
	for id := range sources {
		if _, ok := registered[id]; !ok {
			orphaned = append(orphaned, id)
		}
	}
	sort.Ints(orphaned)
	for _, id := range orphaned {
		t.Errorf("spell %d is declared in ui/core/spells but nothing in sim/ registers it", id)
	}
}

func TestSpellSourcesAreWellFormed(t *testing.T) {
	for id, source := range loadSpellSources(t) {
		if source.Ability == "" {
			t.Errorf("spell %d: no ability name", id)
		}
		if source.File == "" {
			t.Errorf("spell %d (%s): no file", id, source.Ability)
		}
		if !validSpellSources[source.Source] {
			t.Errorf("spell %d (%s): %q is not a source", id, source.Ability, source.Source)
		}

		// An ability Forever changed has to carry its own words, or the site falls back to
		// the Classic tooltip and the entry has bought nothing.
		if source.Source == "forever" || source.Source == "assumed" {
			if source.Tooltip == "" {
				t.Errorf("spell %d (%s): %s with no tooltip of its own", id, source.Ability, source.Source)
			}
		} else if source.Tooltip != "" {
			t.Errorf("spell %d (%s): %s does not need a tooltip, Wowhead's is right", id, source.Ability, source.Source)
		}

		// A guess that does not say what was guessed cannot be checked against the beta.
		if source.Source == "assumed" && len(source.Assumptions) == 0 {
			t.Errorf("spell %d (%s): assumed but lists nothing to confirm", id, source.Ability)
		}
		if source.Source != "assumed" && len(source.Assumptions) > 0 {
			t.Errorf("spell %d (%s): %s does not assume anything", id, source.Ability, source.Source)
		}
	}
}

func TestUnreviewedSpellsOnlyShrink(t *testing.T) {
	var unreviewed []int
	for id, source := range loadSpellSources(t) {
		if source.Source == "unreviewed" {
			unreviewed = append(unreviewed, id)
		}
	}
	sort.Ints(unreviewed)

	if len(unreviewed) > unreviewedSpellBudget {
		t.Errorf("%d spells are unreviewed and the budget is %d: %v", len(unreviewed), unreviewedSpellBudget, unreviewed)
	}
	if len(unreviewed) < unreviewedSpellBudget {
		t.Errorf("only %d spells are unreviewed, lower unreviewedSpellBudget to %d", len(unreviewed), len(unreviewed))
	}
}

// A site the walk cannot read is worse than a missing entry, because nothing downstream
// knows the ability exists. Better to name them than to report coverage that is not there.
func TestEverySpellRegistrationResolvesToAnID(t *testing.T) {
	_, unresolved := registeredSpellIDs(t)

	if len(unresolved) > unresolvedSpellSiteBudget {
		for _, site := range unresolved {
			t.Errorf("%s: the spell id here cannot be read from the source, so the manifest cannot cover it", site)
		}
		t.Errorf("%d unreadable registration sites, the budget is %d", len(unresolved), unresolvedSpellSiteBudget)
	}
	if len(unresolved) < unresolvedSpellSiteBudget {
		t.Errorf("only %d sites are unreadable, lower unresolvedSpellSiteBudget to %d", len(unresolved), len(unresolved))
	}
}
