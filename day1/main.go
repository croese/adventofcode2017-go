package main

import (
	"fmt"
	"os"
)

type circularIntSlice []int

func (i circularIntSlice) elem(index int) int {
	return i[index%len(i)]
}

func main() {
	part := os.Args[1]
	input := os.Args[2]

	var captchaFunc func(circularIntSlice) int
	if part == "a" {
		captchaFunc = calculateFirstCaptcha
	} else {
		captchaFunc = calculateSecondCaptcha
	}

	fmt.Println(captchaFunc(stringToDigits(input)))
}

func stringToDigits(s string) circularIntSlice {
	result := []int{}

	for _, c := range s {
		result = append(result, int(c-'0'))
	}

	return result
}

func calculateFirstCaptcha(digits circularIntSlice) int {
	if digits == nil || len(digits) == 0 {
		return 0
	}

	sum := 0

	for i := 0; i < len(digits); i++ {
		if digits.elem(i) == digits.elem(i+1) {
			sum += digits.elem(i)
		}
	}

	return sum
}

func calculateSecondCaptcha(digits circularIntSlice) int {
	if digits == nil || len(digits) == 0 {
		return 0
	}

	sum := 0
	len := len(digits)
	step := len / 2 // puzzle input length always even
	for i := 0; i < len; i++ {
		if digits.elem(i) == digits.elem(i+step) {
			sum += digits.elem(i)
		}
	}

	return sum
}
