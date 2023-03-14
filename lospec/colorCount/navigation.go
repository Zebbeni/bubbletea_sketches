package colorCount

var (
	navDownLookup = map[Focus]Focus{
		ColorNumberFilterType: ColorNumberForm,
		ColorNumberForm:       ColorNumberForm,
	}
	navUpLookup = map[Focus]Focus{
		ColorNumberFilterType: ColorNumberFilterType,
		ColorNumberForm:       ColorNumberFilterType,
	}
)
