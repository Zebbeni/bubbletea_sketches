package sortingType

type option struct {
	name  string
	value SelectedSortingType
}

func buildOptions() []option {
	return []option{
		{name: "Default", value: Default},
		{name: "A-Z", value: Alphabetical},
		{name: "Downloads", value: Downloads},
		{name: "Newest", value: Newest},
	}
}
