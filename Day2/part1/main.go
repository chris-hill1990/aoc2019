package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var arr []string

	for scanner.Scan() {
		input := scanner.Text()
		arr = strings.Split(input, ",")
		fmt.Println(arr)

	}

	length := len(arr)
	fmt.Println(length)
	index := 0
	code := 0
	arr[1] = "12"
	arr[2] = "2"

	for ok := true; ok; ok = (code != 99) {

		code, err = strconv.Atoi(arr[index])
		check(err)
		fmt.Println("Operation: ", code)
		fmt.Println("Index: ", index)

		switch code {
		case 1:
			a, err := strconv.Atoi(arr[index+1])
			check(err)
			b, err := strconv.Atoi(arr[index+2])
			check(err)
			in1, err := strconv.Atoi(arr[a])
			check(err)
			in2, err := strconv.Atoi(arr[b])
			check(err)
			c := in1 + in2
			loc, err := strconv.Atoi(arr[index+3])
			check(err)
			arr[loc] = strconv.Itoa(c)
		case 2:
			a, err := strconv.Atoi(arr[index+1])
			check(err)
			b, err := strconv.Atoi(arr[index+2])
			check(err)
			in1, err := strconv.Atoi(arr[a])
			check(err)
			in2, err := strconv.Atoi(arr[b])
			check(err)
			c := in1 * in2
			loc, err := strconv.Atoi(arr[index+3])
			check(err)
			arr[loc] = strconv.Itoa(c)
		}
		index += 4
		fmt.Println(arr)

	}

	fmt.Println(arr[0])

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
