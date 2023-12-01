package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var mapping = map[string]string{
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

func findFirstWord(str string) (int, string) {
	lowestIdx := -1
	val := ""

	for k := range mapping {
		idx := strings.Index(str, k)

		if (idx > -1 && lowestIdx == -1) || (idx > -1 && idx < lowestIdx) {
			lowestIdx = idx
			val = string(k)
		}
	}

	return lowestIdx, val
}

func findLastWord(str string) (int, string) {
	biggestIdx := -1
	val := ""
	for k := range mapping {
		idx := strings.LastIndex(str, k)

		if idx > biggestIdx {
			biggestIdx = idx
			val = string(k)
		}
	}

	return biggestIdx, val
}

func findFirstDigit(str string) (int, string) {
	idx := strings.IndexFunc(str, func(r rune) bool {
		return r >= '0' && r <= '9'
	})

	if idx == -1 {
		return idx, ""
	}

	return idx, string(str[idx])
}

func findFirstVal(str string) string {
	firstDigitIdx, firstDigit := findFirstDigit(str)
	firstWordIdx, firstWord := findFirstWord(str)

	if (firstWordIdx == -1) || firstDigitIdx > -1 && firstDigitIdx < firstWordIdx {
		return firstDigit
	}

	return mapping[firstWord]
}

func findLastDigit(str string) (int, string) {
	idx := strings.LastIndexFunc(str, func(r rune) bool {
		return r >= '0' && r <= '9'
	})

	if idx == -1 {
		return idx, ""
	}

	return idx, string(str[idx])
}

func findLastVal(str string) string {
	lastDigitIdx, lastDigit := findLastDigit(str)

	lastWordIdx, lastWord := findLastWord(str)

	if lastDigitIdx > lastWordIdx {
		return lastDigit
	}

	return mapping[lastWord]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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

		fmt.Println("str", str)

		x, w := findFirstWord(str)
		fmt.Println("first word", x, w)

		x, w = findLastWord(str)
		fmt.Println("last word", x, w)

		x, w = findFirstDigit(str)
		fmt.Println("first digit", x, w)

		x, w = findLastDigit(str)
		fmt.Println("last digit", x, w)

		fmt.Println("first", findFirstVal(str))
		fmt.Println("last", findLastVal(str))

		num, err := strconv.Atoi(findFirstVal(str) + findLastVal(str))
		check(err)

		fmt.Println("num", num)

		if err != nil {
			panic(err)
		}
		sum += num
	}

	fmt.Println(sum)

}
