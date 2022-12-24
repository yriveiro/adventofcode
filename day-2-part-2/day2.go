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
// X -> Lose
// Y -> Draw
// Z -> Win

func main() {
  score := 0
	outcomes := map[string]int{
		"A X": 3, //  L 0 + 3
		"A Y": 4, //  D 3 + 1
		"A Z": 8, //  W 6 + 2
		
    "B X": 1, //  L 0 + 1
		"B Y": 5, //  D 3 + 2
		"B Z": 9, //  W 6 + 3

		"C X": 2, // L 0 + 2
		"C Y": 6, // D 3 + 3
		"C Z": 7, // W 6 + 1

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
