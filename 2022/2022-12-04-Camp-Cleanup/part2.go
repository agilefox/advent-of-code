package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func getIntArray(r string) []int {
    x := strings.Split(r,"-")
    ret := []int{}
    for _, i := range x {
        j, err := strconv.Atoi(i)
        if err != nil {
            fmt.Println(err)
        }
        ret = append(ret, j)
    }
    return ret
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
        line := fileScanner.Text()
        pair := strings.Split(line, ",")
        
        first := getIntArray(pair[0])
        second := getIntArray(pair[1])
        
        

        // if the first number is inside
        if second[0] >= first[0] && second[0] <= first[1] {
            total += 1
            fmt.Println(second[0], "in", first)
        } else if second[1] >= first[0] && second[1] <= first[1] {
            total += 1
            fmt.Println(second[1], "in", first)
        } else if first[0] >= second[0] && first[0] <= second[1] {
            total += 1
            fmt.Println(first[0], "ins", second)
        } else if first[1] >= second[0] && first[1] <= second[1] {
            total += 1
            fmt.Println(first[1], "ins", second)
        } else {
            fmt.Println(first, "X", second)
        }
    }
    readFile.Close()    
    fmt.Println("Final Score:", total)      
}