package day11

import (
	"aoc2022/utils"
	"log"
	"sort"
)

func Run() {
	content, err := utils.ReadFile("day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	monkeys := parseMonkeys(content)
	inspect := make([]int, len(monkeys))

	divisor := 1
	for j, monkey := range monkeys {
		log.Println("Monkey", j, "has", monkey.Items)
		divisor = divisor * monkey.DivideBy
	}

	for i := 0; i < 10000; i++ {
		for j, monkey := range monkeys {
			handleMonkey(monkey, inspect, j, monkeys, divisor)
		}
		if i%100 == 0 {
			log.Println("Iteration", i)
		}
		//for j, monkey := range monkeys {
		//	log.Println("Monkey", j, "has", monkey.Items)
		//}
	}

	for i, i2 := range inspect {
		log.Printf("Monkey %d inspected %d items", i, i2)
	}
	sort.Ints(inspect)
	log.Printf("Inspect: %v", inspect[len(inspect)-2]*inspect[len(inspect)-1])
}

func handleMonkey(monkey *Monkey, inspect []int, j int, monkeys []*Monkey, divisor int) {
	for _, item := range monkey.Items {
		calc(inspect, j, monkey, item, monkeys, divisor)
	}
	monkey.Items = make([]int, 0)
}

func calc(inspect []int, j int, monkey *Monkey, item int, monkeys []*Monkey, divisor int) {
	inspect[j] = inspect[j] + 1
	value := monkey.Operation(item)
	// value = value / 3
	value = value % divisor
	if value%monkey.DivideBy == 0 {
		monkeys[monkey.SuccessTarget].Items = append(monkeys[monkey.SuccessTarget].Items, value)
	} else {
		monkeys[monkey.ErrorTarget].Items = append(monkeys[monkey.ErrorTarget].Items, value)
	}
}
