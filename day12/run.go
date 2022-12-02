package day12

import (
	"aoc2022/utils"
	"fmt"
	"log"
)

func Run() {
	content, err := utils.ReadFile("day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	field, start, end := parseInput(content)
	// log.Println(field, start, end)

	var distances [][]int
	for i := 0; i < len(field); i++ {
		row := make([]int, len(field[i]))
		for j := 0; j < len(field[i]); j++ {
			row[j] = -1
		}
		distances = append(distances, row)
	}
	distances[start.I][start.J] = 0
	// log.Println(distances)

	rows, cols := len(field), len(field[0])
	edges := make(map[Position]bool)

	// edges[start] = true
	for i := range field {
		for j := range field[i] {
			if field[i][j] == 1 {
				distances[i][j] = 0
				edges[Position{i, j}] = true
			}
		}
	}

	for len(edges) > 0 {
		next := make(map[Position]bool)

		try := func(row, col, nextRow, nextCol int) {
			if nextRow < 0 || nextCol < 0 || nextRow >= rows || nextCol >= cols {
				return
			}
			if field[nextRow][nextCol] > field[row][col]+1 {
				return
			}
			if distances[nextRow][nextCol] != -1 && distances[nextRow][nextCol] <= distances[row][col]+1 {
				return
			}
			distances[nextRow][nextCol] = distances[row][col] + 1
			next[Position{nextRow, nextCol}] = true
		}

		for p := range edges {
			try(p.I, p.J, p.I-1, p.J)
			try(p.I, p.J, p.I+1, p.J)
			try(p.I, p.J, p.I, p.J-1)
			try(p.I, p.J, p.I, p.J+1)
		}
		edges = next
	}

	// log.Println(distances)
	log.Println(distances[end.I][end.J])
}

type Position struct {
	I int
	J int
}

func (p Position) String() string {
	return fmt.Sprintf("%d-%d", p.I, p.J)
}

func parseInput(content []string) ([][]int, Position, Position) {
	field := make([][]int, len(content))
	start := Position{}
	end := Position{}
	for i, s := range content {
		field[i] = make([]int, len(s))
		for j, c := range s {
			switch c {
			case 'S':
				field[i][j] = 1
				start.I = i
				start.J = j
			case 'E':
				field[i][j] = 26
				end.I = i
				end.J = j
			default:
				field[i][j] = int(c - 'a' + 1)
			}
		}
	}
	return field, start, end
}
