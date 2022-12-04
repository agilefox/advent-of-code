package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)
// A for Rock, B for Paper, and C for Scissors
// 1 for Rock, 2 for Paper, and 3 for Scissors
var throws = map[string]int{"A": 1,"B": 2,"C": 3}
var invthrows = map[int]string{1: "A",2: "B",3: "C"}

// X for lose, Y for draw, and Z for win
func getMyThrow(opp string, end string) string {
    //draw
    if (end == "Y") {
        return opp
    } 
    
    // win
    if (end == "Z") {
        return invthrows[(throws[opp] % 3)+1]
    }
    
    // lose
    /**
     * A=1=3=C
     * B=2=1=A
     * C=3=2=B
     * 1 -> 3
     * 2 -> 1
     * 3 -> 2
     * 0->2->2->3
     * 1->3->0->1
     * 2->4->1->2
     */
     c := ((throws[opp]+1) % 3) +1
    return invthrows[c] 
}
// 0 if you lost, 3 if the round was a draw, and 6 if you won
func getRoundScore(opp string, me string) int {
    // get val for each throw
    if throws[opp] == throws[me] {
        // fmt.Println(opp, "=", me)
        return 3 + throws[me]
    }
    // R v P, P v S, 
    if throws[opp] == (throws[me]-1) || (throws[opp] == 3 && throws[me] == 1) {
        // fmt.Println(opp, "<", me)
        return 6 + throws[me]
    }
    // fmt.Println(opp, ">", me)
    
    return throws[me]
}

func getThrows(line string) (string, string) {
    split := strings.Split(line, " ")
    return split[0],split[1]

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
    	throws := fileScanner.Text()
        opp, end := getThrows(throws)
        // X for lose, Y for draw, and Z for win
        // if end == "X" {
        //     fmt.Println("LOSE")
        // } else if end == "Y" {
        //     fmt.Println("DRAW")
        // } else {
        //     fmt.Println("WIN")
        // }
        me := getMyThrow(opp, end)
        total += getRoundScore(opp, me)
    }
    readFile.Close()
    fmt.Println("Final Score:", total)    	
}