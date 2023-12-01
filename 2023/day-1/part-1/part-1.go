package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

		if input != "" {
			input = regexp.MustCompile(`[a-z]+`).ReplaceAllString(input, "")
			input = fmt.Sprintf("%s%s", string(input[0]), string(input[len(input)-1]))

			if i, err := strconv.Atoi(input); err == nil {
				fmt.Printf("Processing input input '%d'\n", i)
				acc = acc + i
			} else {
				os.Exit(1)
			}
		}
	}
	fmt.Printf("The sum of all input values is '%d'", acc)
}
