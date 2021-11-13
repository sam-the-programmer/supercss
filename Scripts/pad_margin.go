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

func gen_pad_margin(output string) string {
	output += "\n" + ".p-tenth{padding: 0.1vh;}"
	output += "\n" + ".p-fifth{padding: 0.2vh;}"
	output += "\n" + ".p-half{padding: 0.5vh;}"

	output += "\n" + ".m-tenth{margin: 0.1vh;}"
	output += "\n" + ".m-fifth{margin: 0.2vh;}"
	output += "\n" + ".m-half{margin: 0.5vh;}"

	for i := 0; i < 6; i++ {
		output += "\n" + ".p-" + strconv.Itoa(i) + "{padding: " + strconv.Itoa(i) + "vh;}"
		output += "\n" + ".m-" + strconv.Itoa(i) + "{margin: " + strconv.Itoa(i) + "vh;}"
	}

	for i := 6; i < 20; i = i + 2 {
		output += "\n" + ".p-" + strconv.Itoa(i) + "{padding: " + strconv.Itoa(i) + "vh;}"
		output += "\n" + ".m-" + strconv.Itoa(i) + "{margin: " + strconv.Itoa(i) + "vh;}"
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

	output = gen_pad_margin(output)

	WriteFile(output, "./CSS/pad.css")
	output = ""
}
