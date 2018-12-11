package main

import (
	"fmt"
	utils "github.com/Mabrollin/AoC/utils"
	"strconv"
	"strings"
)

type point struct {
	x  int
	y  int
	vx int
	vy int
}

func (p *point) move(time int) {
	p.x += p.vx * time
	p.y += p.vy * time
}
func (p point) hash() int {
	return hash(p.x, p.y)
}

func hash(x, y int) int {
	x += 600000
	y += 600000
	return (x+y)*(x+y+1)/2 + x
}

func main() {
	rawPoints := utils.ReadArgumentFile()
	points := make([]*point, len(rawPoints))
	for i, rawPoint := range rawPoints {
		splitInput := strings.Split(rawPoint, " velocity=")
		rawPosition := strings.TrimPrefix(splitInput[0], "position=")
		rawVelocity := splitInput[1]
		x, y := parseRawVector(rawPosition)
		vx, vy := parseRawVector(rawVelocity)
		points[i] = &point{
			x:  x,
			y:  y,
			vx: vx,
			vy: vy,
		}
	}

	for i := 0; i < 50000; i++ {
		pointsByCoor := make(map[int]*point)
		hasLine := false
		yStart := 0
		yEnd := 0
		xStart := 0
		for _, point := range points {
			point.move(1)
			pointsByCoor[hash(point.x, point.y)] = point
			if !hasLine {
				hasLine, yStart, yEnd = isInLine(point, pointsByCoor, 8)
				xStart = point.x
			}

		}
		if hasLine {
			fmt.Println(yStart, yEnd)
			fmt.Println("len ", len(points), len(pointsByCoor))
			cloud := [160][16]bool{}
			for x := 0; x <160; x++ {
				for y := yStart; y < yEnd+8; y++ {
					key := hash(x-20+xStart-32, y)
					_, ok := pointsByCoor[key]
					if ok {
						fmt.Println("key", key, x, y)
					}
					cloud[x][y-yStart] = ok
				}
			}
			fmt.Println("len ", len(points), len(pointsByCoor))
			fmt.Println(i)
			print(cloud)
		}

	}
}

func parseRawVector(raw string) (int, int) {
	raw = strings.TrimPrefix(raw, "<")
	raw = strings.TrimSuffix(raw, ">")
	rawInts := strings.Split(raw, ", ")
	a, err := strconv.Atoi(strings.TrimPrefix(rawInts[0], " "))
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(strings.TrimPrefix(rawInts[1], " "))
	if err != nil {
		panic(err)
	}
	return a, b
}

func print(grid [160][16]bool) {
	for _, row := range grid {
		for _, slot := range row {
			if slot {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func isInLine(point *point, pointsByCoor map[int]*point, length int) (bool, int, int) {
	yStart := point.y
	yEnd := point.y
	x := point.x
	for {
		if _, ok := pointsByCoor[hash(x, yStart)]; !ok || yEnd-yStart >= length {
			break
		}
		yStart--
	}
	for {
		if _, ok := pointsByCoor[hash(x, yEnd)]; !ok|| yEnd-yStart >= length {
			break
		}
		yEnd++
	}
	isInLine := yEnd-yStart >= length
	return isInLine, yStart, yEnd
}
