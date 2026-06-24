package main

import(
	"strings"
)

func printArt(input string, banner []string) string {
	// Our bucket to collect everything
	var result strings.Builder

	lines := strings.Split(input, "\\n")

	for _, word := range lines {
		if word == "" {
			result.WriteString("\n") // Collect instead of printing
			continue
		}

		for i := 0; i <= 8; i++ {
			for j := 0; j < len(word); j++ {
				index := int(word[j]-32)*9 + i
				if index >= 0 && index < len(banner) {
					result.WriteString(banner[index]) // Drop into bucket
				}
			}
			result.WriteString("\n") // Newline after each line
		}
	}

	// Hand back everything in the bucket
	return result.String()
}