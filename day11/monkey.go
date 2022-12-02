package day11

import (
	"aoc2022/utils"
	"regexp"
	"strings"
)

type Monkey struct {
	Items         []int
	Operation     func(int) int
	DivideBy      int
	SuccessTarget int
	ErrorTarget   int
}

func parseMonkeys(content []string) []*Monkey {
	monkeys := make([]*Monkey, 0)
	for i := 0; i < len(content); i = i + 7 {
		end := i + 7
		if end > len(content) {
			end = len(content)
		}
		monkeys = append(monkeys, parseMonkey(content[i:end]))
	}
	return monkeys
}

func parseMonkey(content []string) *Monkey {
	monkey := Monkey{}
	// starting items
	itemStrings := strings.Split(content[1][18:], ",")
	for _, itemString := range itemStrings {
		monkey.Items = append(monkey.Items, utils.ToInt(strings.TrimSpace(itemString)))
	}
	monkey.Operation = parseOperation()(content[2][19:])
	monkey.DivideBy = utils.ToInt(content[3][21:])
	monkey.SuccessTarget = utils.ToInt(content[4][29:])
	monkey.ErrorTarget = utils.ToInt(content[5][30:])
	return &monkey
}

func parseOperation() func(string) func(int) int {
	singlePattern := regexp.MustCompile(`old ([+*]) (\d+)`)
	nonePattern := regexp.MustCompile(`old ([+*]) old`)
	return func(operation string) func(int) int {
		if singlePattern.MatchString(operation) {
			matches := singlePattern.FindStringSubmatch(operation)
			return func(old int) int {
				if matches[1] == "+" {
					return old + utils.ToInt(matches[2])
				}
				return old * utils.ToInt(matches[2])
			}
		}
		if nonePattern.MatchString(operation) {
			matches := nonePattern.FindStringSubmatch(operation)
			return func(old int) int {
				if matches[1] == "+" {
					return old + old
				}
				return old * old
			}
		}
		return nil
	}
}
