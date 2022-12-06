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

	var tally [12]int
	linecount := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		linecount += 1
		for i := 0; i < len(line); i++ {
			if string(line[i]) == "1" {
				tally[i] += 1
			}
		}
	}
	readFile.Close()
	var gamma, epsilon string
	half := linecount / 2
	for i := 0; i < 12; i++ {
		if tally[i] >= half {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gma, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	} else {
		fmt.Println(gma)
	}
	epi, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	} else {
		fmt.Println(epi)
	}
	fmt.Println(linecount, tally[0], tally, gamma, epsilon, gma, epi, gma*epi)
}
