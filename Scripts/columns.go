package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func floatToStr(f float64) string {
	return strings.ReplaceAll(fmt.Sprintf("%f", f/10), ".", "-")
}

func genColumns(output string) string {
	for i := 1; i < 10; i++ {
		output += "\n.columns-" + strconv.Itoa(i) + "{columns: " + strconv.Itoa(i) + ";}"
	}
	return output
}

// WriteFile writes a string to a file.
func WriteFile(output string, filename string) {
	err := os.WriteFile(filename, []byte(output), os.ModePerm)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	var output string

	output = genColumns(output)

	WriteFile(output, "./CSS/columns.css")
	output = ""
}
