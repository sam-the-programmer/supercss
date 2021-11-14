package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const NEWLINE = ""

func float_to_str(f float64) string {
	return strings.ReplaceAll(fmt.Sprintf("%f", f/10), ".", "-")
}

func gen_bgs(output string) string {
	for _, v := range []string{
		"red", "orange", "yellow",
		"green", "lightblue", "blue",
		"purple", "pink", "gray",
		"black", "brown",
	} {
		output += NEWLINE + ".t-" + string(v) + "{color:" + string(v) + ";}.t-" + string(v) + " * {color: " + string(v) + ";}"
		output += NEWLINE + ".t-" + string(v) + "{color:" + string(v) + ";}.t-" + string(v) + " * {color: " + string(v) + ";}"
		for i := 1; i < 4; i++ {
			output += NEWLINE + ".t-" + string(v) + strconv.Itoa(i) + "{color: var(--" + string(v) + strconv.Itoa(i) + ");}.t-" + string(v) + strconv.Itoa(i) + " * {color: var(--" + string(v) + strconv.Itoa(i) + ");}"
			output += NEWLINE + ".t-" + string(v) + strconv.Itoa(i) + "{text-shadow: var(--" + string(v) + strconv.Itoa(i) + ");}.t-" + string(v) + strconv.Itoa(i) + " * {color: var(--" + string(v) + strconv.Itoa(i) + ");}"
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
.t-white{color:white;}.t-white * {color:white;}`

	output = gen_bgs(output)

	WriteFile(output, "./CSS/textcol.css")
}
