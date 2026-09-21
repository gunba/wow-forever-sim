// Mirrors the Wowhead images the site displays into assets/img/wowhead, so the pages never
// load icons from wow.zamimg.com at runtime. Some networks reject that host's TLS setup,
// which left every class and spec icon broken while the rest of the site worked.
//
// The set is derived from what the UI can show: every icon named in the database and its
// inputs, the talent trees (icons, tree backgrounds, and the icons of talent spells the
// database does not carry, resolved through Wowhead's tooltip endpoint the same way the UI
// does), and every image the UI source references: paths under WOWHEAD_IMAGES and the bare
// icon file names of the talent tree icon tables.
//
// Only missing files are downloaded, so re-running after a database refresh is cheap:
//
//	go run ./tools/icons
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const zamimg = "https://wow.zamimg.com/images/"

var (
	outDir     = flag.String("outDir", "assets/img/wowhead", "directory the images are mirrored into")
	dbFile     = flag.String("db", "assets/database/db.json", "database to take icon names from")
	inputsDir  = flag.String("inputs", "assets/db_inputs", "tooltip csv directory to take icon names from")
	treesDir   = flag.String("trees", "ui/core/talents/trees", "talent tree directory")
	uiDir      = flag.String("ui", "ui", "UI source directory scanned for image references")
	benchFile  = flag.String("bench", "artifacts/forever_dps_5min.json", "optional benchmark containing displayed spell IDs")
	numWorkers = flag.Int("workers", 8, "parallel downloads")
)

// Paths are relative to outDir and mirror Wowhead's layout under images/wow/, except the
// handful of site icons that live directly under images/icons/ (boss.gif, horde.png, ...),
// which sit in icons/ with no size directory.
func remoteUrl(localPath string) string {
	if dir, _ := filepath.Split(localPath); dir == "icons/" {
		return zamimg + localPath
	}
	return zamimg + "wow/" + localPath
}

var (
	iconFieldRegex = regexp.MustCompile(`\\?"icon\\?":\\?"([A-Za-z0-9_\-]+)\\?"`)
	uiImageRegex   = regexp.MustCompile(`WOWHEAD_IMAGES\}([A-Za-z0-9_./\-]+\.(?:jpg|png|gif))`)
	// Icon file names the UI holds without a path (the talent tree icon tables), all large.
	bareIconRegex   = regexp.MustCompile(`'([a-z0-9_]+)\.jpg'`)
	htmlImageRegex  = regexp.MustCompile(`assets/img/wowhead/([A-Za-z0-9_./\-]+\.(?:jpg|png|gif))`)
	spellIDRegex    = regexp.MustCompile(`"?(?:spellId|foreverId)"?\s*:\s*(\d+)`)
	zamimgWowPrefix = zamimg + "wow/"
)

type talentTree struct {
	BackgroundUrl string `json:"backgroundUrl"`
	Talents       []struct {
		SpellIds []int32 `json:"spellIds"`
		Icon     string  `json:"icon"`
	} `json:"talents"`
}

type imageSet map[string]bool

func (s imageSet) addIcon(name string) {
	if name != "" {
		s["icons/large/"+strings.ToLower(name)+".jpg"] = true
	}
}

// Every "icon" field anywhere in the file, whether it is JSON or a CSV of JSON tooltips.
func (s imageSet) addIconFields(filePath string) {
	for _, match := range iconFieldRegex.FindAllStringSubmatch(readFile(filePath), -1) {
		s.addIcon(match[1])
	}
}

// The tools package is not imported for its helpers because it pulls in the sim, and with it
// the generated protos, which the Update Icons workflow has no reason to build.
func readFile(filePath string) string {
	b, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to open %s: %s", filePath, err)
	}
	return string(b)
}

func main() {
	flag.Parse()
	images := imageSet{}

	images.addIconFields(*dbFile)
	inputs, _ := filepath.Glob(filepath.Join(*inputsDir, "*.csv"))
	for _, input := range inputs {
		images.addIconFields(input)
	}

	dbSpellIds := databaseSpellIds()
	unresolvedSpells := map[int32]bool{}
	addSpellIDs := func(source string) {
		for _, match := range spellIDRegex.FindAllStringSubmatch(source, -1) {
			id, _ := strconv.ParseInt(match[1], 10, 32)
			if id > 0 && !dbSpellIds[int32(id)] {
				unresolvedSpells[int32(id)] = true
			}
		}
	}
	trees, _ := filepath.Glob(filepath.Join(*treesDir, "*.json"))
	for _, treeFile := range trees {
		var classTrees []talentTree
		if err := json.Unmarshal([]byte(readFile(treeFile)), &classTrees); err != nil {
			log.Fatalf("Failed to parse %s: %s", treeFile, err)
		}
		for _, tree := range classTrees {
			if strings.HasPrefix(tree.BackgroundUrl, zamimgWowPrefix) {
				images[strings.TrimPrefix(tree.BackgroundUrl, zamimgWowPrefix)] = true
			}
			for _, talent := range tree.Talents {
				images.addIcon(talent.Icon)
				for _, spellId := range talent.SpellIds {
					if spellId != 0 && !dbSpellIds[spellId] {
						unresolvedSpells[spellId] = true
					}
				}
			}
		}
	}
	filepath.WalkDir(*uiDir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		switch filepath.Ext(path) {
		case ".ts", ".tsx":
			source := readFile(path)
			for _, match := range uiImageRegex.FindAllStringSubmatch(source, -1) {
				images[match[1]] = true
			}
			for _, match := range bareIconRegex.FindAllStringSubmatch(source, -1) {
				images.addIcon(match[1])
			}
			addSpellIDs(source)
		case ".json":
			images.addIconFields(path)
			addSpellIDs(readFile(path))
		case ".html":
			for _, match := range htmlImageRegex.FindAllStringSubmatch(readFile(path), -1) {
				images[match[1]] = true
			}
		}
		return nil
	})
	if _, err := os.Stat(*benchFile); err == nil {
		addSpellIDs(readFile(*benchFile))
	}
	spellIDs := make([]int32, 0, len(unresolvedSpells))
	for id := range unresolvedSpells {
		spellIDs = append(spellIDs, id)
	}
	sort.Slice(spellIDs, func(i, j int) bool { return spellIDs[i] < spellIDs[j] })
	for _, icon := range tooltipIcons(spellIDs) {
		images.addIcon(icon)
	}

	var missing []string
	for localPath := range images {
		if _, err := os.Stat(filepath.Join(*outDir, localPath)); err != nil {
			missing = append(missing, localPath)
		}
	}
	sort.Strings(missing)
	fmt.Printf("%d images referenced, %d missing\n", len(images), len(missing))

	// A file Wowhead itself does not serve is not worth failing the run over: the page was
	// already showing a broken image for it.
	if failed := download(missing); len(failed) > 0 {
		sort.Strings(failed)
		fmt.Printf("%d images could not be fetched:\n%s\n", len(failed), strings.Join(failed, "\n"))
	}
}

func databaseSpellIds() map[int32]bool {
	var db struct {
		SpellIcons []struct {
			Id int32 `json:"id"`
		} `json:"spellIcons"`
	}
	if err := json.Unmarshal([]byte(readFile(*dbFile)), &db); err != nil {
		log.Fatalf("Failed to parse %s: %s", *dbFile, err)
	}
	ids := make(map[int32]bool, len(db.SpellIcons))
	for _, spell := range db.SpellIcons {
		ids[spell.Id] = true
	}
	return ids
}

// The UI asks Wowhead for spells the database lacks (ui/core/proto_utils/database.ts), so
// the icons it ends up displaying for them come from the same endpoint.
func tooltipIcons(spellIds []int32) []string {
	var icons []string
	var mu sync.Mutex
	parallel(len(spellIds), func(i int) {
		spellId := spellIds[i]
		body, err := get(fmt.Sprintf("https://nether.wowhead.com/forever/tooltip/spell/%d?lvl=60", spellId))
		var tooltip struct {
			Icon string `json:"icon"`
		}
		if err == nil {
			err = json.Unmarshal(body, &tooltip)
		}
		if err != nil || tooltip.Icon == "" {
			log.Printf("No tooltip icon for spell %d", spellId)
			return
		}
		mu.Lock()
		icons = append(icons, tooltip.Icon)
		mu.Unlock()
	})
	return icons
}

func download(localPaths []string) []string {
	var failed []string
	var mu sync.Mutex
	parallel(len(localPaths), func(i int) {
		localPath := localPaths[i]
		if err := fetch(remoteUrl(localPath), filepath.Join(*outDir, localPath)); err != nil {
			log.Printf("%s: %s", localPath, err)
			mu.Lock()
			failed = append(failed, localPath)
			mu.Unlock()
		}
	})
	return failed
}

func fetch(url string, filePath string) error {
	body, err := get(url)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return err
	}
	return os.WriteFile(filePath, body, 0644)
}

func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s returned %s", url, resp.Status)
	}
	return io.ReadAll(resp.Body)
}

// Runs fn for every index in [0, n) on numWorkers goroutines.
func parallel(n int, fn func(i int)) {
	var wg sync.WaitGroup
	queue := make(chan int)
	for w := 0; w < *numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range queue {
				fn(i)
			}
		}()
	}
	for i := 0; i < n; i++ {
		queue <- i
	}
	close(queue)
	wg.Wait()
}
