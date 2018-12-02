package main

import (
	"fmt"
	utils "github.com/Mabrollin/AoC/utils"
)

func main() {
	frequencies := utils.MapToInts(utils.ReadArgumentFile())

	// Part 1
	totalFrequency := 0
	for _, frequency := range frequencies {
		totalFrequency += frequency
	}

	fmt.Println(totalFrequency)

	// Part 2
	totalFrequency = 0
	var first *int
	reachedFrequencies := make(map[int]bool)

	for first == nil {
		for _, frequency := range frequencies {
			totalFrequency += frequency
			if reachedFrequencies[totalFrequency] {
				if first == nil {
					firstValue := totalFrequency
					first = &firstValue
					break
				}
			}
			reachedFrequencies[totalFrequency] = true
		}
	}

	fmt.Println(*first)
}
