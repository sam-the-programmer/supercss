package main

import (
	"log"
	"os"
	"strings"
)

// ReadFile returns a string of the read file.
func ReadFile(filepath string) string {
	bytes, err := os.ReadFile(filepath)

	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}

// WriteFile writes data to a file.
func WriteFile(filepath string, input string) {
	err := os.WriteFile(filepath, []byte(input), os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	files := []string{
		"bgs", "columns",
		"glow-text", "pad",
		"textcol", // base.css, bigger.css, hover.css and utils.css (along with any other ones desinged for readability for human editors) should not be added to this slice.
	}

	for _, v := range files {
		fileString := ReadFile("css/" + v + ".css")

		// Steps go here...
		fileString = strings.ReplaceAll(fileString, "\n", "")

		WriteFile("css/"+v+".css", fileString)
	}
}
