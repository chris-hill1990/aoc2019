package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var memarr []string
	x := 0
	y := 0

	for ok := true; ok; ok = (x != 99 || y != 99) {
		index := 0
		code := 0
		file, err := os.Open("../input.txt")
		check(err)
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			input := scanner.Text()
			memarr = strings.Split(input, ",")

		}
		file.Close()
		var arr = memarr
		copy(memarr, arr)

		fmt.Println("trying x: ", x)
		fmt.Println("trying y: ", y)

		arr[1] = strconv.Itoa(x)
		arr[2] = strconv.Itoa(y)

		for ok := true; ok; ok = (code != 99) {
			code, err = strconv.Atoi(arr[index])
			check(err)

			a, err := strconv.Atoi(arr[index+1])
			check(err)
			b, err := strconv.Atoi(arr[index+2])
			check(err)
			in1, err := strconv.Atoi(arr[a])
			check(err)
			in2, err := strconv.Atoi(arr[b])
			check(err)

			switch code {
			case 1:
				c := in1 + in2
				loc, err := strconv.Atoi(arr[index+3])
				check(err)
				arr[loc] = strconv.Itoa(c)
			case 2:
				c := in1 * in2
				loc, err := strconv.Atoi(arr[index+3])
				check(err)
				arr[loc] = strconv.Itoa(c)
			default:
				break
			}

			index += 4

		}

		if arr[0] != strconv.Itoa(19690720) {
			if y == 99 {
				x++
				y = 0
			} else {
				y++
			}
		} else {
			fmt.Println("answer: ", 100*x+y)
			os.Exit(0)
		}
		arr = nil
	}

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
