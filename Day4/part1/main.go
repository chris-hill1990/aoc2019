package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

}

func getSingleLineInput() []string {
	file, err := os.Open("../input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var arr []string

	for scanner.Scan() {
		input := scanner.Text()
		arr = strings.Split(input, ",")
	}
	return arr
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
