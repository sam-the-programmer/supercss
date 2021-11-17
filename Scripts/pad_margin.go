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

const newline = "" // Change to "\n" for formatting, leave empty for smaller file sizes.

func floatToStr(f float64) string {
	return strings.ReplaceAll(fmt.Sprintf("%f", f/10), ".", "-")
}

func genPadMargin(output string) string {
	output += newline + ".p-tenth{padding: .1vh;}"
	output += newline + ".pb-tenth{padding-bottom: .1vh;}"
	output += newline + ".pt-tenth{padding-top: .1vh;}"
	output += newline + ".pl-tenth{padding-left: .1vh;}"
	output += newline + ".pr-tenth{padding-right: .1vh;}"

	output += newline + ".m-tenth{margin: .1vh;}"
	output += newline + ".mb-tenth{margin-bottom: .1vh;}"
	output += newline + ".mt-tenth{margin-top: .1vh;}"
	output += newline + ".ml-tenth{margin-left: .1vh;}"
	output += newline + ".mr-tenth{margin-right: .1vh;}"

	output += newline + ".p-fifth{padding: .2vh;}"
	output += newline + ".pb-fifth{padding-bottom: .2vh;}"
	output += newline + ".pt-fifth{padding-top: .2vh;}"
	output += newline + ".pl-fifth{padding-left: .2vh;}"
	output += newline + ".pr-fifth{padding-right: .2vh;}"

	output += newline + ".m-fifth{margin: .2vh;}"
	output += newline + ".mb-fifth{margin-bottom: .2vh;}"
	output += newline + ".mt-fifth{margin-top: .2vh;}"
	output += newline + ".ml-fifth{margin-left: .2vh;}"
	output += newline + ".mr-fifth{margin-right: .2vh;}"

	output += newline + ".p-half{padding: .5vh;}"
	output += newline + ".pb-half{padding-bottom: .5vh;}"
	output += newline + ".pt-half{padding-top: .5vh;}"
	output += newline + ".pl-half{padding-left: .5vh;}"
	output += newline + ".pr-half{padding-right: .5vh;}"

	output += newline + ".m-half{margin: .5vh;}"
	output += newline + ".mb-half{margin-bottom: .5vh;}"
	output += newline + ".mt-half{margin-top: .5vh;}"
	output += newline + ".ml-half{margin-left: .5vh;}"
	output += newline + ".mr-half{margin-right: .5vh;}"

	for i := 0; i < 6; i++ {
		output = _gen(output, 9)
	}

	for i := 6; i < 20; i = i + 2 {
		output = _gen(output, i)
	}

	for i := 50; i < 101; i = i + 10 {
		output = _gen(output, i)
	}

	return output
}

func _gen(output string, i int) string {
	output += newline + ".p-" + strconv.Itoa(i) + "{padding: " + strconv.Itoa(i) + "vh;}"
	output += newline + ".pb-" + strconv.Itoa(i) + "{padding-bottom: " + strconv.Itoa(i) + "vh;}"
	output += newline + ".pt-" + strconv.Itoa(i) + "{padding-top: " + strconv.Itoa(i) + "vh;}"
	output += newline + ".pl-" + strconv.Itoa(i) + "{padding-left: " + strconv.Itoa(i) + "vh;}"
	output += newline + ".pr-" + strconv.Itoa(i) + "{padding-right: " + strconv.Itoa(i) + "vh;}"

	output += newline + ".m-" + strconv.Itoa(i) + "{margin: " + strconv.Itoa(i) + "vh;}"
	output += newline + ".mb-" + strconv.Itoa(i) + "{margin-bottom: " + strconv.Itoa(i) + "vh;}"
	output += newline + ".mt-" + strconv.Itoa(i) + "{margin-top: " + strconv.Itoa(i) + "vh;}"
	output += newline + ".ml-" + strconv.Itoa(i) + "{margin-left: " + strconv.Itoa(i) + "vh;}"
	output += newline + ".mr-" + strconv.Itoa(i) + "{margin-right: " + strconv.Itoa(i) + "vh;}"
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

	output = genPadMargin(output)

	WriteFile(output, "./CSS/pad.css")
	output = ""
}
