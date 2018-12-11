package main

import (
	"fmt"
	utils "github.com/Mabrollin/AoC/utils"
	"strings"
)

type node struct {
	children []*node
	metadata []int
}

func createNode(input []int) (*node, []int) {
	node := &node{
		children: make([]*node, input[0]),
		metadata: make([]int, input[1]),
	}
	input = input[2:]
	for i, child := range node.children {
		child, input = createNode(input)
		node.children[i] = child
	}
	for i, _ := range node.metadata {
		node.metadata[i] = input[i]
	}
	input = input[len(node.metadata):]
	return node, input
}

func (n node) getMetaSum1() int {
	metaSum := 0
	for _, child := range n.children {
		metaSum += child.getMetaSum1()
	}
	for _, meta := range n.metadata {
		metaSum += meta
	}
	return metaSum
}

func (n node) getMetaSum2() int {
	metaSum := 0
	if len(n.children) > 0 {
		fmt.Println(n.metadata)
		for _, meta := range n.metadata {
			fmt.Println(len(n.children), meta)
			if len(n.children) >= meta && meta != 0{
				metaSum += n.children[meta-1].getMetaSum2()
			}
		}
	} else {
		for _, meta := range n.metadata {
			metaSum += meta
		}
	}
	return metaSum
}

func main() {
	input := utils.MapToInts(strings.Split(utils.ReadArgumentFile()[0], " "))
	node, _ := createNode(input)
	fmt.Println(node.getMetaSum1())
	fmt.Println(node.getMetaSum2())

}
