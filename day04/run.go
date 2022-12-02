package day04

import (
	"aoc2022/utils"
	"log"
	"strings"
)

type assignment struct {
	start int
	end   int
}

func Run() {
	content, err := utils.ReadFile("day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	part1(content)
}

func part1(content []string) {
	sum1 := 0
	sum2 := 0
	for _, s := range content {
		elves := strings.Split(s, ",")
		assignments := make([]assignment, len(elves))
		for i, elf := range elves {
			a := strings.Split(elf, "-")
			assignments[i] = assignment{
				start: utils.ToInt(a[0]),
				end:   utils.ToInt(a[1]),
			}
		}
		done := false
		if assignments[0].start <= assignments[1].start && assignments[1].end <= assignments[0].end {
			log.Printf("Assignments: %v", assignments)
			sum1 = sum1 + 1
			sum2 = sum2 + 1
			done = true
		}
		if !done {
			if assignments[1].start <= assignments[0].start && assignments[0].end <= assignments[1].end {
				log.Printf("Assignments: %v", assignments)
				sum1 = sum1 + 1
				sum2 = sum2 + 1
				done = true
			}
		}
		if !done {
			if assignments[0].start <= assignments[1].start && assignments[0].end >= assignments[1].start {
				log.Printf("Assignments: %v", assignments)
				sum2 = sum2 + 1
				done = true
			}
		}
		if !done {
			if assignments[1].start <= assignments[0].start && assignments[1].end >= assignments[0].start {
				log.Printf("Assignments: %v", assignments)
				sum2 = sum2 + 1
				done = true
			}
		}
	}
	log.Printf("Sum part 1: %d", sum1)
	log.Printf("Sum part 2: %d", sum2)
}
