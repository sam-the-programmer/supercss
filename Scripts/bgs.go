package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const newline = ""

func floatToStr(f float64) string {
	return strings.ReplaceAll(fmt.Sprintf("%f", f/10), ".", "-")
}

func genBgs(output string) string {
	for _, v := range []string{
		"red", "orange", "yellow",
		"green", "lightblue", "blue",
		"purple", "pink", "gray",
		"black", "brown",
	} {
		output += newline + ".bg-" + string(v) + "{background-color:" + string(v) + ";}"
		for i := 1; i < 4; i++ {
			output += newline + ".bg-" + string(v) + strconv.Itoa(i) + "{background-color: var(--" + string(v) + strconv.Itoa(i) + ");}"
		}
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
	output := `@import url('./base.css');
.bg-white{background-color:white;}`

	output = genBgs(output)

	WriteFile(output, "./CSS/bgs.css")
}
