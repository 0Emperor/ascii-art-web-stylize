package asciiweb

import (
	"strings"
)

func GenerateAsciiArt(input, bannerFile string) string {
	asciiGraph, err := ReadFile(bannerFile)
	if err != nil {
		return "Error"
	}
	lines := strings.Split(input, "\r\n")

	finalART := ""
	for _, line := range lines {
		finalART += PrintLineAsAscii(line, asciiGraph)
	}
	return finalART
}
