package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)
/** 
 * a = 97
 * A = 65
 * Lowercase item types a through z have priorities 1 through 26.
 * Uppercase item types A through Z have priorities 27 through 52.
 */ 
func getRuneVal(item int) int {
    if int(item) >= 97 {
        return int(item) - 96
    }
    return int(item) - 38
}
func findDuplicateItems(needle string, haystack string) string {
    // 4 -> 2     
    // search the first string for each item in the second half
    var result string
    for i := 0; i < len(needle); i++ {
        val := rune(needle[i])
        if strings.ContainsRune(haystack, val) && !strings.ContainsRune(result,val) {
            result += string(val)
        }
    }
    
    return result
}

func main() {
	readFile, err := os.Open("input.txt")
	 if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  
    var total = 0
    var lines []string
    for fileScanner.Scan() {
    	lines = append(lines,fileScanner.Text())
    }
    readFile.Close()
    var items string
    for i := 0; i < len(lines); i+=3 {
        items = findDuplicateItems(lines[i], lines[i+1])
        items = findDuplicateItems(items, lines[i+2])
        fmt.Println(lines[i], lines[i+1], lines[i+2], items)
        total += getRuneVal(int(items[0]))
    }
    
    
    fmt.Println("Final Score:", total)    	
}