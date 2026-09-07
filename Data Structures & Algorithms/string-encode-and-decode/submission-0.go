type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder

	for _, str := range strs {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteString("#")
		sb.WriteString(str)	
	}

	result := sb.String()
	return result
}

func (s *Solution) Decode(encoded string) []string {
	decoded := []string{}
	// fmt.Printf("Decoding: %s\n", encoded)
	for idx := int64(0); idx < int64(len(encoded)); {
		var numberSb strings.Builder
		for encoded[idx] != '#' {
			numberSb.WriteByte(encoded[idx])
			idx += 1
		}
		// fmt.Printf("Got number %s\n", numberSb.String())
		strLen, _ := strconv.ParseInt(numberSb.String(), 10, 16)
		idx += 1

		nextString := encoded[idx:idx + strLen]
		idx += strLen
		// fmt.Printf("Got next string: %s\n", nextString)

		decoded = append(decoded, nextString)
	}

	return decoded
}
