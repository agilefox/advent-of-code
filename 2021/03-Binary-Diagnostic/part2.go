package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func reduce(lines []int, i int, side bool) int {
	fmt.Println("start", lines[0], len(lines), i)
	// get the needle, this is the existing size + our bit. So if we look at the first item, it will be AT LEAST this size
	// find the location of the first item with our bit
	needle := int(math.Pow(2, float64(11-i)))
	for n := 0; n < i; n++ {
		fmt.Println("n", int(math.Pow(2, float64(11-n))))
		needle += (int(math.Pow(2, float64(11-n))) & lines[0])
		fmt.Println("n2", needle)
	}

	index := sort.Search(len(lines), func(j int) bool { return lines[j] >= needle })
	fmt.Println(fmt.Sprintf("%b", needle), needle, index, len(lines)/2, lines[index])
	// split the array at this location
	if side {
		// more 1s
		if index <= int(math.Ceil(float64(len(lines))/2.0)) {
			// keep the big numbers
			lines = lines[index:]
		} else {
			// keep the small numbers
			lines = lines[:index]
		}
	} else {
		// more 0s
		if index >= int(math.Floor(float64(len(lines))/2.0)) {
			// keep the lil numbers
			lines = lines[0:index]
		} else {
			lines = lines[index:]
		}
	}
	if len(lines) == 1 {
		return lines[0]
	}
	return reduce(lines, i+1, side)
}

// too low 2054272
// 2089472
func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []int
	for fileScanner.Scan() {
		val, _ := strconv.ParseInt(fileScanner.Text(), 2, 64)
		lines = append(lines, int(val))
	}
	readFile.Close()
	lines2 := make([]int, len(lines))
	sort.Ints(lines)
	copy(lines2, lines)
	ox := reduce(lines, 0, true)
	fmt.Println("oxygen generator rating", ox)
	//co := reduce(lines, 0, false)
	//fmt.Println("CO2 scrubber rating", co, co*ox)

}
