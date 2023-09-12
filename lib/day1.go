package lib

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func sumArrayValues(arr []int) int {
	sum := 0
	for _, valueInt := range arr {
		sum += valueInt
	}
	return sum
}

func findElPlace(el int, arr []int) []int {
	finalArr := make([]int, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		finalArr[i] = math.MinInt
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] <= el {
			finalArr[i] = arr[i]
			if i+1 == len(arr) {
				finalArr[i+1] = el
			} else if arr[i+1] > el {
				finalArr[i+1] = el
			}
		} else {
			finalArr[i+1] = arr[i]
		}
	}

	return finalArr[1:]
}

func Day1Solution() int {
	file, err := os.Open("./input/day1.txt")
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	scanner := bufio.NewScanner(file)

	maxSum := make([]int, 3)
	maxSum[0] = 0
	maxSum[1] = 0
	maxSum[2] = 0
	localSum := 0
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			maxSum = findElPlace(localSum, maxSum)
			localSum = 0
			continue
		}

		currNum, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(err.Error())
			return 0
		}

		localSum += currNum

	}

	return sumArrayValues(maxSum)
}

func iDay1Solution() int {
	file, err := os.Open("./input/day1.txt")
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	fmt.Println(file.Name())

	return 0
}
