# WoW: Forever sim

A fork of [ElliotWood/Forever](https://github.com/ElliotWood/Forever), built on
[wowsims/classic](https://github.com/wowsims/classic), for **World of Warcraft: Forever**.

**[Open the simulator and DPS matrix](https://gunba.github.io/wow-forever-sim/classic/review/).**

The current DPS profiles use Forever crafted, dungeon and verified PvP/vendor equipment, updated talent
trees and racials, and an explicit Tier 1 bonus setting. The Classic ruleset is
retained for mechanic comparisons, not as a complete Classic simulator.

The [five-minute benchmark](tools/forever_bench/README.md) covers 26 builds and
171 Horde/Alliance race/build combinations. Its [matrix](artifacts/forever_dps_5min.png),
[CSV](artifacts/forever_dps_5min.csv) and
[raw requests/results](artifacts/forever_dps_5min.json) use paid hit normalization
and 5,000 iterations per result. Legal talent and rotation changes were screened,
independently validated and checked across every available race; these are not
claims of a global optimum or best-in-slot equipment.

The matrix shows race-averaged Tier 1 gains, hypothetical +10% equipment gains
and scaling amplification: the mean +50% gain divided by five times the mean
+10% gain. A value of 1× is linear; larger values indicate accelerating gains.
This measures finite-range curvature, not proof of exponential growth.
[Sensitivity data](artifacts/forever_sensitivity.json) records both scaling runs,
the assumptions and Monte Carlo uncertainty.

[Build reviews](docs/build_reviews.md) describe the retained choices and resource
use. [In-game checks](docs/in_game_checks.md) track important unresolved mechanics.
[Mechanics review](docs/mechanics_review.md) records corrections and provisional models.
[Change-history review](docs/history_review.md) records the fork points,
reviewed changes, regression checks and remaining source gaps.
[Forever upstream review](docs/upstream-forever-review-2026-09-23.md) compares
recent WoWSims changes and Warrior combat logs with this simulator.
[Replay profiles](artifacts/ui_profiles/index.json) are built into the matching
simulator's **Ranked builds** selector. Clicking a matrix cell opens that exact
setup, including its paid hit adjustment; no manual import is needed.
[Frozen benchmark inputs](artifacts/forever_input_profiles.json) separate the
profile roster from the results for replay.
[Equipment-search evidence](artifacts/gear_search/summary.json) records the
comparisons. Lower-level items remain where stronger or needed for documented
coverage gaps.

This project is licensed with MIT license, inherited from the upstream project. As upstream requests, keep a user visible link back to [wowsims/classic](https://github.com/wowsims/classic) in anything built on this.

## What is different from Classic

`SimOptions.ruleset` picks the rules. `RulesetClassic` is the wire default;
the UI and benchmark use `RulesetForever`. Talent and item changes are not
reverted by the engine switch.

| Rule | What changed |
| --- | --- |
| Periodic critical strikes | Spell dots and bleeds roll for crits, against the caster's crit chance at the time of the tick |
| Unified hit and critical strike | Generic item ratings cover physical and spell attacks; traps still use an inherited special hit rule |
| Bonus healing on gear | Healing and damage use separate sourced effects; healing-only records do not grant spell damage |
| Racials | Resistance racials removed, weapon skill racials pay crit instead, several races reshaped |
| Races | The Skyborne, plus six new race and class pairings |

Talent trees are **not** switchable. All nine classes carry their Forever tree under both rulesets, because there is only one talent proto and the field numbers are positional. Running under `RulesetClassic` therefore gives you Classic engine rules with Forever talents, which is useful for isolating a rule change and is not a faithful Classic character.

## Where the talent data comes from

The trees started as a transcription of the BlizzCon 2026 stream (`tools/forever_talents/data/`, from
[Deradon/wow-forever-talent-calc](https://github.com/Deradon/wow-forever-talent-calc)), which only ever showed rank 1.
Since the beta client went out on 17 September 2026 **the beta client is the source of truth** for talent names,
positions, prerequisites, rank counts and per-rank values. Where the two disagree, the client wins, including over
the talentsforever.com crawl, which is also BlizzCon footage.

### How the beta data is read

No client install or extraction is needed. [wago.tools](https://wago.tools) publishes every DB2 table of every build
as CSV, so the whole pipeline is HTTP:

1. **Find the build.** Forever's beta ships under the `wow_classic_beta` product as version `1.60.x`
   (`https://wago.tools/api/builds`, sort by `created_at`; the list is not ordered). The first was `1.60.1.69893`.
2. **Download the tables.** `https://wago.tools/db2/<Table>/csv?build=<build>`. Talents live in the retail-style
   Trait tables: `TraitTree`, `TraitNode`, `TraitNodeEntry`, `TraitNodeXTraitNodeEntry`, `TraitDefinition`,
   `TraitDefinitionEffectPoints`, `TraitEdge`, plus `CurvePoint`, `SpellName`, `Spell` and `SpellEffect`.
   **Ignore `Talent` and `TalentTab`**: in the beta they are unchanged Classic Era leftovers.
3. **Rebuild the trees.** One `TraitTree` per class holds all three tabs side by side on one canvas: tabs start at
   `PosX` 1020 / 5020 / 9080, rows at `PosY` 2130 + 600 per row, columns 600 apart. A few nodes carry a stray extra
   zero in a position, and where two nodes share a spell the newer node id is the live one. Prerequisites are
   `TraitEdge` rows with `Type` other than 0. Tab names are not in the tables, so tabs are matched to the existing
   trees by talent-name overlap.
4. **Read per-rank values from curves, not tooltips.** `TraitDefinitionEffectPoints` gives each effect a `CurveID`,
   and `CurvePoint(rank)` is that effect's value at that rank (verified on Ignite 8/16/24/32/40, Ruin, Improved Life
   Tap). `SpellEffect.EffectBasePointsF` only holds one value and is sometimes stale.
5. **Resolve the tooltip text.** `$s1`/`$m1` is effect 1's value, `$/1000;s1` and `${$m1/1000}` apply a divisor
   (durations are stored in milliseconds, rage in tenths), and a trailing `.1` means one decimal place. Tokens that
   point at other spells, formulas or durations are left as `<d>`, `<o>`, `<other spell s>` and never copied into
   the trees as numbers.

In practice:

    tools/forever_talents/export_beta.py 1.60.1.69893 beta/     # trees in the data/<class>.json schema
    tools/forever_talents/diff_trees.py beta/                   # what changed against the sim
    tools/forever_talents/apply_beta_tooltips.py beta/          # tooltips, rank values, assets/confirmed_talents.json

Structural changes (renames, removals, moves) are applied by hand because they renumber the talent protos, and
`TestConfirmedTalentRanksMatchTheSim` and `TestGoTalentScalingMatchesTheTrees` then hold the trees and the Go to
the client's numbers.

**Watching for changes.** `.github/workflows/watch_wowhead_forever.yml` runs every 6 hours: it snapshots Wowhead's
Forever gear-planner data and diffs the newest `1.60.x` client build's DB2 tables
(`tools/data_watch/wago_db2_diff.py`), opening a `data-change` pull request when either moves.

The tree json and the proto message have to stay in the same order: `FillTalentsProto` maps the nth character of a
talent string to proto field number n. Changing the order invalidates saved talent strings.

## Known gaps

Worth knowing before reading any number out of this sim:

- **Beta mechanics are not all verified.** Warrior rage, parts of the combat
  table, some pet behavior and server-scripted proc interactions remain inherited
  or approximate. Druid Omen provisionally uses the client's 100% proc entry
  and ten-second internal cooldown; its effective server rate needs testing.
  See the benchmark's evidence notes.
- **Gear coverage is incomplete.** The current DPS loadouts are complete and use
  actual Forever stats, but the published ring, neck and trinket pools are sparse.
  Missing records are not filled with Classic equivalents. See
  [gear data](docs/forever_gear_data.md).
- **Tier 1 is a separate scenario setting.** It grants the role's full set
  bonuses without raid-item stats. Creature restrictions still apply. Utility
  effects and dummy-effect assumptions are listed in the
  [Tier 1 notes](docs/forever_tier1.md).
- **The benchmark is not a raid composition.** It uses a fixed external buff
  package without world buffs. It does not prove that a particular twenty-player
  roster supplies every buff.
- **Tank/healer profiles and other item effects have not received the same
  review as the current DPS baselines.**

## Running it

There are no published builds for this fork and the deploy workflow does not run here — build and host it locally with the instructions below. Upstream's releases and [live sims](https://wowsims.github.io/classic) are Classic Era and do not include any of this.

With the development dependencies installed:

```sh
npm ci
make proto
GOFLAGS=-buildvcs=false make dist/classic/.dirstamp
go build -buildvcs=false -o /tmp/forever-web ./sim/web/main.go
/tmp/forever-web --usefs=true --launch=false --host=127.0.0.1:3333
```

Open <http://127.0.0.1:3333/classic/>. The gear, talent and rotation tabs offer
matching build presets. **Settings → Other → Tier 1 bonuses** controls the
gear-independent set effects. Paid hit normalization belongs to the benchmark;
ordinary UI simulations use the displayed character stats.

# Local Dev Installation

This project requires Go >=1.23.4, protobuf-compiler and the Go protobuf plugin,
and Node >=20.

## Ubuntu
Do not use apt to install any dependencies, the versions they install are all too old.
Script below will curl latest versions and install them.
```sh
# Standard Go installation script
curl -O https://dl.google.com/go/go1.24.13.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.24.13.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> $HOME/.bashrc
echo 'export GOPATH=$HOME/go' >> $HOME/.bashrc
echo 'export PATH=$PATH:$GOPATH/bin' >> $HOME/.bashrc
source $HOME/.bashrc

cd forever

# Install protobuf compiler and Go plugins
sudo apt update && sudo apt upgrade
sudo apt install protobuf-compiler
go get -u -v google.golang.org/protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Install node
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash
nvm install 20.13.1

# Install the npm package dependencies using node
npm install
```

## Docker
Alternatively, install Docker and your workflow will look something like this:
```sh
git clone https://github.com/ElliotWood/forever.git
cd forever

# Build the docker image and install npm dependencies (only need to run these once).
docker build --tag wowsims-classic .
docker run --rm -v $(pwd):/classic wowsims-classic npm install

# Now you can run the commands as shown in the Commands sections, preceding everything with, "docker run --rm -it -p 8080:8080 -v $(pwd):/classic wowsims-classic".
# For convenience, set this as an environment variable:
CLASSIC_CMD="docker run --rm -it -p 8080:8080 -v $(pwd):/classic wowsims-classic"

#For the watch commands assign this environment variable:
CLASSIC_WATCH_CMD="docker run --rm -it -p 8080:8080 -p 3333:3333 -p 5173:5173 -e WATCH=1 -v $(pwd):/classic wowsims-classic"

# ... do some coding on the sim ...

# Run tests
$(echo $CLASSIC_CMD) make test

# ... do some coding on the UI ...

# Host a local site
$(echo $CLASSIC_CMD) make host
```

## Windows
If you want to develop on Windows, we recommend setting up a Ubuntu virtual machine (VM) or running Docker using [this guide](https://docs.docker.com/desktop/windows/wsl/ "https://docs.docker.com/desktop/windows/wsl/") and then following the Ubuntu or Docker instructions, respectively.

## Mac OS
* Docker is available in OS X as well, so in theory similar instructions should work for the Docker method
* You can also use the Ubuntu setup instructions as above to run natively, with a few modifications:
  * You may need a different Go installer if `go1.18.3.linux-amd64.tar.gz` is not compatible with your system's architecture; you can do the Go install manually from `https://go.dev/doc/install`.
  * OS X uses Homebrew instead of apt, so in order to install protobuf-compiler you’ll instead need to run `brew install protobuf-c` (note the package name is also a little different than in apt). You might need to first update or upgrade brew.
  * The provided install script for Node will not included a precompiled binary for OS X, but it’s smart enough to compile one. Be ready for your CPU to melt on running `curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash`.

# Commands
We use a makefile for our build system. These commands will usually be all you need while developing for this project:
```sh
# Installs a pre-commit git hook so that your go code is automatically formatted (if you don't use an IDE that supports that).  If you want to manually format go code you can run make fmt.
# Also installs `air` to reload the dev servers automatically
make setup

# Run all the tests. Currently only the backend sim has tests.
make test

# Update the expected test results. This will need to be run after adding/removing any tests, and also if test results change due to code changes.
make update-tests

# Host a local version of the UI at http://localhost:8080. Visit it by pointing a browser to
# http://localhost:8080/classic/YOUR_SPEC_HERE, where YOUR_SPEC_HERE is the directory under ui/ with your custom code.
# Recompiles the entire client before launching using `make dist/classic`
make host

# With file-watching so the server auto-restarts and recompiles on Go or TS changes:
WATCH=1 make host

# Delete all generated files (.pb.go and .ts proto files, and dist/)
make clean

# Recompiles the ts only for the given spec (e.g. make host_elemental_shaman)
make host_$spec

# Recompiles the `wowsimclassic` server binary and runs it, hosting /dist directory at http://localhost:3333/classic.
# This is the fastest way to iterate on core go simulator code so you don't have to wait for client rebuilds.
# To rebuild client for a spec just do 'make $spec' and refresh browser.
make rundevserver

# With file-watching so the server auto-restarts and recompiles on Go or TS changes:
WATCH=1 make rundevserver


# The same as rundevserver, recompiles  `wowsimclassic` binary and runs it on port 3333. Instead of serving content from the dist folder,
# this command also runs `vite serve` to start the Vite dev server on port 5173 (or similar) and automatically reloads the page on .ts changes in less than a second.
# This allows for more rapid development, with sub second reloads on TS changes. This combines the benefits of `WATCH=1 make rundevserver` and `WATCH=1 make host`
# to create something that allows you to work in any part of the code with ease and speed.
# This might get rolled into `WATCH=1 make rundevserver` at some point.
WATCH=1 make devmode

# This is just the same as rundevserver currently
make devmode

# This command recompiles the workers in the /ui/worker folder for easier debugging/development
# Can be used with or without WATCH command
make webworkers

# With file watch enabled
WATCH=1 make webworkers

# Creates the 'wowsimclassic' binary that can host the UI and run simulations natively (instead of with wasm).
# Builds the UI and the compiles it into the binary so that you can host the sim as a server instead of wasm on the client.
# It does this by first doing make dist/classic and then copying all those files to binary_dist/classic and loading all the files in that directory into its binary on compile.
make wowsimclassic

# Using the --usefs flag will instead of hosting the client built into the binary, it will host whatever code is found in the /dist directory.
# Use --wasm to host the client with the wasm simulator.
# The server also disables all caching so that refreshes should pickup any changed files in dist/. The client will still call to the server to run simulations so you can iterate more quickly on client changes.
# make dist/classic && ./wowsimclassic --usefs would rebuild the whole client and host it. (you would have had to run `make devserver` to build the wowsimclassic binary first.)
./wowsimclassic --usefs

# Generate code for items. Only necessary if you changed the items generator.
make items
```

# Converting a class to Forever

The nine classes are already converted; this is the shape of it if a tree needs revisiting.

1. Regenerate the tree and the proto message together with `tools/forever_talents/import_talents.py $CLASS --write`, then `make proto`. They have to stay in the same order or every saved talent string breaks.
2. Implement the talents in `sim/$CLASS/talents.go`. A talent that fed a `RaidBuffs` or `Debuffs` field other classes read went baseline in Forever; a self-only stat passive was deleted outright.
3. Gate anything that changes existing Classic behaviour on `sim.IsForever()` or `character.Env.IsForever()`, and prove Classic did not move: run the spec under `RulesetClassic` before and after, and diff the DPS and TPS.
4. `make test && make update-tests`, then `./node_modules/.bin/tsc --noEmit` — use the pinned binary, not `npx tsc`, which resolves a TypeScript the repo's tsconfig rejects.

# Adding a Sim
So you want to make a new sim for your class/spec! The basic steps are as follows:
 - [Create the proto interface between sim and UI.](#create-the-proto-interface-between-sim-and-ui)
 - [Implement the UI.](#implement-the-ui)
 - [Implement the sim.](#implement-the-sim)
 - [Launch the site.](#launch-the-site)


## Create the proto interface between Sim and UI
This project uses [Google Protocol Buffers](https://developers.google.com/protocol-buffers/docs/gotutorial "https://developers.google.com/protocol-buffers/docs/gotutorial") to pass data between the sim and the UI. TLDR; Describe data structures in .proto files, and the tool can generate code in any programming language. It lets us avoid repeating the same code in our Go and Typescript worlds without losing type safety.

For a new sim, make the following changes:
  - Add a new value to the `Spec` enum in proto/common.proto. __NOTE: The name you give to this enum value is not just a name, it is used in our templating system. This guide will refer to this name as `$SPEC` elsewhere.__
  - Add a 'proto/YOUR_CLASS.proto' file if it doesn't already exist and add data messages containing all the class/spec-specific information needed to run your sim.
  - Update the `PlayerOptions.spec` field in `proto/api.proto` to include your shiny new message as an option.

That's it! Now when you run `make` there will be generated .go and .ts code in `sim/core/proto` and `ui/core/proto` respectively. If you aren't familiar with protos, take a quick look at them to see what's happening.

## Implement the UI
The UI and sim can be done in either order, but it is generally recommended to build the UI first because it can help with debugging. The UI is very generalized and it doesn't take much work to build an entire sim UI using our templating system. To use it:
  - Modify `ui/core/proto_utils/utils.ts` to include boilerplate for your `$SPEC` name if it isn't already there.
  - Create a directory `ui/$SPEC`. So if your Spec enum value was named, `elemental_shaman`, create a directory, `ui/elemental_shaman`.
  - Copy+paste from another spec's UI code.
  - Modify all the files for your spec; most of the settings are fairly obvious, if you need anything complex just ask and we can help!
  - Finally, add a rule to the `makefile` for the new sim site. Just copy from the other site rules already there and change the `$SPEC` names.

No .html is needed, it will be generated based on `ui/index_template.html` and the `$SPEC` name.

When you're ready to try out the site, run `make host` and navigate to `http://localhost:8080/classic/$SPEC`.

## Implement the Sim
This step is where most of the magic happens. A few highlights to start understanding the sim code:
  - `sim/wasm/main.go` This file is the actual main function, for the [.wasm binary](https://webassembly.org/ "https://webassembly.org/") used by the UI. You shouldn't ever need to touch this, but just know its here.
  - `sim/core/api.go` This is where the action starts. This file implements the request/response messages defined in `proto/api.proto`.
  - `sim/core/sim.go` Orchestrates everything. Main event loop is in `Simulation.RunOnce`.
  - `sim/core/agent.go` An Agent can be thought of as the 'Player', i.e. the person controlling the game. This is the interface you'll be implementing.
  - `sim/core/character.go` A Character holds all the stats/cooldowns/gear/etc common to any WoW character. Each Agent has a Character that it controls.

Read through the core code and some examples from other classes/specs to get a feel for what's needed. Hopefully `sim/core` already includes what you need, but most classes have at least 1 unique mechanic so you may need to touch `core` as well.

Finally, add your new sim to `RegisterAll()` in `sim/register_all.go`.

Don't forget to write unit tests! Again, look at existing tests for examples. Run them with `make test` when you're ready.

# Launch the site
When everything is ready for release, modify `ui/core/launched_sims.ts` and `ui/index.html` to include the new spec value. This will add the sim to the dropdown menu so anyone can find it from the existing sims. This will also remove the UI warning that the sim is under development. Now tell everyone about your new sim!

# Add your spec to the raid sim
Don't touch the raid sim until the individual sim is ready for launch; anything in the raid sim is publicly accessible. To add your new spec to the raid sim, do the following:
 - Add a reference to the individual sim in `ui/raid/tsconfig.json`. DO NOT FORGET THIS STEP or Typescipt will silently do very bad things.
 - Import the individual sim's css file from `ui/raid/index.scss`.
 - Update `ui/raid/presets.ts` to include a constructor factory in the `specSimFactories` variable and add configurations for new Players in the `playerPresets` variable.

# Deployment
`.github/workflows/deploy.yml` is inherited from upstream and deploys on pushes to `master` there. Actions do not run on this fork, so there is no site to deploy to and no published build — host it locally with `make host`.
