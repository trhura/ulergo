
package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main () {
	fmt.Println()
	graph := loadGraph("18.txt")

	for cRow, row := range graph {
		for cCol, _ := range row {
			parents := getParents(graph, cRow, cCol)
			if len(parents) > 0 {
				graph[cRow][cCol] += findMax(parents)
			}
		}
	}

	fmt.Println("The maximum total is ", findMax(graph[len(graph)-1]))
}

func findMax (numbers []int) int {
	if len(numbers) < 1 {
		panic("findMax received an empty slice.")
	}

	max := numbers[0]
	for _, n := range numbers {
		if n > max { max = n}
	}

	return max
}

func getParents (graph [][]int, curRow int, curCol int) []int{
	if (curRow == 0) { return []int{} }
	prevRow := graph[curRow-1]

	if curCol > 0 && curCol < (len(graph[curRow]) - 1) {
		return []int {prevRow[curCol-1], prevRow[curCol]}
	}

	if curCol > 0 {
		return []int {prevRow[curCol-1]}
	}

	if curCol < (len(graph[curRow]) - 1) {
		return []int {prevRow[curCol]}

	}

	panic("Shouldn't reach here.")
}

func loadGraph (filename string) [][]int {
	f, err := os.Open(filename)
	defer f.Close()
	check(err)

	lineScanner := bufio.NewScanner(f)
	lineScanner.Split(bufio.ScanLines)

	graph := make([][]int, 0)
	for lineScanner.Scan() {
		wordScanner := bufio.NewScanner(strings.NewReader(lineScanner.Text()))
		wordScanner.Split(bufio.ScanWords)
		nslice := make([]int, 0)

		for wordScanner.Scan() {
			n, _ := strconv.Atoi(wordScanner.Text())
			nslice = append(nslice, n)
		}

		graph = append(graph, nslice)
	}

	return graph
}

func check(e error) {
	if e != nil {
        panic(e)
    }
}
