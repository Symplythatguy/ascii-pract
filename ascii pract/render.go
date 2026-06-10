package main

import(
	"fmt"
	"strings"
)

func printArt(input string, banner []string) {
	lines := strings.Split(input, "\\n")

	for _, word := range lines {
		if word == "" {
			fmt.Println()
			continue
		}

		for i := 0; i <= 8; i++ {
			for j := 0; j < len(word); j++ {
				index := int(word[j]-32)*9 + i
				if index >= 0 && index < len(banner) {
					fmt.Print(banner[index])
				}
			}
			fmt.Println()
		}
	}
}