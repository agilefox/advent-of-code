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

var sum = 0

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
		leftHeight := -1
		for i, h := range line {
			// plant a tree
			height, _ := strconv.Atoi(string(h))
			trees[n][i] = Tree{height, false}

			if height > leftHeight {
				leftHeight = height
				trees[n][i].Seen = true
			}

		}
		rightHeight := -1
		for i := len(lines) - 1; i >= 0; i-- {
			if trees[n][i].Height > rightHeight {
				rightHeight = trees[n][i].Height
				trees[n][i].Seen = true
			}
		}
	}

	for i := 0; i < len(trees[0]); i++ {
		// traverse U/D
		topHeight := -1
		for n := 0; n < len(trees); n++ {
			if trees[n][i].Height > topHeight {
				topHeight = trees[n][i].Height
				trees[n][i].Seen = true
			}
		}

		// traverse D/U
		bottomHeight := -1
		for n := len(trees) - 1; n >= 0; n-- {
			if trees[n][i].Height > bottomHeight {
				bottomHeight = trees[n][i].Height
				trees[n][i].Seen = true
			}
		}
	}
	sum = 0
	for _, row := range trees {
		for _, tree := range row {
			bit := 0
			if tree.Seen {
				sum++
				bit = 1
			}
			fmt.Print(bit)
		}
		fmt.Printf("\n")
	}
	// 1253 too low
	// 1719
	log.Println(sum)
}
