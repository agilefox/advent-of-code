package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
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

func gatherDocs(lines []string) ([]string, []string) {
	var manifest, instructions []string
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			manifest = lines[:i-1]
			instructions = lines[i+1:]
			break
		}
	}
	return manifest, instructions
}

func convertManifest(manifest []string) []string {
	containers := make([]string, 9)

	for _, line := range manifest {
		var stack = 0
		for i := 1; i < len(line); i += 4 {
			if string(line[i]) != " " {
				containers[stack] = containers[stack] + string(line[i])
				// fmt.Printf("Adding %s to %d\n", string(line[i]), stack)
			}
			stack += 1
		}
	}
	return containers
}

func parseInstuctions(line string) (int, int, int) {
	num := false
	var s string
	var vals []string
	for _, c := range line {

		if unicode.IsNumber(c) {
			num = true
			s = s + string(c)
		} else if num {
			num = false
			vals = append(vals, s)
			s = ""
		}
	}
	x, _ := strconv.Atoi(vals[0])
	y, _ := strconv.Atoi(vals[1])
	z, _ := strconv.Atoi(s)
	return x, y, z
}
func reverse(s string) string {
	var r string
	for _, c := range s {
		r = string(c) + r
	}
	return r
}
func moveCrates(containers []string, num int, src int, dst int) {
	fmt.Printf("Move %d from %d to %d\n", num, src+1, dst+1)
	stack := containers[src][:num]
	// REVERSE STACK!
	//stack = reverse(stack)
	fmt.Println(containers[src], fmt.Sprintf("[%s]", stack), "->", containers[dst])
	containers[dst] = stack + containers[dst]
	containers[src] = containers[src][num:]
	fmt.Println(containers[src], containers[dst])
}

func main() {
	lines := readFile("input.txt")

	// split into 2 arrays
	manifest, instructions := gatherDocs(lines)

	// convert manifest to 9 slices
	containers := convertManifest(manifest)

	fmt.Println("Containers", containers)
	// carry out instuctions
	for _, line := range instructions {
		x, y, z := parseInstuctions(line)
		moveCrates(containers, x, y-1, z-1)
	}

	// fmt.Println("Manifest", manifest[0])
	// fmt.Println("Instructions", instructions[0])
	fmt.Println("Containers", containers)
}

// RNZLFZSJH
// CNSFCGJSM
