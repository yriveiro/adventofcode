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
	var valid bool

	max := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	f, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	fs := bufio.NewScanner(f)

	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		input := fs.Text()

		parsed_input := strings.Split(input, ": ")

		id := strings.Replace(parsed_input[0], "Game ", "", 1)
		games := strings.Split(parsed_input[1], "; ")

		valid = true

		fmt.Printf("\n\t┏ Assesing game id: %s\n", id)

		for _, game := range games {
			fmt.Printf("\t┣ Game: %#v\n", game)
			groups := regexp.MustCompile(`(?P<count>\d+) (?P<color>\w+)`).FindAllStringSubmatch(game, -1)

			for _, group := range groups {
				v, _ := strconv.Atoi(group[1])

				if v > max[group[2]] {
					fmt.Printf("\t┣ color %s | value: %d, max: %d | not valid game %#v\n", group[2], v, max[group[2]], group)
					valid = false
					break
				}

				fmt.Printf("\t┣ %#v | valid game\n", v)
			}
		}

		fmt.Printf("\t┗ Is Valid? %v\n", valid)

		if valid {
			if i, err := strconv.Atoi(id); err == nil {
				fmt.Printf("\nGame id '%d' added | Total %d\n", i, acc)
				acc = acc + i
			} else {
				os.Exit(1)
			}
		}
	}
	fmt.Printf("The sum of all input values is '%d'", acc)
}
