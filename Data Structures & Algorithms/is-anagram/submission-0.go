import "maps"

func isAnagram(s string, t string) bool {
	sf, st := map[rune]int{}, map[rune]int{}

	for _, c := range s {
		cc, ok := sf[c]
		if !ok {
			sf[c] = 1
			continue
		}

		sf[c] = cc + 1
	}


	for _, c := range t {
		cc, ok := st[c]
		if !ok {
			st[c] = 1
			continue
		}

		st[c] = cc + 1
	}

	return maps.Equal(sf, st)
}
