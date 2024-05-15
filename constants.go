package goterra

var Factions = map[int]string{
	0:  "DE",
	1:  "FR",
	2:  "IO",
	3:  "NX",
	4:  "PZ",
	5:  "SI",
	6:  "BW",
	9:  "MT",
	7:  "SH",
	10: "BC",
	12: "RU",
}

func idFromFactionString(lv string) int {
	for i, v := range Factions {
		if v == lv {
			return i
		}
	}
	return -1
}

var FactionsVersion = map[uint8][2]uint8{
	0:  {1, 1},
	1:  {1, 1},
	2:  {1, 1},
	3:  {1, 1},
	4:  {1, 1},
	5:  {1, 1},
	6:  {1, 2},
	9:  {1, 2},
	7:  {1, 3},
	10: {1, 4},
	12: {1, 5},
}
