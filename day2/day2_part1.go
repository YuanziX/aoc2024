package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func input_is_safe(str_nums []string) bool {
	last_diff := 0
	safe := true

	for i := range str_nums[:len(str_nums)-1] {
		curr_num, _ := strconv.Atoi(str_nums[i])
		next_num, _ := strconv.Atoi(str_nums[i+1])

		new_diff := curr_num - next_num

		if last_diff < 0 && new_diff > 0 {
			safe = false
			break
		} else if last_diff > 0 && new_diff < 0 {
			safe = false
			break
		} else if new_diff > 3 || new_diff < -3 || new_diff == 0 {
			safe = false
			break
		}

		last_diff = new_diff
	}

	return safe
}

func part1() {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatalf("Something went wrong: %v\n", err)
	}

	input := string(file)

	count_safe := 0
	lines := strings.Split(input, "\n")

	// -1 to account for empty new line in input
	for _, line := range lines[:len(lines)-1] {
		str_nums := strings.Split(line, " ")
		if input_is_safe(str_nums) {
			count_safe++
		}
	}

	fmt.Printf("There are %d safe reports\n", count_safe)
}
