package main

import "fmt"

func main() {
	fmt.Println(transverse(Node{0, 0}, 0, 20))
}

type Node struct {
	x int
	y int
}

func transverse (currentNode Node, paths int, gridSize int) int {
	if (currentNode.x == gridSize && currentNode.y == gridSize) {
		//fmt.Println(paths+1)
		return paths + 1
	}
	//go down
	if currentNode.x < gridSize && currentNode.y < gridSize {
		return transverse(Node{currentNode.x+1, currentNode.y}, paths, gridSize) +
			transverse(Node{currentNode.x, currentNode.y+1}, paths, gridSize)
	}

	if  currentNode.x < gridSize {
		return transverse(Node{currentNode.x+1, currentNode.y}, paths, gridSize)
	}

	if currentNode.y < gridSize {
		return transverse(Node{currentNode.x, currentNode.y+1}, paths, gridSize)
	}

	//will not reach here
	return paths
}
