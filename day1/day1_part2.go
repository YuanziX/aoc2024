package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part2() {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatalf("Something went wrong: %v\n", err)
	}

	input := string(file)

	ls := strings.Split(input, "\n")

	nums := make(map[int]int, 0)

	for i := 0; i < len(ls); i++ {
		dat := strings.Split(ls[i], "   ")
		if len(dat) < 2 {
			continue
		}

		num1, err := strconv.Atoi(dat[0])
		if err != nil {
			log.Fatalf("Something went wrong: %v\n", err)
		}

		_, ok := nums[num1]
		if !ok {
			nums[num1] = 0
		}
	}

	for i := 0; i < len(ls); i++ {
		dat := strings.Split(ls[i], "   ")
		if len(dat) < 2 {
			continue
		}

		num2, err := strconv.Atoi(dat[1])
		if err != nil {
			log.Fatalf("Something went wrong: %v\n", err)
		}

		val, ok := nums[num2]
		if !ok {
			continue
		}

		nums[num2] = val + 1
	}

	res := 0

	for i := 0; i < len(ls); i++ {
		dat := strings.Split(ls[i], "   ")
		if len(dat) < 2 {
			continue
		}

		num1, err := strconv.Atoi(dat[0])
		if err != nil {
			log.Fatalf("Something went wrong: %v\n", err)
		}

		val := nums[num1]
		res += num1 * val
	}
	fmt.Printf("Similarity score is %v\n", res)
}
