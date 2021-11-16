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

func genGlows(output string) string {
	for _, v := range []string{
		"red", "orange", "yellow",
		"green", "lightblue", "blue",
		"purple", "pink", "gray",
		"black", "brown",
	} {
		output += newline + ".tglow-" + string(v) + "{text-shadow: 0 0 .5ch" + string(v) + ";}"
		output += newline + ".tdrop-" + string(v) + "{text-shadow: .2ch .2ch 0" + string(v) + ";}"
		for i := 1; i < 4; i++ {
			output += newline + ".tglow-" + string(v) + strconv.Itoa(i) + "{text-shadow: 0 0 .5ch" + string(v) + ";}"
			output += newline + ".tglow-" + string(v) + strconv.Itoa(i) + "{text-shadow: .2ch .2ch 0" + string(v) + ";}"
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
.tglow-white{text-shadow: 0 0 .5ch white;}.tdrop-whtie{text-shadow: .2ch .2ch 0 white;}`

	output = genGlows(output)

	WriteFile(output, "./CSS/glow-text.css")
}
