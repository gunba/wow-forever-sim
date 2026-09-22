package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/tools/database"
)

func TestForeverCatalogReplacesLegacyAvailability(t *testing.T) {
	path := filepath.Join(t.TempDir(), "catalog.json")
	if err := os.WriteFile(path, []byte(`{"items":[{"id":1,"name":"Reviewed item","icon":"inv_misc_questionmark"}]}`), 0600); err != nil {
		t.Fatal(err)
	}
	db := database.NewWowDatabase()
	db.Items[2] = &proto.UIItem{Id: 2, Name: "Provisional legacy item"}
	replaceForeverEquipment(db, path)
	if len(db.Items) != 1 || db.Items[1] == nil || db.Items[2] != nil {
		t.Fatal("source-filtered catalog retained legacy availability")
	}
}
