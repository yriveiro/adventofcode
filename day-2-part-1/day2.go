package main

import (
	"bufio"
	"fmt"
	"os"
)

// answers:
// A -> Rocks
// B -> Paper
// C -> sccisors
//
// responses:
// X -> Rocks    - 1 point
// Y -> Paper    - 2 points
// Z -> sccisors - 3 points

func main() {
  score := 0
	outcomes := map[string]int{
		"A X": 4, //  D 3 + 1 
		"A Y": 8, //  W 6 + 2 
		"A Z": 3, //  L 0 + 3 
		
    "B X": 1, //  L 0 + 1 
		"B Y": 5, //  D 3 + 2 
		"B Z": 9, //  W 6 + 3 

		"C X": 7, // W 6 + 1
		"C Y": 2, // L 0 + 2
		"C Z": 6, // D 3 + 3

	}

	f, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	fs := bufio.NewScanner(f)

	fs.Split(bufio.ScanLines)

	for fs.Scan() {
    fmt.Printf("Round: '%d', current score: '%d'\n", outcomes[fs.Text()], score)
		score += outcomes[fs.Text()]
	}

  fmt.Println("Score: ", score)
}
