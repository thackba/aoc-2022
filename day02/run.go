package day02

import (
	"aoc2022/utils"
	"log"
	"strings"
)

func Run() {
	lines, err := utils.ReadFile("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	matrix := getMatrix()
	targetMatrix := getTargetMatrix()
	sumPart1 := 0
	sumPart2 := 0
	for _, line := range lines {
		choices := strings.Split(line, " ")
		if len(choices) != 2 {
			log.Fatal("invalid input")
		}
		e := getElfChoice(choices)
		m := getMyChoice(choices)
		// Part 1
		r := getResult(matrix, e, m)
		// log.Println(e, m, r)
		sumPart1 = sumPart1 + r + m
		// Part 2
		t := getResult(targetMatrix, e, m)
		// log.Println(e, m, t)
		sumPart2 = sumPart2 + t + ((m - 1) * 3)

	}
	log.Printf("Lösung Teil 1: %d", sumPart1)
	log.Printf("Lösung Teil 2: %d", sumPart2)
}

func getElfChoice(choices []string) int {
	return strings.Index("ABC", strings.TrimSpace(choices[0])) + 1
}

func getMyChoice(choices []string) int {
	return strings.Index("XYZ", strings.TrimSpace(choices[1])) + 1
}

func getResult(matrix [][]int, elfChoice, myChoice int) int {
	return matrix[elfChoice-1][myChoice-1]
}

func getMatrix() [][]int {
	return [][]int{
		{3, 6, 0},
		{0, 3, 6},
		{6, 0, 3},
	}
}

func getTargetMatrix() [][]int {
	return [][]int{
		{3, 1, 2},
		{1, 2, 3},
		{2, 3, 1},
	}
}
