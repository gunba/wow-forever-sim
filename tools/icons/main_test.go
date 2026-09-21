package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDatabaseIconsAreBundled(t *testing.T) {
	type row struct {
		ID   int32  `json:"id"`
		Icon string `json:"icon"`
	}
	var db struct {
		Items, ItemIcons, SpellIcons []row
	}
	data, err := os.ReadFile("../../assets/database/db.json")
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(data, &db); err != nil {
		t.Fatal(err)
	}
	for kind, rows := range map[string][]row{"items": db.Items, "itemIcons": db.ItemIcons, "spellIcons": db.SpellIcons} {
		for _, item := range rows {
			path := filepath.Join("../../assets/img/wowhead/icons/large", strings.ToLower(item.Icon)+".jpg")
			if _, err := os.Stat(path); err != nil || item.Icon == "" {
				t.Errorf("%s %d: missing bundled icon %q; refresh with go run ./tools/icons", kind, item.ID, item.Icon)
			}
		}
	}
}
