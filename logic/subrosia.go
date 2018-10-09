package logic

// subrosia has several large areas which are guaranteed to be traverseable as
// long as you can get there in the first place:
//
// 1. "temple": rosa portal, dance hall, temple, smithy
// 2. "beach": swamp portal, market, beach
// 3. "hide and seek": H&S, mountain portal, spring tower
// 4. "pirate house": village portal, pirates
// 5. "furnace": lake portal, furnace, bomb flower
// 6. "bridge": bridge area (large but not visited in any%)
//
// the other locations are isolated and only traverseable with some combination
// of jumping and boulder removal.

var subrosiaNodes = map[string]*Node{
	"temple": Or("rosa portal",
		And("beach", "ribbon"),
		And("beach", "jump"),
		And("hide and seek", "pegasus jump L-2"),
		And("bridge", "jump")),

	"beach": Or("swamp portal",
		And("hide and seek", "bracelet", "long jump"),
		And("hide and seek", "jump", "bracelet", "magnet gloves"),
		And("furnace", "bracelet", "jump"),
		And("furnace", Or("feather L-2", Hard("long jump"))),
		And("furnace", "jump", "magnet gloves"),
		And("temple", "jump")),

	"hide and seek": Or("mountain portal",
		And("pirate house", "jump"),
		And("temple", "pegasus jump L-2"),
		And("bridge", "pegasus jump L-2")),

	"pirate house": Or("village portal", And("hide and seek", "jump")),

	"furnace": Or("lake portal",
		And("beach", Or("feather L-2", Hard("long jump"))),
		And("beach", "magnet gloves", "jump")),

	"bridge": Or(
		And("temple", "jump"),
		And("remains portal", "bracelet", "feather L-2"),
		And("hide and seek", "pegasus jump L-2")),

	"dance hall prize":   AndSlot("temple"),
	"rod gift":           AndSlot("temple"),
	"star ore spot":      AndSlot("beach", "shovel"),
	"pirate's bell":      And("temple", "rusty bell"),
	"cross winter tower": Or("hit far switch", "jump"),
	"winter tower":       AndSlot("temple", "cross winter tower"),
	"summer tower":       AndSlot("beach", "ribbon", "bracelet"),
	"spring tower":       AndSlot("hide and seek", "jump"),
	"autumn tower":       AndSlot("temple", "jump", "bomb flower"),
	"blue ore chest": AndSlot("hide and seek",
		Or("feather L-2", "magnet gloves")),
	"red ore chest":        AndSlot("furnace", Or("feather L-2", "magnet gloves")),
	"non-rosa gasha chest": AndSlot("bridge"),
	"rosa gasha chest":     AndSlot("bridge", "beach", "ribbon", "jump"),
	"subrosian market 1":   AndSlot("beach", "star ore"),
	"subrosian market 2":   AndSlot("beach", "ore chunks", "ember seeds"),
	"subrosian market 5":   AndSlot("beach", "ore chunks"),
	"hard ore slot": AndSlot("furnace", "red ore", "blue ore",
		"temple", "bomb flower"),
	"iron shield gift": AndSlot("temple", "hard ore"),

	"enter d8": Or("d8 portal"),
}
