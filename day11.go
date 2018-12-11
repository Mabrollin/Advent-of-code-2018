package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const GRID_SIZE = 300

func main() {
	serial, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	fixedGrid := true
	if strings.HasPrefix(os.Args[2], "fixedGrid=") {
		fixedGrid, err = strconv.ParseBool(strings.TrimPrefix(os.Args[2], "fixedGrid="))
		if err != nil {
			panic(err)
		}
	}

	grid := [300][300]int{}
	for x := 0; x < GRID_SIZE; x++ {
		for y := 0; y < GRID_SIZE; y++ {
			grid[x][y] = generateFuelLevel(x, y, serial)
		}
	}
	bestX := 0
	bestY := 0
	bestSize := 0
	best := 0
	foundScores := make(map[int]int)

	startSize := 2
	endSize := 3
	if !fixedGrid {
		startSize = 0
		endSize = GRID_SIZE - 1
	}
	for size := startSize; size < endSize; size++ {
		fmt.Printf("Progress: %d/%d\r", size+1, endSize)
		for x := 0; x < GRID_SIZE-size; x++ {
			for y := 0; y < GRID_SIZE-size; y++ {
				totalPower := 0
				if val, ok := foundScores[hash3(x, y, size-1)]; ok {
					totalPower += val
					for a := 0; a <= size; a++ {
						totalPower += grid[x+a][y+size]
					}
					for b := 0; b <= size-1; b++ {
						totalPower += grid[x+size][y+b]
					}
				} else {
					for a := 0; a <= size; a++ {
						for b := 0; b <= size; b++ {
							totalPower += grid[x+a][y+b]
						}
					}
				}
				foundScores[hash3(x, y, size)] = totalPower
				if totalPower > best {
					best = totalPower
					bestX = x
					bestY = y
					bestSize = size
				}
			}
		}
	}
	fmt.Printf("\nbest: (%d,%d,%d) with the score of %d\n", bestX+1, bestY+1, bestSize+1, best)
}

func generateFuelLevel(x, y, serial int) int {
	rackId := (x + 1) + 10
	power := (y + 1) * rackId
	power += serial
	power *= rackId
	power = (power - (power/1000)*1000 - power%100) / 100
	power -= 5
	return power
}

func hash2(x, y int) int {
	x += 600000
	y += 600000
	return (x+y)*(x+y+1)/2 + x
}

func hash3(x, y, z int) int {
	return hash2(hash2(x, y), z)
}
