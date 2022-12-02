package day06

import (
	"aoc2022/utils"
	"log"
)

func Run() {
	content, err := utils.ReadFile("day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	targetLength := 14
	line := content[0]
	for i := 0; i < len(line); i++ {
		end := i + targetLength
		if end > len(line) {
			end = len(line)
		}
		if isValid(line[i:end], targetLength) {
			log.Println(i+targetLength, line[i:end])
			break
		}
	}
}

func isValid(content string, targetLength int) bool {
	letterMap := make(map[rune]int)
	for _, letter := range content {
		letterMap[letter]++
	}
	return len(letterMap) == targetLength
}
