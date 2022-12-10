package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

var sum int

func cycleCheck(cycle int, x int) {
	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		sum += x * cycle
		log.Println(cycle, x, sum)
	}
}

func main() {

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)

	lines := readFile("input.txt")

	cycle := 1
	x := 1

	// nnop
	// addx 3
	// 1 -
	// 2 start addx
	// 3 addx running
	// 4 addx completed
	for _, line := range lines {
		cycle += 1 // nop time
		cycleCheck(cycle, x)
		// get current instruction
		var instructions = strings.Split(line, " ")

		if len(instructions) > 1 {
			var add, _ = strconv.Atoi(instructions[1])
			x += add
			cycle += 1
			cycleCheck(cycle, x)
		}

	}

	log.Println(sum)

}
