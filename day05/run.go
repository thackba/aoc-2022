package day05

import (
	"aoc2022/utils"
	"log"
	"regexp"
	"strings"
)

type Movement struct {
	From   int
	To     int
	Amount int
}

func Run() {
	stacks := 9
	part2 := true
	content, err := utils.ReadFile("day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	emptyLine := utils.FindEmptyLine(content)
	containers := setup(content[:emptyLine-1], stacks)
	movements := getMovements(content[emptyLine+1:])
	if !part2 {
		for _, move := range movements {
			for i := 0; i < move.Amount; i++ {
				containers[move.To] = append(containers[move.To], containers[move.From][len(containers[move.From])-1])
				containers[move.From] = containers[move.From][:len(containers[move.From])-1]
			}
		}
	} else {
		for _, move := range movements {
			containers[move.To] = append(containers[move.To], containers[move.From][len(containers[move.From])-move.Amount:]...)
			containers[move.From] = containers[move.From][:len(containers[move.From])-move.Amount]
		}
	}
	target := make([]string, stacks)
	for i, container := range containers {
		target[i] = container[len(container)-1]
	}
	log.Println(strings.Join(target, ""))
}

func setup(content []string, stacks int) [][]string {
	containers := make([][]string, stacks)
	for i := 0; i < stacks; i++ {
		containers[i] = make([]string, 0)
	}
	for _, line := range content {
		for i := 0; i < len(line); i = i + 4 {
			end := i + 4
			if end > len(line) {
				end = len(line)
			}
			container := line[i:end]
			if strings.Index(container, "[") != -1 {
				containers[i/4] = append([]string{container[1:2]}, containers[i/4]...)
			}
		}
	}
	return containers
}

func getMovements(content []string) []Movement {
	movements := make([]Movement, len(content))
	r := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")
	for i, move := range content {
		match := r.FindStringSubmatch(move)
		movements[i] = Movement{
			From:   utils.ToInt(match[2]) - 1,
			To:     utils.ToInt(match[3]) - 1,
			Amount: utils.ToInt(match[1]),
		}
	}
	return movements
}
