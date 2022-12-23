package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func compare(top sort.IntSlice, acc int) {
  top[3] = acc
  sort.Sort(sort.Reverse(top))
}

func main() {
  top := sort.IntSlice(make([]int, 4))
	acc := 0
	elf := 1

	f, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}

	fs := bufio.NewScanner(f)

	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		if fs.Text() == "" {
			fmt.Printf("Processing elf number '%d' with '%d' calories\n", elf, acc)

			compare(top, acc)

			elf = elf + 1
			acc = 0
		}

		c, err := strconv.Atoi(fs.Text())
		if err == nil {
			fmt.Printf("\tTotal caloires '%d', adding '%d' more\n", acc, c)
			acc = acc + c
		}
	}

	// check last entry, as no empty line is emmited
	compare(top, acc)

	f.Close()

	fmt.Println(top)
	fmt.Println("The total calories of the top 3 elfs are:", func(input []int) int {
		sum := 0
		for _, v := range input {
			fmt.Println(v)
			sum += v
		}

		return sum
  }(top[0:3]))
}
