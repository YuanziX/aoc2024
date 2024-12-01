package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partition(arr []int, l, r int) int {
	piv := arr[r]
	i := l - 1

	for j := l; j <= r; j++ {
		if arr[j] < piv {
			i++
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	temp := arr[i+1]
	arr[i+1] = arr[r]
	arr[r] = temp

	return i + 1
}

func qs(arr []int, l, r int) {
	if l < r {
		p := partition(arr, l, r)
		qs(arr, l, p-1)
		qs(arr, p+1, r)
	}
}

func abs_diff(num1, num2 int) int {
	if num1 > num2 {
		return num1 - num2
	}
	return num2 - num1
}

func part1() {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatalf("Something went wrong: %v\n", err)
	}

	input := string(file)

	ls := strings.Split(input, "\n")

	c1 := make([]int, len(ls))
	c2 := make([]int, len(ls))

	for i := 0; i < len(ls); i++ {
		dat := strings.Split(ls[i], "   ")
		if len(dat) < 2 {
			continue
		}

		num1, err := strconv.Atoi(dat[0])
		if err != nil {
			log.Fatalf("Something went wrong: %v\n", err)
		}
		c1[i] = num1

		num2, err := strconv.Atoi(dat[1])
		if err != nil {
			log.Fatalf("Something went wrong: %v\n", err)
		}
		c2[i] = num2
	}

	qs(c1, 0, len(c1)-1)
	qs(c2, 0, len(c2)-1)

	res := 0

	for i := 0; i < len(c1); i++ {
		res += abs_diff(c1[i], c2[i])
	}

	fmt.Printf("Distance is %v\n", res)
}
