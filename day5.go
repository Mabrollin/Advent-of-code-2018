package main

import (
	"fmt"
	utils "github.com/Mabrollin/AoC/utils"
	"strings"
)

func main() {
	chain := strings.Split(utils.ReadArgumentFile()[0], "")
	fmt.Println(len(chain))

	// Part 1
	for i := 1; i < len(chain); i++ {
		if shouldReact(chain[i-1], chain[i]) {
			chain = append(chain[:i-1], chain[i+1:]...)
			i -= 3
			if i < 0 {
				i = 0
			}
		}
	}

	fmt.Println(len(chain))

	// part 2

	lowest := len(chain)

	for i := 0; i < len(chain); i++ {
		newChain := make([]string, len(chain))
		copy(newChain, chain)
		newChain = strings.Split(strings.Replace(strings.Replace(strings.Join(newChain, ""), strings.ToUpper(newChain[i]), "", -1), strings.ToLower(newChain[i]), "", -1), "")
		for i := 1; i < len(newChain); i++ {
			if shouldReact(newChain[i-1], newChain[i]) {
				newChain = append(newChain[:i-1], newChain[i+1:]...)
				i -= 3
				if i < 0 {
					i = 0
				}
			}
		}

		if len(newChain) < lowest {
			lowest = len(newChain)
		}
	}
	fmt.Println(lowest)
}

func shouldReact(first string, second string) bool {
	return (first != second) && (strings.ToUpper(first) == strings.ToUpper(second))
}
