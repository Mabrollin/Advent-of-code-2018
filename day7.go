package main

import (
	"fmt"
	utils "github.com/Mabrollin/AoC/utils"
)

type link struct {
	letter    byte
	leadsTo   []*link
	blockedBy []*link
}

func (l link) print() {
	fmt.Printf("\n\nNode: %c", l.letter)
	fmt.Print("\nblocked by: ")
	for _, blocked := range l.blockedBy {
		fmt.Printf("%c ", blocked.letter)
	}
	fmt.Print("\nleads to: ")
	for _, lead := range l.leadsTo {
		fmt.Printf("%c ", lead.letter)
	}
}

type work struct {
	link     *link
	progress int
}

func (w work) isDone() bool {
	return w.progress >= int(w.link.letter)-4

}

func main() {
	lines := utils.ReadArgumentFile()
	linkMap := make(map[byte]*link)
	for _, line := range lines {
		condition := line[5]
		leadsTo := line[36]
		if val1, ok := linkMap[condition]; ok {
			if val2, ok := linkMap[leadsTo]; ok {
				val1.leadsTo = append(val1.leadsTo, val2)
				val2.blockedBy = append(val2.blockedBy, val1)
				val1.print()
				val2.print()
			} else {
				newLink := link{
					letter:    leadsTo,
					leadsTo:   []*link{},
					blockedBy: []*link{},
				}
				val1.leadsTo = append(val1.leadsTo, &newLink)
				newLink.blockedBy = append(newLink.blockedBy, val1)
				linkMap[newLink.letter] = &newLink
				val1.print()
				newLink.print()
			}
		} else {
			if val2, ok := linkMap[leadsTo]; ok {
				newLink := link{
					letter:    condition,
					leadsTo:   []*link{},
					blockedBy: []*link{},
				}
				newLink.leadsTo = append(newLink.leadsTo, val2)
				val2.blockedBy = append(val2.blockedBy, &newLink)
				linkMap[newLink.letter] = &newLink
				newLink.print()
				val2.print()
			} else {
				newLink1 := link{
					letter:    condition,
					leadsTo:   []*link{},
					blockedBy: []*link{},
				}
				newLink2 := link{
					letter:    leadsTo,
					leadsTo:   []*link{},
					blockedBy: []*link{},
				}
				newLink1.leadsTo = append(newLink1.leadsTo, &newLink2)
				newLink2.blockedBy = append(newLink2.blockedBy, &newLink1)
				linkMap[newLink1.letter] = &newLink1
				linkMap[newLink2.letter] = &newLink2
				newLink1.print()
				newLink2.print()
			}
		}
	}
	currentNodes := []*link{}
	for _, link := range linkMap {
		if len(link.blockedBy) == 0 {
			currentNodes = append(currentNodes, link)
		}
	}

	output := ""

	for len(currentNodes) > 0 {
		lowest := currentNodes[0]
		lowestIndex := 0
		for i, node := range currentNodes {
			if node.letter < lowest.letter {
				lowest = node
				lowestIndex = i
			}
		}
		output += string([]byte{lowest.letter})

		currentNodes = currentNodes[:lowestIndex+copy(currentNodes[lowestIndex:], currentNodes[lowestIndex+1:])]
		for _, lead := range lowest.leadsTo {
			if len(lead.blockedBy) == 1 && lead.blockedBy[0].letter == lowest.letter {
				currentNodes = append(currentNodes, lead)
			} else {
				for i, val := range lead.blockedBy {
					if val == lowest {
						lead.blockedBy = lead.blockedBy[:i+copy(lead.blockedBy[i:], lead.blockedBy[i+1:])]
					}
				}
			}
		}
	}
	fmt.Println(output)

	// Part 2
	queue := []*work{}
	for _, link := range linkMap {
		if len(link.blockedBy) == 0 {
			work := work{
				link:     link,
				progress: 0,
			}
			queue = append(queue, &work)
		}
	}

	workers := [4]*work{}
	minute := 0
	for len(queue) != 0 || !idle(workers) {
		for hasFree(workers) && len(queue) != 0 {
			lowest := queue[0]
			lowestIndex := 0
			for i, work := range queue {
				if work.link.letter < lowest.link.letter {
					lowest = work
					lowestIndex = i
				}
			}

			worker := getFirstFree(workers)
			workers[worker] = lowest

			queue = queue[:lowestIndex+copy(queue[lowestIndex:], queue[lowestIndex+1:])]
		}
		minute++
		for i, worker := range workers {
			if worker == nil {
				continue
			}
			worker.progress++
			if worker.isDone() {
				for _, lead := range worker.link.leadsTo {
					if len(lead.blockedBy) == 1 && lead.blockedBy[0].letter == worker.link.letter {
						work := work{
							link:     lead,
							progress: 0,
						}
						queue = append(queue, &work)
					} else {
						for i, val := range lead.blockedBy {
							if val == worker.link {
								lead.blockedBy = lead.blockedBy[:i+copy(lead.blockedBy[i:], lead.blockedBy[i+1:])]
							}
						}
					}
				}
				workers[i] = nil
			}

		}
	}
	fmt.Println(minute)

}

func idle(workers [4]*work) bool {
	idle := true
	for _, worker := range workers {
		idle = idle && (worker == nil || worker.isDone())
	}
	return idle
}

func hasFree(workers [4]*work) bool {
	return getFirstFree(workers) != -1
}

func getFirstFree(workers [4]*work) int {
	for i, worker := range workers {
		if worker == nil || worker.isDone() {
			return i
		}
	}
	return -1
}
