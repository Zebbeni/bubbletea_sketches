package app

var (
	navDownLookup = map[Focus]Focus{
		ColorCount:  TagSearch,
		TagSearch:   Sorting,
		Sorting:     PaletteList,
		PaletteList: PaletteList,
	}
	navUpLookup = map[Focus]Focus{
		ColorCount:  ColorCount,
		TagSearch:   ColorCount,
		Sorting:     TagSearch,
		PaletteList: Sorting,
	}
)
