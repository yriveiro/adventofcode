package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var top int
	var acc int
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

			if acc > top {
				top = acc
			}

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
	if acc > top {
		top = acc
	}

	f.Close()
	fmt.Printf("The elf with most calories has '%d'", top)
}
