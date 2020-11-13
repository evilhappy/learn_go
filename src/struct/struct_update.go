package main

import "fmt"

type Tree struct {
	nodeMap map[int]Node
	left int
	right int
}

type Node struct {
	left int
	right int
}

func main()  {
	var node Node
	node.left = 1
	node.right = 2

	nodeMap := make(map[int]Node)
	nodeMap[0] = node

	var tree Tree
	tree.nodeMap = nodeMap
	tree.left = 3
	tree.right = 4

	tree.left = 5
	treeNode := tree.nodeMap[0]
	treeNode.left = 6
	treeNode.right = 7
	fmt.Println(tree)
}
