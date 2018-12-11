package main

import (
	"fmt"
	utils "github.com/Mabrollin/AoC/utils"
	"strconv"
	"strings"
)

type marble struct {
	left   *marble
	right  *marble
	number int
}

func (m marble) find(number int) *marble {
	if m.number == number {
		return &m
	}
	p := m.left
	for p != &m {
		if p.number == number {
			return p
		}
		p = p.left
	}
	return nil
}

func (m marble) print() {
	p := m.left
	for p.number != m.number {
		fmt.Printf("%d ", p.number)
		p = p.left
	}
	fmt.Println(m.number)
}

func (m marble) delete() {
	m.left.right = m.right
	m.right.left = m.left
//	fmt.Printf("wgaga: %d\n", m.number)
}
func (m marble) add(new marble) *marble {
	right := m.right.right
	left := m.right
	left.right = &new
	new.left = left
	right.left = &new
	new.right = right
	return &new
}

func (m marble) place(new marble) (int, *marble) {
	score := 0
	if new.number%23 == 0 {
		score += new.number
		seven := m.left.left.left.left.left.left.left
		score += seven.number
		seven.delete()
		return score, seven.right
	}
	return score, m.add(new)

}

func main() {
	input := utils.ReadArgumentFile()[0]
	splitInput := strings.Split(input, "; ")

	players, err := strconv.Atoi(strings.TrimSuffix(splitInput[0], " players"))
	if err != nil {
		panic(err)
	}

	finalScore, err := strconv.Atoi(strings.TrimPrefix(strings.TrimSuffix(splitInput[1], " points"), "last marble is worth "))
	if err != nil {
		panic(err)
	}
	scoreByPlayer := make([]int, players)

	current := &marble{}
	current.left = current
	current.right = current
	player := 0
currentNumber := 0

	for {
		//current.print()
		currentNumber ++
		next := marble{
			number: currentNumber,
		}
		score := 0
		score, current = current.place(next)
		scoreByPlayer[player] += score
		player++
		if player >= players {
			player = 0
		}
		if next.number == finalScore {
			break
		}
	}
	high := 0
for _, score := range scoreByPlayer {
	if score > high {
		high = score
	}
}
	fmt.Println(high)
}
