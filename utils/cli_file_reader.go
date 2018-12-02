package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MapToInts(strings []string) []int {
	integers := make([]int, 0)
	for _, str := range strings {
		integer, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		integers = append(integers, integer)
	}
	return integers
}

func ReadArgumentFile() []string {
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputData := make([]string, 0)

	for scanner.Scan() {
		inputData = append(inputData, scanner.Text())
	}
  return inputData
}
