func generatehs(str string) string {
	start := byte('a')
	bs := []byte(str)
	var charCount []int = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	for _, c := range bs {
		charCount[c - start] = charCount[c - start] + 1
	}

	s := ""
	for _, count := range(charCount) {
		s += "," + string(count)
	}

	return s
}


func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	for _, s := range strs {
		hs := generatehs(s)
		// fmt.Printf("%s => %s\n", s, hs)
		if _, ok := groups[hs]; ok {
			groups[hs] = append(groups[hs], s)
			continue
		}
		groups[hs] = []string{s}
	}

	result := [][]string{}
	for _, s := range(groups) {
		result = append(result, s)
	}

	return result
}
