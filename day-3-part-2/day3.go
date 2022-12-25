package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Group struct {
	Current []string
	EOF     bool
	Size    int
}

func New(size int) *Group{
  g := &Group{
    Current: make([]string, size),
    EOF: false,
    Size: size,
  }

  return g
}

func (g *Group) Scan(fs *bufio.Scanner) bool {
	// hard coded the size of the group
  slice := make([]int, g.Size)

  // cleanup slice to reuse in next group
  g.Current = nil

	for range slice {
		g.EOF = fs.Scan()

		if !g.EOF {
			break
		}

		g.Current = append(g.Current, fs.Text())
	}

  return g.EOF
}

func (g *Group) GetBadget() rune {
	var bag rune

	fmt.Printf("Comparing %s\n", g.Current)

	for _, v := range g.Current[0] {
		if strings.ContainsRune(g.Current[1], v) && strings.ContainsRune(g.Current[2], v) {
      fmt.Printf("Found badget %s\n", string(v))
			bag = v
		}
	}

	return bag
}

func main() {
  g := New(3)
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

  for g.Scan(fs) {
    acc += priorities[g.GetBadget()]
  }

	fmt.Printf("Sum of priorities: %d", acc)
}
