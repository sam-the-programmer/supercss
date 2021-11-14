/*
This script generates the padding and the margins for the
framework.
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const NEWLINE = "" // Change to "\n" for formatting, leave empty for smaller file sizes.

func float_to_str(f float64) string {
	return strings.ReplaceAll(fmt.Sprintf("%f", f/10), ".", "-")
}

func gen_pad_margin(output string) string {
	output += NEWLINE + ".p-tenth{padding: .1vh;}"
	output += NEWLINE + ".pb-tenth{padding-bottom: .1vh;}"
	output += NEWLINE + ".pt-tenth{padding-top: .1vh;}"
	output += NEWLINE + ".pl-tenth{padding-left: .1vh;}"
	output += NEWLINE + ".pr-tenth{padding-right: .1vh;}"

	output += NEWLINE + ".m-tenth{margin: .1vh;}"
	output += NEWLINE + ".mb-tenth{margin-bottom: .1vh;}"
	output += NEWLINE + ".mt-tenth{margin-top: .1vh;}"
	output += NEWLINE + ".ml-tenth{margin-left: .1vh;}"
	output += NEWLINE + ".mr-tenth{margin-right: .1vh;}"

	output += NEWLINE + ".p-fifth{padding: .2vh;}"
	output += NEWLINE + ".pb-fifth{padding-bottom: .2vh;}"
	output += NEWLINE + ".pt-fifth{padding-top: .2vh;}"
	output += NEWLINE + ".pl-fifth{padding-left: .2vh;}"
	output += NEWLINE + ".pr-fifth{padding-right: .2vh;}"

	output += NEWLINE + ".m-fifth{margin: .2vh;}"
	output += NEWLINE + ".mb-fifth{margin-bottom: .2vh;}"
	output += NEWLINE + ".mt-fifth{margin-top: .2vh;}"
	output += NEWLINE + ".ml-fifth{margin-left: .2vh;}"
	output += NEWLINE + ".mr-fifth{margin-right: .2vh;}"

	output += NEWLINE + ".p-half{padding: .5vh;}"
	output += NEWLINE + ".pb-half{padding-bottom: .5vh;}"
	output += NEWLINE + ".pt-half{padding-top: .5vh;}"
	output += NEWLINE + ".pl-half{padding-left: .5vh;}"
	output += NEWLINE + ".pr-half{padding-right: .5vh;}"

	output += NEWLINE + ".m-half{margin: .5vh;}"
	output += NEWLINE + ".mb-half{margin-bottom: .5vh;}"
	output += NEWLINE + ".mt-half{margin-top: .5vh;}"
	output += NEWLINE + ".ml-half{margin-left: .5vh;}"
	output += NEWLINE + ".mr-half{margin-right: .5vh;}"

	for i := 0; i < 6; i++ {
		output += NEWLINE + ".p-" + strconv.Itoa(i) + "{padding: " + strconv.Itoa(i) + "vh;}"
		output += NEWLINE + ".pb-" + strconv.Itoa(i) + "{padding-bottom: " + strconv.Itoa(i) + "vh;}"
		output += NEWLINE + ".pt-" + strconv.Itoa(i) + "{padding-top: " + strconv.Itoa(i) + "vh;}"
		output += NEWLINE + ".pl-" + strconv.Itoa(i) + "{padding-left: " + strconv.Itoa(i) + "vh;}"
		output += NEWLINE + ".pr-" + strconv.Itoa(i) + "{padding-right: " + strconv.Itoa(i) + "vh;}"

		output += NEWLINE + ".m-" + strconv.Itoa(i) + "{margin: " + strconv.Itoa(i) + "vh;}"
		output += NEWLINE + ".mb-" + strconv.Itoa(i) + "{margin-bottom: " + strconv.Itoa(i) + "vh;}"
		output += NEWLINE + ".mt-" + strconv.Itoa(i) + "{margin-top: " + strconv.Itoa(i) + "vh;}"
		output += NEWLINE + ".ml-" + strconv.Itoa(i) + "{margin-left: " + strconv.Itoa(i) + "vh;}"
		output += NEWLINE + ".mr-" + strconv.Itoa(i) + "{margin-right: " + strconv.Itoa(i) + "vh;}"
	}

	for i := 6; i < 20; i = i + 2 {
		output += NEWLINE + ".p-" + strconv.Itoa(i) + "{padding: " + strconv.Itoa(i) + "vh;}"
		output += NEWLINE + ".pb-" + strconv.Itoa(i) + "{padding-bottom: " + strconv.Itoa(i) + "vh;}"
		output += NEWLINE + ".pt-" + strconv.Itoa(i) + "{padding-top: " + strconv.Itoa(i) + "vh;}"
		output += NEWLINE + ".pl-" + strconv.Itoa(i) + "{padding-left: " + strconv.Itoa(i) + "vh;}"
		output += NEWLINE + ".pr-" + strconv.Itoa(i) + "{padding-right: " + strconv.Itoa(i) + "vh;}"

		output += NEWLINE + ".m-" + strconv.Itoa(i) + "{margin: " + strconv.Itoa(i) + "vh;}"
		output += NEWLINE + ".mb-" + strconv.Itoa(i) + "{margin-bottom: " + strconv.Itoa(i) + "vh;}"
		output += NEWLINE + ".mt-" + strconv.Itoa(i) + "{margin-top: " + strconv.Itoa(i) + "vh;}"
		output += NEWLINE + ".ml-" + strconv.Itoa(i) + "{margin-left: " + strconv.Itoa(i) + "vh;}"
		output += NEWLINE + ".mr-" + strconv.Itoa(i) + "{margin-right: " + strconv.Itoa(i) + "vh;}"
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
