package day01

import (
	"aoc2022/utils"
	"log"
	"sort"
	"strconv"
	"strings"
)

func Run() {
	calories := make([]int, 0)
	counter := 0
	lines, err := utils.ReadFile("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if len(trimmed) == 0 && counter > 0 {
			calories = append(calories, counter)
			counter = 0
		} else {
			value, err := strconv.Atoi(trimmed)
			if err != nil {
				log.Fatal(err)
			}
			counter = counter + value
		}
	}
	if counter > 0 {
		calories = append(calories, counter)
	}
	sort.Ints(calories)
	// Part 1
	log.Printf("Max calories: %d", calories[len(calories)-1])
	// Part 2
	sort.Ints(calories)
	sum := 0
	for _, calorie := range calories[len(calories)-3:] {
		sum = sum + calorie
	}
	log.Printf("Sum most three calorie values: %d", sum)
}
