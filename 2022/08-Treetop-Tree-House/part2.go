package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const marker = 14

func readFile(filename string) []string {
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()
	return lines
}

type Tree struct {
	Height int
	Seen   bool
}

func getScenicScore(trees [][]Tree, n int, i int) int {
	const max = 99
	h := trees[n][i].Height
	cur_h := -1
	acc := 1
	// walk left
	local := 0
	for x := i - 1; cur_h < h && x >= 0 && x < max; x-- {
		local++
		cur_h = trees[n][x].Height
	}
	acc *= local
	// walk right
	local = 0
	cur_h = -1
	for x := i + 1; cur_h < h && x >= 0 && x < max; x++ {
		local++
		cur_h = trees[n][x].Height
	}
	acc *= local

	// walk up
	local = 0
	cur_h = -1
	for y := n + 1; cur_h < h && y >= 0 && y < max; y++ {
		local++
		cur_h = trees[y][i].Height
	}
	acc *= local

	// walk down
	local = 0
	cur_h = -1
	for y := n - 1; cur_h < h && y >= 0 && y < max; y-- {
		local++
		cur_h = trees[y][i].Height
	}
	acc *= local
	return acc
}

func main() {

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)

	lines := readFile("input.txt")
	fmt.Println(len(lines[0]), len(lines))
	// make a map of Trees
	trees := make([][]Tree, len(lines))

	// Fill in the map (and tranverse horizontally)
	for n, line := range lines {
		trees[n] = make([]Tree, len(line))
		for i, h := range line {
			// plant a tree
			height, _ := strconv.Atoi(string(h))
			trees[n][i] = Tree{height, false}
		}
	}

	highScore := 0
	for n, row := range lines {
		for i, _ := range row {
			score := getScenicScore(trees, n, i)
			if score > highScore {
				highScore = score
			}
		}
	}

	log.Println(highScore)
}
