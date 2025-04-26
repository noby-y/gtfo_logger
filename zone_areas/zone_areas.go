package zone_areas

type RangeLetter struct {
	Min, Max int
	Letter   string
}

var E_528 = []RangeLetter{
	{0, 3, "A"},
	{4, 12, "B"},
	{13, 16, "C"},
	{17, 22, "D"},
	{23, 28, "E"},
	{29, 34, "F"},
	{35, 35, "G"},
	{36, 39, "H"},
	{40, 40, "I"},
	{41, 49, "J"},
	{50, 50, "K"},
	{51, 51, "L"},
}

var E_533 = []RangeLetter{
	{0, 13, "A"},
	{14, 19, "B"},
	{20, 33, "C"},
	{34, 39, "D"},
	{40, 45, "E"},
}

var E_DEFAULT = []RangeLetter{
	{0, 9999, "?"},
}
