package day10

import (
	"aoc2022/utils"
	"log"
	"strings"
)

type Operation struct {
	Delta    int
	Duration int
}

func Run() {
	content, err := utils.ReadFile("day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	operations := parseOperations(content)
	cycles := 0

	marker := []int{20, 60, 100, 140, 180, 220}
	counter := []int{1, 1, 1, 1, 1, 1}

	oldValue := 1
	var positions []int

	for _, o := range operations {
		for i := 0; i < o.Duration; i++ {
			positions = append(positions, oldValue)
		}
		cycles += o.Duration
		for i := 0; i < len(marker); i++ {
			if cycles < marker[i] {
				counter[i] += o.Delta
			}
		}
		oldValue += o.Delta
	}

	sum := 0
	for i, mark := range marker {
		sum += mark * counter[i]
	}
	log.Printf("Sum: %d", sum)

	crt := make([]string, 0)
	for i := 0; i < len(positions); i++ {
		if i%40 == 0 {
			crt = append(crt, "")
		}
		if i%40 >= positions[i]-1 && i%40 < positions[i]+2 {
			crt[i/40] += "#"
		} else {
			crt[i/40] += " "
		}
	}
	log.Print("\n" + strings.Join(crt, "\n"))
}

func parseOperations(input []string) []Operation {
	operations := make([]Operation, len(input))
	for i, line := range input {
		if line == "noop" {
			operations[i] = Operation{0, 1}
		} else {
			operations[i] = Operation{utils.ToInt(line[5:]), 2}
		}
	}
	return operations
}
