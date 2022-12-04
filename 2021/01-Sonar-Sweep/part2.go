package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var total = 0
	var depths []int
	for fileScanner.Scan() {
		d, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			fmt.Println("wtf")
			os.Exit(1)
		}
		depths = append(depths, d)
	}
	readFile.Close()
	for i := 3; i < len(depths); i++ {
		if depths[i-3]+depths[i-2]+depths[i-1] < depths[i-2]+depths[i-1]+depths[i] {
			total += 1
		}
	}

	fmt.Println("Final Score:", total)
}
