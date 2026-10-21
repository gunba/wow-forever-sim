package env

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// The class's talent trees from ui/core/talents/trees, flattened in talent-string order.
type talentNode struct {
	Field  string
	Name   string
	Tree   int
	Row    int
	Col    int
	Max    int
	Prereq int // flat index of the talent that must be full first, or -1
}

type talentTrees struct {
	Names []string
	Nodes []talentNode
	Sizes []int
}

func loadTalentTrees(repo, class string) (*talentTrees, error) {
	data, err := os.ReadFile(filepath.Join(repo, "ui/core/talents/trees", class+".json"))
	if err != nil {
		return nil, err
	}
	var raw []struct {
		Name    string `json:"name"`
		Talents []struct {
			Field    string `json:"fieldName"`
			Name     string `json:"name"`
			Max      int    `json:"maxPoints"`
			Location struct {
				Row int `json:"rowIdx"`
				Col int `json:"colIdx"`
			} `json:"location"`
			Prereq *struct {
				Row int `json:"rowIdx"`
				Col int `json:"colIdx"`
			} `json:"prereqLocation"`
		} `json:"talents"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	t := &talentTrees{}
	for ti, tr := range raw {
		t.Names = append(t.Names, tr.Name)
		t.Sizes = append(t.Sizes, len(tr.Talents))
		start := len(t.Nodes)
		for _, n := range tr.Talents {
			max := n.Max
			if max == 0 {
				max = 1
			}
			t.Nodes = append(t.Nodes, talentNode{Field: n.Field, Name: n.Name, Tree: ti, Row: n.Location.Row, Col: n.Location.Col, Max: max, Prereq: -1})
		}
		for i, n := range tr.Talents {
			if n.Prereq == nil {
				continue
			}
			for j, m := range tr.Talents {
				if m.Location.Row == n.Prereq.Row && m.Location.Col == n.Prereq.Col {
					t.Nodes[start+i].Prereq = start + j
				}
			}
		}
	}
	return t, nil
}

func (t *talentTrees) decode(s string) ([]int, error) {
	parts := strings.Split(s, "-")
	if len(parts) > len(t.Sizes) {
		return nil, fmt.Errorf("talent string %q has too many trees", s)
	}
	points := make([]int, len(t.Nodes))
	offset := 0
	for ti, size := range t.Sizes {
		if ti < len(parts) {
			if len(parts[ti]) > size {
				return nil, fmt.Errorf("talent string %q: tree %d too long", s, ti)
			}
			for i, ch := range parts[ti] {
				if ch < '0' || ch > '9' {
					return nil, fmt.Errorf("talent string %q: bad rank", s)
				}
				points[offset+i] = int(ch - '0')
			}
		}
		offset += size
	}
	return points, nil
}

func (t *talentTrees) encode(points []int) string {
	var trees []string
	offset := 0
	for _, size := range t.Sizes {
		var b strings.Builder
		for i := 0; i < size; i++ {
			b.WriteByte(byte('0' + points[offset+i]))
		}
		trees = append(trees, strings.TrimRight(b.String(), "0"))
		offset += size
	}
	return strings.TrimRight(strings.Join(trees, "-"), "-")
}

func (t *talentTrees) treeSpent(points []int, tree int) int {
	n := 0
	for i, node := range t.Nodes {
		if node.Tree == tree {
			n += points[i]
		}
	}
	return n
}

// canAdd reports whether one more point can go into talent i: under its cap, 5 points in the tree per row above it,
// and its prerequisite full. Points are only ever added, so a build made this way is legal.
func (t *talentTrees) canAdd(points []int, i int) bool {
	node := t.Nodes[i]
	if points[i] >= node.Max {
		return false
	}
	if t.treeSpent(points, node.Tree) < 5*node.Row {
		return false
	}
	if node.Prereq >= 0 && points[node.Prereq] < t.Nodes[node.Prereq].Max {
		return false
	}
	return true
}

// teacherPoint is the next point of the target build that can be added now (lowest row first), or -1.
func (t *talentTrees) teacherPoint(points, target []int) int {
	best := -1
	for i, node := range t.Nodes {
		if points[i] < target[i] && t.canAdd(points, i) {
			if best < 0 || node.Row < t.Nodes[best].Row {
				best = i
			}
		}
	}
	return best
}
