package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		masstotal := mass/3 - 2
		fuelmass := masstotal
		for ok := true; ok; ok = (fuelmass >= 0) {
			fuelmass = fuelmass/3 - 2
			if fuelmass >= 0 {
				masstotal += fuelmass
			}
		}
		check(err)
		sum += masstotal
	}
	check(err)
	fmt.Println(sum)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
