"""Run the publication checks shared by local builds and GitHub Pages."""

import os
from functools import partial
from http.server import SimpleHTTPRequestHandler, ThreadingHTTPServer
from pathlib import Path
import shlex
import subprocess
import tempfile
import threading

ROOT = Path(__file__).resolve().parents[1]


def run(*args, env=None):
    print("+ " + shlex.join(map(str, args)), flush=True)
    subprocess.run(args, cwd=ROOT, env=env, check=True)


class SiteHandler(SimpleHTTPRequestHandler):
    def log_message(self, format, *args):
        pass


def main():
    for command in (
        "go test -tags with_db ./tools/forever_bench ./sim/core ./sim/druid ./sim/warlock ./tools/database ./tools/database/gen_db ./tools/icons",
        "go test -tags with_db ./tools/forever_tanks ./sim/warrior/dps_warrior ./sim/warrior/tank_warrior ./sim/paladin/protection ./sim/druid/tank",
        "go test -tags with_db ./sim -run 'Test(EveryRegisteredSpellSaysWhereItsNumbersCameFrom|EverySpellRegistrationResolvesToAnID|SpellSourcesAreWellFormed|UnreviewedSpellsOnlyShrink|SpellSourceScannerKeepsBranchesAndRejectsRankScalars|SpellSourceScannerFollowsTypedRangeParameter)$'",
        "python3 -m unittest discover -s tools/forever_bench",
        "python3 tools/forever_bench/uncertainties.py --check",
        "python3 -m unittest discover -s tools/database",
        "python3 -m unittest discover -s tools/data_watch",
    ):
        run(*shlex.split(command))

    with tempfile.TemporaryDirectory(prefix="forever-verify-") as temporary:
        work = Path(temporary)
        scrub = work / "scrub-test.cjs"
        run("node_modules/.bin/esbuild", "ui/scrub/scrub.test.ts",
            "--bundle", "--platform=node", f"--outfile={scrub}")
        run("node", scrub)
        run("python3", "tools/forever_bench/build_review_site.py",
            "--results", "artifacts/modelled_gear/forever_dps_5min.json",
            "--profiles", "artifacts/modelled_gear/ui_profiles",
            "--sensitivity", "artifacts/modelled_gear/forever_sensitivity.json",
            "--gear-search", "artifacts/modelled_gear_search/current")

        base = "/" + os.environ.get("SITE_BASE", "/classic/").strip("/")
        base = base.rstrip("/") + "/"
        site = work / "site"
        mount = site / base.strip("/")
        mount.parent.mkdir(parents=True, exist_ok=True)
        mount.symlink_to(ROOT / "dist/classic", target_is_directory=True)
        handler = partial(SiteHandler, directory=str(site))
        with ThreadingHTTPServer(("127.0.0.1", 0), handler) as server:
            thread = threading.Thread(target=server.serve_forever, daemon=True)
            thread.start()
            env = {**os.environ, "SITE_URL": f"http://127.0.0.1:{server.server_port}{base}"}
            try:
                for check in (
                    "check_pages", "check_questions", "check_tank_matrix",
                    "check_scrub", "check_icons", "check_benchmark_replay",
                    "check_ranked_defaults",
                ):
                    run("node", f"tools/smoke/{check}.mjs", env=env)
                tank = work / "forever-tank-bench"
                run("go", "build", "-tags", "with_db", "-o", tank, "./tools/forever_tanks")
                run("node", "tools/smoke/check_tank_defaults.mjs",
                    env={**env, "TANK_BENCH": str(tank)})
            finally:
                server.shutdown()
                thread.join()


if __name__ == "__main__":
    main()
