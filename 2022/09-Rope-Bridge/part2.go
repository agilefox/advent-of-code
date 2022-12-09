package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

type Point struct {
	X int
	Y int
}

func getVector(dir string) Point {
	switch dir {
	case "U":
		return Point{0, 1}
	case "D":
		return Point{0, -1}
	case "L":
		return Point{-1, 0}
	}
	return Point{1, 0}
}

func add(a Point, b Point) Point {
	return Point{a.X + b.X, a.Y + b.Y}
}

func far(a Point, b Point) bool {
	if math.Abs(float64(a.X-b.X)) > 1 || math.Abs(float64(a.Y-b.Y)) > 1 {
		return true
	}
	return false
}

func stepToward(a Point, b Point) Point {
	// 3,5 = H-T = -2/abs(-2) = -1
	// 5,3 = H-T = 2/abs(2) = 1
	if a.X != b.X {
		d := a.X - b.X
		b.X += d / int(math.Abs(float64(d)))
	}
	if a.Y != b.Y {
		d := a.Y - b.Y
		b.Y += d / int(math.Abs(float64(d)))
	}
	return Point{b.X, b.Y}
}

func main() {

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)

	lines := readFile("input.txt")
	const head = 0
	const tail = 9
	const rope_length = 10
	var rope [rope_length]Point
	visited := make(map[Point]bool)
	visited[rope[tail]] = true // started at 0,0
	for _, cmd := range lines {
		var velocity = strings.Split(cmd, " ")
		var vector = getVector(velocity[0])
		var speed, _ = strconv.Atoi(velocity[1])
		for i := 0; i < speed; i++ {
			rope[head] = add(rope[head], vector)
			for i := head + 1; i < rope_length; i++ {
				if far(rope[i-1], rope[i]) {
					rope[i] = stepToward(rope[i-1], rope[i])
				}
			}
			visited[rope[tail]] = true
		}
	}

	log.Println(len(visited))

}
