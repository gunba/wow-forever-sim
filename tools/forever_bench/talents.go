package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
)

type location struct {
	Row int `json:"rowIdx"`
	Col int `json:"colIdx"`
}
type talent struct {
	Field        string    `json:"fieldName"`
	Location     location  `json:"location"`
	Max          int       `json:"maxPoints"`
	Prereq       *location `json:"prereqLocation"`
	NotSimulated bool      `json:"notSimulated"`
}
type tree struct {
	Name    string   `json:"name"`
	Talents []talent `json:"talents"`
}
type talentConfig struct {
	Trees []tree
	Build build
}

func loadTalents(b build) talentConfig {
	class := strings.ToLower(strings.TrimPrefix(b.Class.String(), "Class"))
	var trees []tree
	if err := json.Unmarshal(mustRead(filepath.Join("ui/core/talents/trees", class+".json")), &trees); err != nil {
		panic(err)
	}
	for ti := range trees {
		for i := range trees[ti].Talents {
			if trees[ti].Talents[i].Max == 0 {
				trees[ti].Talents[i].Max = 1
			}
		}
	}
	return talentConfig{trees, b}
}

func (c talentConfig) decode(s string) ([]int, error) {
	parts := strings.Split(s, "-")
	if len(parts) > 3 {
		return nil, fmt.Errorf("too many talent trees")
	}
	var points []int
	for ti, tr := range c.Trees {
		part := ""
		if ti < len(parts) {
			part = parts[ti]
		}
		if len(part) > len(tr.Talents) {
			return nil, fmt.Errorf("%s: too many ranks", tr.Name)
		}
		for i := range tr.Talents {
			p := 0
			if i < len(part) {
				if part[i] < '0' || part[i] > '9' {
					return nil, fmt.Errorf("invalid talent rank")
				}
				p = int(part[i] - '0')
			}
			points = append(points, p)
		}
	}
	return points, nil
}

func (c talentConfig) encode(points []int) string {
	var trees []string
	offset := 0
	for _, tr := range c.Trees {
		var part strings.Builder
		for i := range tr.Talents {
			part.WriteByte(byte(points[offset+i]) + '0')
		}
		trees = append(trees, strings.TrimRight(part.String(), "0"))
		offset += len(tr.Talents)
	}
	return strings.TrimRight(strings.Join(trees, "-"), "-")
}

func (c talentConfig) validate(s string) error {
	points, err := c.decode(s)
	if err != nil {
		return err
	}
	offset, total := 0, 0
	hasRequired := c.Build.Required == ""
	for ti, tr := range c.Trees {
		rows := make(map[int]int)
		byLocation := make(map[location]int)
		treeTotal := 0
		for i, t := range tr.Talents {
			p := points[offset+i]
			if p > t.Max {
				return fmt.Errorf("%s exceeds rank cap", t.Field)
			}
			rows[t.Location.Row] += p
			byLocation[t.Location] = p
			treeTotal += p
			if t.Field == c.Build.Required && p > 0 {
				hasRequired = true
			}
		}
		for i, t := range tr.Talents {
			if points[offset+i] == 0 {
				continue
			}
			above := 0
			for row := 0; row < t.Location.Row; row++ {
				above += rows[row]
			}
			if above < 5*t.Location.Row {
				return fmt.Errorf("%s needs %d points above its row; has %d", t.Field, 5*t.Location.Row, above)
			}
			if t.Prereq != nil {
				for _, parent := range tr.Talents {
					if parent.Location == *t.Prereq && byLocation[parent.Location] != parent.Max {
						return fmt.Errorf("%s needs %s", t.Field, parent.Field)
					}
				}
			}
		}
		if treeTotal < c.Build.MinimumTrees[ti] {
			return fmt.Errorf("%s does not meet build's tree identity", tr.Name)
		}
		total += treeTotal
		offset += len(tr.Talents)
	}
	if total != 51 {
		return fmt.Errorf("spent %d talent points, want 51", total)
	}
	if !hasRequired {
		return fmt.Errorf("missing build talent %s", c.Build.Required)
	}
	return nil
}

func (c talentConfig) neighbors(s string) []string {
	points, err := c.decode(s)
	if err != nil {
		panic(err)
	}
	var nodes []talent
	for _, tr := range c.Trees {
		nodes = append(nodes, tr.Talents...)
	}
	var candidates []string
	for from, p := range points {
		if p == 0 {
			continue
		}
		points[from]--
		for to, node := range nodes {
			if to == from || points[to] >= node.Max || node.NotSimulated {
				continue
			}
			points[to]++
			encoded := c.encode(points)
			if c.validate(encoded) == nil {
				candidates = append(candidates, encoded)
			}
			points[to]--
		}
		points[from]++
	}
	return candidates
}
