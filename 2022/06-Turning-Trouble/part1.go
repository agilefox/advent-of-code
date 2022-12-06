package main

import (
	"bufio"
	"fmt"
	"os"
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

func findStartOfPacket(line string) int {
	var i int
	for i = marker; i < len(line); i++ {
		var bad = false
		var four = line[i-marker : i]
		keys := make(map[rune]bool)
		for _, v := range four {
			if keys[v] {
				bad = true
				break
			}
			keys[v] = true
		}
		if !bad {
			return i
		}
	}
	return len(line)
}

func main() {

	lines := readFile("input.txt")

	// find it
	i := findStartOfPacket(lines[0])
	fmt.Println(i, string(lines[0][i]), lines[0][i-marker:i])
}

// RNZLFZSJH
