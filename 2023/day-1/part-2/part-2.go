package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

func main() {
	var acc int

	// in this case the numbers can overlap, to break that overlaping we do a
	// two phase pass one removing the overlaping the second to
	overlap := map[string]string{
		"one":   "1e",
		"two":   "2o",
		"three": "3e",
		"four":  "4",
		"five":  "5e",
		"six":   "6",
		"seven": "7n",
		"eight": "8t",
		"nine":  "9e",
	}

	numbers := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
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

		if input != "" {
			fmt.Printf("%s -> ", input)
			pattern := strings.Join(maps.Keys(overlap), "|")

			// fix overlaps
			input = regexp.MustCompile(pattern).ReplaceAllStringFunc(input, func(m string) string { return overlap[m] })

			// final pass
			input = regexp.MustCompile(pattern).ReplaceAllStringFunc(input, func(m string) string { return numbers[m] })
			fmt.Printf("%s-> ", input)

			input = regexp.MustCompile(`[a-z]+`).ReplaceAllString(input, "")
			input = fmt.Sprintf("%s%s", string(input[0]), string(input[len(input)-1]))

			if i, err := strconv.Atoi(input); err == nil {
				fmt.Printf("%d\n", i)
				acc = acc + i
			} else {
				os.Exit(1)
			}
		}
	}
	fmt.Printf("The sum of all input values is '%d'", acc)
}
