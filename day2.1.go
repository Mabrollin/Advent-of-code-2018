package main

import (
	"fmt"
	utils "github.com/Mabrollin/AoC/utils"
	"strings"
)

func main() {
	codes := utils.ReadArgumentFile()
	doubles := 0
	tripples := 0

	for _, code := range codes {
		doubleAdd, trippleAdd := findDoublesAndTripples(code)
		doubles += doubleAdd
		tripples += trippleAdd
	}

	fmt.Printf("doubles: %d, tripples: %d, checksum: %d\n", doubles, tripples, doubles*tripples)

}

func findDoublesAndTripples(code string) (int, int) {
	letters := strings.Split(code, "")
	letterOccurrance := make(map[string]int)
	doubles := 0
	tripples := 0

	for _, letter := range letters {
		if _, ok := letterOccurrance[letter]; ok {
			letterOccurrance[letter] += 1
		} else {
			letterOccurrance[letter] = 1
		}

		if letterOccurrance[letter] == 2 {
			doubles++
		}
		if letterOccurrance[letter] == 3 {
			doubles--
			tripples++
		}
		if letterOccurrance[letter] == 4 {
			tripples--
		}
	}

	if doubles > 0 {
		doubles = 1
	}
	if tripples > 0 {
		tripples = 1
	}

	return doubles, tripples
}
