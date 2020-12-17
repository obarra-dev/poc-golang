package main

import (
	"sort"
	"strings"
	"testing"
)

func TestHasSummedPairEqualValue(t *testing.T) {
	numbers := []int{1, 2, 3, 5, 8, 13}
	result, _, _ := hasSummedPairEqualValue(numbers, 8)

	if !result {
		t.Error("Error", result)
	}
}

func TestHasSummedPairEqualValueDay1(t *testing.T) {
	numbers := []int{1721, 366, 299, 675, 1456}
	result, fnumber, snumber := hasSummedPairEqualValue(numbers, 2020)

	if !result || fnumber != 299 || snumber != 1721 {
		t.Error("Error", result, fnumber, snumber)
	}
}

func hasSummedPairEqualValue(numbers []int, value int) (bool, int, int) {
	sort.Ints(numbers)

	var f, l int
	l = len(numbers) - 1
	for f < l {
		sum := numbers[f] + numbers[l]
		if value == sum {
			return true, numbers[f], numbers[l]
		} else if sum < value {
			f++
		} else {
			l--
		}
	}

	return false, -1, -1
}

func TestFindMissingNumber6(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 6}
	result := findMissingNumber(numbers, 6)

	if result != 5 {
		t.Error("Error", result)
	}
}

func TestFindMissingNumber8(t *testing.T) {
	numbers := []int{3, 7, 1, 2, 8, 4, 5}
	result := findMissingNumber(numbers, 8)

	if result != 6 {
		t.Error("Error", result)
	}
}

func findMissingNumber(numbers []int, leng int) int {
	sumTotal := leng * (leng + 1) / 2

	for _, v := range numbers {
		sumTotal = sumTotal - v
	}

	return sumTotal
}

func TestFindMissingNumberds8(t *testing.T) {
	result := reverseSenteces("omar barra")

	if result != "barra omar" {
		t.Error("Error", result)
	}
}

func reverseSenteces(phrase string) string {
	sentences := strings.Split(phrase, " ")
	var sb strings.Builder

	for i := len(sentences) - 1; i >= 0; i-- {
		sb.WriteString(sentences[i])
		if i != 0 {
			sb.WriteString(" ")
		}
	}

	return sb.String()
}
