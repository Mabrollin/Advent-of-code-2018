package main

import (
	"fmt"
	utils "github.com/Mabrollin/AoC/utils"
	//"sort"
	"strings"
)

const EMPTY_CHAR_SLOT = " "

func main() {
	codes := utils.ReadArgumentFile()
	partCodeOccurrance := make(map[string]int)
	var similarCode *string

	for _, code := range codes {
		letters := strings.Split(code, "")

		for i := range letters {
			partOfCodeLetters := make([]string, len(letters))
			copy(partOfCodeLetters, letters)
			partOfCodeLetters[i] = EMPTY_CHAR_SLOT
			partCode := strings.Join(partOfCodeLetters, "")
			if _, ok := partCodeOccurrance[partCode]; ok {
				similarPartCode := trimPartCode(partCode)
				similarCode = &similarPartCode
			} else {
				partCodeOccurrance[partCode] = 1
			}
		}
	}
	fmt.Printf("similar code: %s\n", *similarCode)
}

func trimPartCode(partCode string) string {
	return strings.Replace(partCode, " ", "", -1)
}
