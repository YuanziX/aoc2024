package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func part2() {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatalf("Something went wrong: %v\n", err)
	}

	input := string(file)

	count_safe := 0
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1] // -1 to account for empty new line in input

	for _, line := range lines {
		str_nums := strings.Split(line, " ")

		if input_is_safe(str_nums) {
			count_safe++
			continue
		}

		for i := 0; i < len(str_nums); i++ {
			new_list := make([]string, 0)
			new_list = append(new_list, str_nums[:i]...)
			new_list = append(new_list, str_nums[i+1:]...)

			if input_is_safe(new_list) {
				count_safe++
				break
			}
		}
	}

	fmt.Printf("There are %d safe reports\n", count_safe)
}
