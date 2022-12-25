package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func split(input string) (string, string) {
	l := len(input)

	return input[0 : l/2], input[(l / 2):l]
}

func common(l, r string) rune {
	var bag rune

	fmt.Printf("Comparing %s <-> %s\n", l, r)

	for _, v := range l {
		if strings.ContainsRune(r, v) {
      fmt.Printf("Found misplaced item %s\n", string(v))

      // if more than one is found are always duplicated (checked before hand)
			bag = v
		}
	}

	return bag
}

func main() {
	acc := 0
	priorities := map[rune]int{
		'a': 1,
		'b': 2,
		'c': 3,
		'd': 4,
		'e': 5,
		'f': 6,
		'g': 7,
		'h': 8,
		'i': 9,
		'j': 10,
		'k': 11,
		'l': 12,
		'm': 13,
		'n': 14,
		'o': 15,
		'p': 16,
		'q': 17,
		'r': 18,
		's': 19,
		't': 20,
		'u': 21,
		'v': 22,
		'w': 23,
		'x': 24,
		'y': 25,
		'z': 26,
		'A': 27,
		'B': 28,
		'C': 29,
		'D': 30,
		'E': 31,
		'F': 32,
		'G': 33,
		'H': 34,
		'I': 35,
		'J': 36,
		'K': 37,
		'L': 38,
		'M': 39,
		'N': 40,
		'O': 41,
		'P': 42,
		'Q': 43,
		'R': 44,
		'S': 45,
		'T': 46,
		'U': 47,
		'V': 48,
		'W': 49,
		'X': 50,
		'Y': 51,
		'Z': 52,

	}

	f, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	fs := bufio.NewScanner(f)

	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		left, right := split(fs.Text())
		c := common(left, right)
		acc += priorities[c]
	}

	fmt.Printf("Sum of priorities: %d", acc)
}
