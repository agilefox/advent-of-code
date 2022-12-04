package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var x, y = 0, 0

	for fileScanner.Scan() {
		components := strings.Split(fileScanner.Text(), " ")
		magnitude, err := strconv.Atoi(components[1])
		if err != nil {
			fmt.Println("wtf")
			os.Exit(1)
		}
		switch components[0] {
		case "forward":
			x += magnitude
		case "down":
			y += magnitude
		case "up":
			y -= magnitude
		}
	}
	readFile.Close()

	fmt.Println("Final Score:", x*y)
}
