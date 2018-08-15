package main

import (
	"math/rand"
	"testing"

	"github.com/jangler/oos-randomizer/graph"
	"github.com/jangler/oos-randomizer/prenode"
)

// item placement from a 1.2.2 rom used for a race, updated for rod/seasons
// being slotted
var testData1 = map[string]string{
	"winter":             "d0 sword chest",
	"bracelet":           "maku tree gift",
	"gnarled key":        "blaino gift",
	"satchel":            "d1 satchel spot",
	"gale tree seeds 1":  "scent tree",
	"ember tree seeds":   "ember tree",
	"feather L-1":        "rod gift",
	"magnet gloves":      "shovel gift",
	"slingshot L-2":      "dance hall prize",
	"mystery tree seeds": "sunken gale tree",
	"star ore":           "x-shaped jewel chest",
	"pegasus tree seeds": "pegasus tree",
	"master's plaque":    "floodgate key spot",
	"fool's ore":         "diver gift",
	"shovel":             "d8 HSS chest",
	"x-shaped jewel":     "star ore spot",
	"sword L-1":          "d2 bracelet chest",
	"feather L-2":        "master's plaque chest",
	"round jewel":        "dragon key spot",
	"dragon key":         "spring banana tree",
	"floodgate key":      "desert pit",
	"square jewel":       "d5 magnet gloves chest",
	"flippers":           "d3 feather chest",
	"pyramid jewel":      "d4 slingshot chest",
	"slingshot L-1":      "pyramid jewel spot",
	"boomerang L-2":      "square jewel chest",
	"rusty bell":         "round jewel gift",
	"gale tree seeds 2":  "mystery tree",
	"sword L-2":          "d6 boomerang chest",
	"spring banana":      "d7 cape chest",
	"boomerang L-1":      "noble sword spot",
	"scent tree seeds":   "tarm gale tree",

	"spring": "spring tower",
	"summer": "summer tower",
	"autumn": "autumn tower",
}

// item placement from a 1.3.0 rom used for a race, updated for rod/seasons
// being slotted
var testData2 = map[string]string{
	"fool's ore":         "d0 sword chest",
	"bracelet":           "maku tree gift",
	"slingshot L-2":      "blaino gift",
	"mystery tree seeds": "ember tree",
	"pyramid jewel":      "x-shaped jewel chest",
	"ember tree seeds":   "scent tree",
	"winter":             "dance hall prize",
	"feather L-1":        "d2 bracelet chest",
	"magnet gloves":      "rod gift",
	"gale tree seeds 1":  "sunken gale tree",
	"shovel":             "shovel gift",
	"gnarled key":        "star ore spot",
	"master's plaque":    "d1 satchel spot",
	"sword L-2":          "diver gift",
	"satchel":            "master's plaque chest",
	"pegasus tree seeds": "mystery tree",
	"star ore":           "d8 HSS chest",
	"boomerang L-2":      "floodgate key spot",
	"spring banana":      "spring banana tree",
	"round jewel":        "desert pit",
	"flippers":           "dragon key spot",
	"slingshot L-1":      "round jewel gift",
	"square jewel":       "pyramid jewel spot",
	"floodgate key":      "d5 magnet gloves chest",
	"rusty bell":         "square jewel chest",
	"dragon key":         "d7 cape chest",
	"feather L-2":        "d4 slingshot chest",
	"x-shaped jewel":     "d3 feather chest",
	"gale tree seeds 2":  "tarm gale tree",
	"sword L-1":          "noble sword spot",
	"scent tree seeds":   "pegasus tree",
	"ricky's gloves":     "d6 boomerang chest",

	"temple remains default summer": "start",
	"temple remains default winter": "",
	"north horon default autumn":    "start",
	"north horon default winter":    "",
	"holodrum plain default winter": "start",
	"holodrum plain default spring": "",
	"lost woods default spring":     "start",
	"lost woods default autumn":     "",
	"tarm ruins default summer":     "start",
	"tarm ruins default spring":     "",
	"spool swamp default summer":    "start",
	"spool swamp default autumn":    "",

	"spring": "spring tower",
	"summer": "summer tower",
	"autumn": "autumn tower",
}

// partial item placement from a dev ~1.4 build, updated for rod/seasons being
// slotted
var testData3 = map[string]string{
	"fool's ore":        "d0 sword chest",
	"satchel":           "maku tree gift",
	"ember tree seeds":  "ember tree",
	"feather L-1":       "dance hall prize",
	"gale tree seeds 1": "sunken gale tree",
	"winter":            "rod gift",
	"sword L-1":         "shovel gift",
	"master's plaque":   "master's plaque chest",
	"flippers":          "diver gift",

	"spool swamp default winter": "start",
	"spool swamp default autumn": "",

	"spring": "spring tower",
	"summer": "summer tower",
	"autumn": "autumn tower",
}

func TestFeatherLockCheck(t *testing.T) {
	r := NewRoute([]string{"horon village"})
	g := r.HardGraph

	// make sure reaching H&S with mandatory shovel does not error
	checkSoftlockWithSlots(t, canFeatherSoftlock, g,
		map[string]string{
			"bracelet":           "d0 sword chest",
			"flippers":           "maku tree gift",
			"shovel":             "blaino gift",
			"feather L-2":        "star ore spot",
			"winter":             "rod gift",
			"satchel":            "shovel gift",
			"pegasus tree seeds": "ember tree",
		}, "hide and seek", false)

	// make sure reaching H&S with optional shovel errors
	checkSoftlockWithSlots(t, canFeatherSoftlock, g,
		map[string]string{
			"bracelet":           "d0 sword chest",
			"flippers":           "maku tree gift",
			"feather L-2":        "blaino gift",
			"winter":             "rod gift",
			"pegasus tree seeds": "ember tree",
			"shovel":             "dance hall prize",
		}, "hide and seek", true)

	// a softlock case from a real rom produced by 1.2.2
	checkSoftlockWithSlots(t, canFeatherSoftlock, g, testData1,
		"hide and seek", true)
}

func TestD7ExitLockChest(t *testing.T) {
	r := NewRoute([]string{"horon village"})
	g := r.HardGraph

	// a softlock case from a real rom produced by 1.2.2
	checkSoftlockWithSlots(t, canD7ExitSoftlock, g, testData1,
		"enter d7", true)
}

func TestD2ExitCheck(t *testing.T) {
	r := NewRoute([]string{"horon village"})
	g := r.HardGraph

	// check for false positive
	checkSoftlockWithSlots(t, canD2ExitSoftlock, g,
		map[string]string{
			"bracelet":       "d0 sword chest",
			"ricky's gloves": "blaino gift",
			"feather L-1":    "maku tree gift",

			"eastern suburbs default winter": "start",
			"eastern suburbs default autumn": "",
		}, "central woods of winter", false)

	// check for false negative
	checkSoftlockWithSlots(t, canD2ExitSoftlock, g,
		map[string]string{
			"sword L-1":      "d0 sword chest",
			"ricky's gloves": "blaino gift",
			"flippers":       "maku tree gift",
			"summer":         "rod gift",

			"eastern suburbs default winter": "start",
			"eastern suburbs default autumn": "",
		}, "central woods of winter", true)
}

func TestSquareJewelCheck(t *testing.T) {
	r := NewRoute([]string{"horon village"})

	// check for false positive
	checkSoftlockWithSlots(t, canSquareJewelSoftlock, r.HardGraph,
		map[string]string{
			"sword L-1":        "d0 sword chest",
			"satchel":          "maku tree gift",
			"ember tree seeds": "ember tree",
			"shovel":           "dance hall prize",
			"feather L-1":      "rod gift",
			"flippers":         "star ore spot",

			"spool swamp default winter": "start",
			"spool swamp default autumn": "",
		}, "square jewel chest", false)

	// check for false negative
	checkSoftlockWithSlots(t, canSquareJewelSoftlock, r.HardGraph, testData3,
		"square jewel chest", true)
}

// helper function used for the other benchmarks
func benchGraphCheck(b *testing.B, check func(graph.Graph) error) {
	// make a list of base item nodes to use for testing
	r := NewRoute([]string{"horon village"})
	g := r.HardGraph
	baseItems := make([]*graph.Node, 0, len(prenode.BaseItems()))
	for name := range prenode.BaseItems() {
		baseItems = append(baseItems, g[name])
	}

	for i := 0; i < b.N; i++ {
		// create a fresh graph and shuffle the item list
		b.StopTimer()
		r = NewRoute([]string{"horon village"})
		g = r.HardGraph
		reached := map[*graph.Node]bool{g["horon village"]: true}

		rand.Shuffle(len(baseItems), func(i, j int) {
			baseItems[i], baseItems[j] = baseItems[j], baseItems[i]
		})
		b.StartTimer()

		// gradually add items to the graph to get a picture of performance at
		// various stages in the exploration
		for _, itemNode := range baseItems {
			itemNode.AddParents(g["d0 sword chest"])
			reached = g.Explore(reached, []*graph.Node{itemNode})

			// run 10 times to get a better proportion of check runtime vs
			// explore runtime. just ignoring the explore runtime results in
			// really long tests
			for j := 0; j < 10; j++ {
				check(g)
			}
		}
	}
}

func BenchmarkCanSoftlock(b *testing.B) {
	benchGraphCheck(b, canSoftlock)
}

func BenchmarkCanFlowerSoftlock(b *testing.B) {
	benchGraphCheck(b, canFlowerSoftlock)
}

func BenchmarkCanFeatherSoftlock(b *testing.B) {
	benchGraphCheck(b, canFeatherSoftlock)
}

func BenchmarkCanEmberSeedSoftlock(b *testing.B) {
	benchGraphCheck(b, canEmberSeedSoftlock)
}

// the keys in the "parents" map MUST be root nodes. this function errors if
// the target node cannot be reached (meaning the test setup is incorrect), or
// if the target node can be reached with the children having the assigned
// parents.
func checkSoftlockWithSlots(t *testing.T, check func(g graph.Graph) error,
	g graph.Graph, parents map[string]string, target string,
	expectError bool) {
	t.Helper()

	// add parents at the start of the function, and remove them at the end. if
	// a parent is blank, remove it at the start and add it at the end (only
	// useful for default seasons).
	for child, parent := range parents {
		if parent == "" {
			g[child].ClearParents()
		} else {
			g[child].AddParents(g[parent])
		}
	}
	defer func() {
		for child, parent := range parents {
			if parent == "" {
				g[child].AddParents(g["start"])
			} else {
				g[child].ClearParents()
			}
		}
	}()
	g.ExploreFromStart()

	softlock := check(g)

	if g[target].GetMark(g[target], nil) != graph.MarkTrue {
		t.Errorf("test invalid: cannot reach %s", target)
	} else if !expectError && softlock != nil {
		t.Errorf("false positive %s softlock", target)
	} else if expectError && softlock == nil {
		t.Errorf("false negative %s softlock", target)
	}
}
