package colorNumberFilterType

type option struct {
	name  string
	value FilterType
}

func buildOptions() []option {
	return []option{
		{name: "Any", value: Any},
		{name: "Max", value: Max},
		{name: "Min", value: Min},
		{name: "Exact", value: Exact},
	}
}
