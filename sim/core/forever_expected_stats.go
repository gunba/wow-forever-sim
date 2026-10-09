package core

// Compact client baselines, not a decoded Rage script.
// Source: https://wago.tools/db2/ExpectedStat/csv?build=1.60.1.70291
// ExpansionID=-2, ContentSetID=0, levels 1-60.
// Captured CSV SHA256: f8e8918417d6f852bd4e443d80c534db89b489915d7127de6d48d4c3d5d92848
// The provisional incoming-Rage model chooses player level and the reference
// Armor ratio from these rows. Neither choice is verified server behavior.
type foreverExpectedStat struct {
	creatureHealth float64
	creatureArmor  float64
	armorConstant  float64
}

var foreverExpectedStats = [...]foreverExpectedStat{
	{},                 // No level zero row.
	{42, 20, 485},      // 1
	{55, 21, 570},      // 2
	{71, 46, 655},      // 3
	{86, 82, 740},      // 4
	{102, 126, 825},    // 5
	{120, 180, 910},    // 6
	{137, 245, 995},    // 7
	{156, 322, 1080},   // 8
	{180, 412, 1165},   // 9
	{208, 518, 1250},   // 10
	{239, 545, 1335},   // 11
	{272, 580, 1420},   // 12
	{307, 615, 1505},   // 13
	{345, 650, 1590},   // 14
	{385, 685, 1675},   // 15
	{427, 721, 1760},   // 16
	{473, 756, 1845},   // 17
	{521, 791, 1930},   // 18
	{572, 826, 2015},   // 19
	{629, 861, 2100},   // 20
	{677, 897, 2185},   // 21
	{731, 932, 2270},   // 22
	{787, 967, 2355},   // 23
	{846, 1002, 2440},  // 24
	{909, 1037, 2525},  // 25
	{975, 1072, 2610},  // 26
	{1040, 1108, 2695}, // 27
	{1109, 1142, 2780}, // 28
	{1177, 1177, 2865}, // 29
	{1242, 1212, 2950}, // 30
	{1308, 1247, 3035}, // 31
	{1374, 1283, 3120}, // 32
	{1443, 1317, 3205}, // 33
	{1512, 1353, 3290}, // 34
	{1586, 1387, 3375}, // 35
	{1660, 1494, 3460}, // 36
	{1737, 1607, 3545}, // 37
	{1814, 1724, 3630}, // 38
	{1897, 1849, 3715}, // 39
	{1981, 1980, 3800}, // 40
	{2061, 2117, 3885}, // 41
	{2146, 2262, 3970}, // 42
	{2231, 2414, 4055}, // 43
	{2317, 2574, 4140}, // 44
	{2402, 2742, 4225}, // 45
	{2495, 2798, 4310}, // 46
	{2587, 2853, 4395}, // 47
	{2681, 2907, 4480}, // 48
	{2779, 2963, 4565}, // 49
	{2909, 3018, 4650}, // 50
	{3040, 3072, 4735}, // 51
	{3174, 3128, 4820}, // 52
	{3317, 3183, 4905}, // 53
	{3458, 3237, 4990}, // 54
	{3602, 3292, 5075}, // 55
	{3755, 3348, 5160}, // 56
	{3909, 3402, 5245}, // 57
	{4105, 3457, 5330}, // 58
	{4234, 3512, 5415}, // 59
	{4365, 3566, 5500}, // 60
}

func foreverExpectedStatAtLevel(level int32) (foreverExpectedStat, bool) {
	if level < 1 || int(level) >= len(foreverExpectedStats) {
		return foreverExpectedStat{}, false
	}
	return foreverExpectedStats[level], true
}
