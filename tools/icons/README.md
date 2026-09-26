# Display metadata and icons

```sh
python3 tools/database/import_forever_spell_icons.py
go run ./tools/database/gen_db -outDir=assets -gen=db
go run ./tools/icons
make dist/classic/.dirstamp
```

The metadata importer uses Forever tooltips and the spell manifest's current-ID
mappings. Internal-only proc IDs use the matching client-derived talent icon;
this does not substitute another spell's mechanics or tooltip.

`forever_spell_icons.json` is bundled into the database. The icon mirror copies
the corresponding images locally. Rebuild both the database and staged assets
after adding records; running Vite alone does not copy them.

The cached starter tooltips use the Classic parser; the generator then applies
the compiled Forever overlays. Passing `-wowhead=forever` to this cached-data
rebuild drops the starter random-suffix definitions.

The page smoke test checks local image responses while blocking third-party
tooltip hosts.
