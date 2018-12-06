package main

import (
	"fmt"
	utils "github.com/Mabrollin/AoC/utils"
	"math"
	"strconv"
	"strings"
)

func main() {
	coors := utils.ReadArgumentFile()

	// Part 1
	coorMap := make(map[int]map[int]int)
	scoreMap := make(map[int]int)

	scoreMap[-1] = math.MinInt32

	xMax := 0
	yMax := 0
	for i, coor := range coors {
		x, y := parseCoor(coor)
		if _, ok := coorMap[x]; !ok {
			coorMap[x] = make(map[int]int)
		}
		coorMap[x][y] = i
		scoreMap[i] = 0
		if x > xMax {
			xMax = x
		}
		if y > yMax {
			yMax = y
		}
	}

	xMax = 400
	yMax = 400

	rangeMap := make([][]int, xMax)
safe := 0
	for x := 0; x < xMax; x++ {
		rangeMap[x] = make([]int, yMax)
		for y := 0; y < yMax; y++ {
			 close, total := getClosest(x, y, coorMap)
			 rangeMap[x][y] = close
			 if total < 10000 {
				 safe ++
			 }
			if x == 0 || x == (xMax-1) || y == 0 || y == (yMax-1) {
				scoreMap[rangeMap[x][y]] = math.MaxInt32
			} else {
				scoreMap[rangeMap[x][y]]++
			}
		}
	}
	fmt.Printf("safe: %d", safe)
	scoreMap = reduce(scoreMap)
}

func parseCoor(input string) (int, int) {
	splitInput := strings.Split(input, ", ")
	x, err := strconv.Atoi(splitInput[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(splitInput[1])
	if err != nil {
		panic(err)
	}
	return x, y
}

func getClosest(x int, y int, coorMap map[int]map[int]int) (int, int) {
	closestId := -1
	total := 0
	closestDistance := math.MaxInt32
	for xCoor, xMap := range coorMap {
		for yCoor, id := range xMap {
			distance := abs(x-xCoor) + abs(y-yCoor)
			total += distance
			if distance == closestDistance {
				closestId = -1
			}
			if distance < closestDistance {
				closestDistance = distance
				closestId = id
			}
		}
	}
	return closestId, total
}

func reduce(scoreMap map[int]int) map[int]int {
	for id, score := range scoreMap {
		fmt.Printf("\nid: %d with score: %d", id, score)
	}
	return nil
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
