package main

import(
	"fmt"
	"os"
)

func main() {
	length := len(os.Args)
	if length < 2 || length > 3 {
		fmt.Println("Error: Invalid number of arguments")
		return
	}

	bannerFile := "standard.txt"
	if length == 3 {
		bannerFile = os.Args[2] + ".txt"
	}

	input := os.Args[1]

	banner, err := loadBanner(bannerFile)
	if err != nil {
		fmt.Println("Error reading banner:", err)
		return
	}

	// Catch what printArt returns, then print it
	result := printArt(input, banner)
	fmt.Print(result)
}
