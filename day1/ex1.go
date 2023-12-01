package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		str := scanner.Text()

		firstDigit := strings.IndexFunc(str, func(r rune) bool {
			return r >= '0' && r <= '9'
		})

		lastDigit := strings.LastIndexFunc(str, func(r rune) bool {
			return r >= '0' && r <= '9'
		})

		fmt.Println(str)
		fmt.Println(firstDigit, lastDigit)

		num, err := strconv.Atoi(string(str[firstDigit]) + string(str[lastDigit]))

		fmt.Println(num)

		if err != nil {
			panic(err)
		}
		sum += num
	}

	fmt.Println(sum)

}
