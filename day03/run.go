package day03

import (
	"aoc2022/utils"
	"log"
	"strings"
)

func Run() {
	content, err := utils.ReadFile("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	part1(content)
	part2(content)
}

func part1(content []string) {
	sum := 0
	for _, s := range content {
		common := make(map[int32]bool, 0)
		if len(s)%2 != 0 {
			log.Fatalf("Line %s has an odd number of characters", s)
		}
		part1 := s[0 : len(s)/2]
		part2 := s[len(s)/2:]
		for _, c := range part1 {
			if strings.ContainsAny(part2, string(c)) {
				prio := c - 'a' + 1
				if prio < 0 {
					prio = c - 'A' + 27
				}
				if _, ok := common[prio]; !ok {
					common[prio] = true
				}
			}
		}
		for i := range common {
			sum = sum + int(i)
		}
	}
	log.Printf("Sum part 1: %d", sum)
}

func part2(content []string) {
	sum := 0
	if len(content)%3 != 0 {
		log.Fatalf("Input has %d lines, which is not a multiple of 3", len(content))
	}
	for i := 0; i < len(content); i = i + 3 {
		common := make(map[int32]bool, 0)
		part1 := content[i]
		part2 := content[i+1]
		part3 := content[i+2]
		for _, c := range part1 {
			if strings.ContainsAny(part2, string(c)) && strings.ContainsAny(part3, string(c)) {
				prio := c - 'a' + 1
				if prio < 0 {
					prio = c - 'A' + 27
				}
				if _, ok := common[prio]; !ok {
					common[prio] = true
				}
			}
		}
		for i := range common {
			sum = sum + int(i)
		}
	}
	log.Printf("Sum part 2: %d", sum)
}
