package ccwc

import (
	"flag"
	"fmt"
	"os"
	"path"
)

const DefaultFilePath = "./test.txt"

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func readFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	handleError(err)
	return data, nil
}

func calcChars(data string) int {
	return len(data)
}

func calcLines(data string) int {
	var count = 0
	for char := range data {
		if char == '\n' {
			count++
		}
	}
	return count
}

func calcWords(data string) int {
	if len(data) == 0 {
		return 0
	}
	var count = 0
	for char := range data {
		if char == ' ' {
			count++
		}
	}
	return count
}

func parseIntToString(num int) string {
	return string(rune(num))
}

func Ccwc() {
	args := os.Args
	filePath := os.Args[len(args)-1]
	var result string

	cFlag := flag.Bool("-c", false, "bytes count")
	lFlag := flag.Bool("-l", false, "lines count")
	wFlag := flag.Bool("-w", false, "words count")
	mFlag := flag.Bool("-m", false, "characters count")

	_, filename := path.Split(filePath)

	data, err := readFile(filePath)
	handleError(err)

	if len(args) < 3 {
		filePath = DefaultFilePath
	}

	// Default: No flags provided, so result should be the result of -c -l -w
	if !*cFlag && !*lFlag && !*wFlag && !*mFlag {
		fmt.Println("No flag provided")
		result =
			parseIntToString(len(data)) + " " +
				parseIntToString(calcLines(string(data))) + " " +
				parseIntToString(calcWords(string(data))) + " " +
				filename
	}
	if *cFlag {
		result = parseIntToString(len(data)) + " " + filename
	}
	if *lFlag {
		result = parseIntToString(calcLines(string(data))) + " " + filename
	}
	if *wFlag {
		result = parseIntToString(calcWords(string(data))) + " " + filename
	}
	if *mFlag {
		result = parseIntToString(calcChars(string(data))) + " " + filename
	}

	fmt.Println(result)
}
