package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var acc int

	f, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	fs := bufio.NewScanner(f)

	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		input := fs.Text()
		buf := map[string]int{}

		parsed_input := strings.Split(input, ": ")

		id := strings.Replace(parsed_input[0], "Game ", "", 1)
		game := strings.Split(parsed_input[1], "; ")

		fmt.Printf("\n\t┏ Assesing game id: %s\n", id)

		buf = map[string]int{
			"red":   1,
			"green": 1,
			"blue":  1,
		}

		for _, iteration := range game {
			fmt.Printf("\t┣ iteration: %#v\n", iteration)
			groups := regexp.MustCompile(`(?P<count>\d+) (?P<color>\w+)`).FindAllStringSubmatch(iteration, -1)

			for _, group := range groups {
				v, _ := strconv.Atoi(group[1])

				if v > buf[group[2]] {
					buf[group[2]] = v
				}

				fmt.Printf("\t┣ Current min valid iteration: %#v\n", buf)
			}
		}

		power := buf["red"] * buf["green"] * buf["blue"]
		fmt.Printf("\t┣ iteration power: %d\n", power)

		acc = acc + power
		fmt.Printf("\t┗ Current total power '%d' \n", acc)

	}
	fmt.Printf("The sum of all input values is '%d'", acc)
}
