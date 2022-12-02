package day09

import (
	"aoc2022/utils"
	"fmt"
	"log"
	"strconv"
)

type Move struct {
	Direction string
	Amount    int
}

type Position struct {
	X int
	Y int
}

func (p Position) String() string {
	return "(" + strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y) + ")"
}

func (p Position) TooFar(pos Position) bool {
	deltaX := p.X - pos.X
	deltaY := p.Y - pos.Y
	return (deltaX*deltaX) > 1 || (deltaY*deltaY) > 1
}

type Tail struct {
	Pos Position
}

func (t Tail) String() string {
	return t.Pos.String()
}

type Bridge struct {
	Head      *Position
	Tail      []Tail
	Positions map[Position]map[int]bool
}

func (b *Bridge) String() string {
	return fmt.Sprintf("Bridge: %v - %v", b.Head, b.Tail)
}

func (b *Bridge) Move(move Move) {
	for i := 0; i < move.Amount; i++ {
		switch move.Direction {
		case "U":
			b.Head.Y = b.Head.Y + 1
		case "D":
			b.Head.Y = b.Head.Y - 1
		case "R":
			b.Head.X = b.Head.X + 1
		case "L":
			b.Head.X = b.Head.X - 1
		}
		if b.Tail[0].Pos.TooFar(*b.Head) {
			b.Tail[0].Pos = b.getNearestPosition(*b.Head, b.Tail[0].Pos)
			if _, ok := b.Positions[b.Tail[0].Pos]; !ok {
				b.Positions[b.Tail[0].Pos] = make(map[int]bool)
			}
			b.Positions[b.Tail[0].Pos][0] = true
		}

		for j := 1; j < len(b.Tail); j++ {
			if b.Tail[j].Pos.TooFar(b.Tail[j-1].Pos) {
				b.Tail[j].Pos = b.getNearestPosition(b.Tail[j-1].Pos, b.Tail[j].Pos)
				if _, ok := b.Positions[b.Tail[j].Pos]; !ok {
					b.Positions[b.Tail[j].Pos] = make(map[int]bool)
				}
				b.Positions[b.Tail[j].Pos][j] = true
			}
		}
		// log.Println(b)
	}
}

func (b *Bridge) getNearestPosition(target Position, pos Position) Position {
	deltaX := target.X - pos.X
	deltaY := target.Y - pos.Y
	if deltaX > 0 {
		pos.X = pos.X + 1
	} else if deltaX < 0 {
		pos.X = pos.X - 1
	}
	if deltaY > 0 {
		pos.Y = pos.Y + 1
	} else if deltaY < 0 {
		pos.Y = pos.Y - 1
	}
	return pos
}

func Run() {
	content, err := utils.ReadFile("day09/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	moves := parseMoves(content)
	tailAmount := 9
	bridge := &Bridge{
		Head:      &Position{X: 0, Y: 0},
		Tail:      make([]Tail, tailAmount),
		Positions: make(map[Position]map[int]bool),
	}
	for i := 0; i < tailAmount; i++ {
		bridge.Tail[i] = Tail{
			Pos: Position{X: 0, Y: 0},
		}
	}
	for _, move := range moves {
		bridge.Move(move)
	}
	// log.Println(bridge.Positions)
	positionAmount := 0
	for _, positions := range bridge.Positions {
		if _, ok := positions[tailAmount-1]; ok {
			// log.Println(key)
			positionAmount++
		}
	}
	log.Printf("Tail Positions: %d", positionAmount+1) // plus startposition
}

func parseMoves(content []string) []Move {
	moves := make([]Move, len(content))
	for i, s := range content {
		moves[i] = Move{
			Direction: s[0:1],
			Amount:    utils.ToInt(s[2:]),
		}
	}
	return moves
}
