package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func findDuplicateItem(line string) int {
    // 4 -> 2 
    a := line[0:len(line)/2]
    b := line[len(line)/2:len(line)]
    // search the first half of the string for each item in the second half
    for i := len(line)/2; i < len(line); i++ {
        val := rune(line[i])
        if strings.ContainsRune(a, val) {

            /** 
             * a = 97
             * A = 65
             * Lowercase item types a through z have priorities 1 through 26.
             * Uppercase item types A through Z have priorities 27 through 52.
             */ 
             ret := 0
            if int(val) >= 97 {
                ret = int(val) - 96
            } else {
                ret = int(val) - 38
            }

            fmt.Println(a,b,string(line[i]),val, ret)
            return ret
        }
    }
    
    return 0

}

func main() {
	readFile, err := os.Open("input.txt")
	 if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  
    var total = 0
    for fileScanner.Scan() {
    	contents := fileScanner.Text()
        total += findDuplicateItem(contents)

    }
    readFile.Close()
    fmt.Println("Final Score:", total)    	
}