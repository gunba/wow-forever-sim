# WoW Forever simulator

## Local development

Choose the existing check for the changed behaviour:

- Go: `go test -tags with_db ./sim/<package> -run '<TestName>'` (or the affected package under `tools/`).
- Python: `python3 -m unittest discover -s tools/<area> -p '<test_file>.py'`.
- TypeScript: `npm run type-check`.
- Built UI: serve `dist/` locally, then run `SITE_URL=http://localhost:8080/classic/ node tools/smoke/check_pages.mjs --page <page>`. Repeat `--page` for multiple pages; `.` selects the landing page. `--list` prints the built pages.
- Other UI behaviour: use the relevant existing script under `tools/smoke/` against the same local server.

`make dist/classic/.dirstamp` rebuilds the site. Use the same `SITE_BASE` for building and serving. `CHROMIUM_PATH` can select an installed browser instead of downloading Playwright Chromium.

On the Linux workspace without native Go, use the cached Go container for focused Go checks:

```sh
docker run --rm -v "$PWD":/src -v "$HOME/.cache/go-build":/root/.cache/go-build -v "$HOME/go/pkg/mod":/go/pkg/mod -w /src golang:1.24-bookworm go test -tags with_db ./sim/<package> -run '<TestName>'
```

Native builds and `make verify` require the development toolchain described in `README.md`; a `go env GOROOT` shim is not a Go compiler.

After the relevant check passes, stop validation until the implementation changes. `npm test` invokes the upstream Bazel suite; it is not the default development check.

## Publication

`make verify` builds the site and runs the publication checks, including all browser smoke checks, using a temporary local server. It is the shared entry point for local release verification and `.github/workflows/deploy.yml`, not an every-edit check.

The maintained fork is remote `publish` (`gunba/wow-forever-sim`), branch `forever`. The `origin` remote is upstream. Push completed changes to `publish`; a push does not deploy.

When deployment is requested, dispatch `deploy.yml` on `forever` once the changes are ready. The workflow builds, verifies and publishes GitHub Pages. Reuse passing local checks rather than repeating them solely before pushing.
