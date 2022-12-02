package day08

import (
	"aoc2022/utils"
	"log"
)

func Run() {
	content, err := utils.ReadFile("day08/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(content)
	field := make([][]int32, len(content))
	for i, s := range content {
		field[i] = make([]int32, len(s))
		for j, c := range s {
			field[i][j] = c - '0'
		}
	}
	amountOfTrees := 0
	maxScore := 0
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if i == 0 || j == 0 || i == len(field)-1 || j == len(field[i])-1 {
				amountOfTrees = amountOfTrees + 1
			} else {
				isTreeVisible, score := isVisible(&field, i, j)
				if isTreeVisible {
					amountOfTrees = amountOfTrees + 1
				}
				if score > maxScore {
					maxScore = score
				}
			}
		}
	}
	log.Printf("Amount of trees: %d", amountOfTrees)
	log.Printf("Max score: %d", maxScore)
}

func isVisible(field *[][]int32, i int, j int) (bool, int) {
	treeHeight := (*field)[i][j]
	// log.Printf("Tree height: %d", treeHeight)
	// Check up
	scores := make([]int, 4)
	isVisibleUp := true
	for k := i - 1; k > -1; k-- {
		scores[0] = scores[0] + 1
		// log.Printf("Checking up: %d, %d -> %d", k, j, (*field)[k][j])
		if (*field)[k][j] >= treeHeight {
			isVisibleUp = false
			break
		}
	}
	// Check down
	isVisibleDown := true
	for k := i + 1; k < len(*field); k++ {
		scores[1] = scores[1] + 1
		// log.Printf("Checking down: %d, %d -> %d", k, j, (*field)[k][j])
		if (*field)[k][j] >= treeHeight {
			isVisibleDown = false
			break
		}
	}
	// Check left
	isVisibleLeft := true
	for k := j - 1; k > -1; k-- {
		scores[2] = scores[2] + 1
		// log.Printf("Checking left: %d, %d -> %d", k, j, (*field)[i][k])
		if (*field)[i][k] >= treeHeight {
			isVisibleLeft = false
			break
		}
	}
	// Check right
	isVisibleRight := true
	for k := j + 1; k < len((*field)[i]); k++ {
		scores[3] = scores[3] + 1
		// log.Printf("Checking right: %d, %d -> %d", k, j, (*field)[i][k])
		if (*field)[i][k] >= treeHeight {
			isVisibleRight = false
			break
		}
	}
	return isVisibleUp || isVisibleDown || isVisibleLeft || isVisibleRight, scores[0] * scores[1] * scores[2] * scores[3]
}
