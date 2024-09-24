package asciiweb

func PrintLineAsAscii(line string, asciiGraph []string) string {
	var asciiChars []string

	partialArt := ""
	if line != "" {
		for _, char := range line {
			// add each chars to the asciiChars line by line
			for i := 8; i > 0; i-- {
				asciiChars = append(asciiChars, string(asciiGraph[findLastLine(char)-i]))
			}
		}

		for i := 0; i < 8; i++ {
			for j := 0; j < len(asciiChars); j += 8 {
				partialArt += asciiChars[i+j]
			}
			partialArt += "\n"

		}
	} else {
		partialArt += "\n"
	}
	return partialArt
}

// find the last line after the char
func findLastLine(char rune) int {
	return int((char - 31) * (9))
}
