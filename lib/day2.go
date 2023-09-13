package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getGameScore(them string, us string) int {
	switch us {
	case "X":
		us = "A"
	case "Y":
		us = "B"
	case "Z":
		us = "C"
	}

	if us == them {
		return 3
	}
	if (us == "A" && them == "C") ||
		(us == "B" && them == "A") ||
		(us == "C" && them == "B") {
		return 6
	}

	return 0
}

func getShapeScore(shape string) int {
	switch shape {
	case "X", "A":
		return 1
	case "Y", "B":
		return 2
	case "Z", "C":
		return 3
	}
	return 0
}

func sumTo3(num1 int, num2 int) int {
	sum := num1 + num2
	if sum > 3 {
		sum -= 3
	}
	return sum
}

func getMatchScore(shape string, result string) int {
	shapeScore := getShapeScore(shape)
	switch result {
	case "X":
		return 0 + sumTo3(shapeScore, 2)
	case "Y":
		return 3 + shapeScore
	case "Z":
		return 6 + sumTo3(shapeScore, 1)
	default:
		return 0
	}
}

func Day2Solution() int {
	file, err := os.Open("./input/day2.txt")
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	scanner := bufio.NewScanner(file)
	totalScore := 0
	for scanner.Scan() {
		text := scanner.Text()
		scores := strings.Split(text, " ")
		// First part
		// totalScore += getGameScore(scores[0], scores[1])
		// totalScore += getShapeScore(scores[1])

		//Second part
		totalScore += getMatchScore(scores[0], scores[1])
	}
	return totalScore
}
