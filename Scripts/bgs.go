package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		output += "\n.bg-" + string(v) + "{background-color:" + string(v) + ";}"
		for i := 1; i < 4; i++ {
			output += "\n.bg-" + string(v) + strconv.Itoa(i) + "{background-color: var(--" + string(v) + strconv.Itoa(i) + ");}"
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

	output = gen_bgs(output)

	WriteFile(output, "./CSS/bgs.css")
}
