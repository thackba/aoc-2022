package day07

import (
	"aoc2022/utils"
	"fmt"
	"log"
	"regexp"
)

type TreeNode struct {
	Name     string
	Size     int
	Total    int
	Parent   *TreeNode
	Children []*TreeNode
}

func (t *TreeNode) String() string {
	return fmt.Sprintf("%s %d %d, %v", t.Name, t.Size, t.Total, t.Children)
}

func Run() {
	content, err := utils.ReadFile("day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rootNode := &TreeNode{"/", 0, 0, nil, make([]*TreeNode, 0)}
	parseInput(content[1:], rootNode)
	updateTotal(rootNode)
	log.Printf("At Most: %d", findAtMost(rootNode, 100000))
	maxSpace := 70000000
	minSpaceFree := 30000000
	log.Println(rootNode)
	minDelete := minSpaceFree - (maxSpace - rootNode.Total)
	log.Printf("Free Space: %d / %d", maxSpace-rootNode.Total, minSpaceFree)
	log.Printf("Delete: %d", minDelete)
	possibleDirectories := findDirectoryToDelete(rootNode, minDelete)
	var dirToDelete *TreeNode
	minValue := maxSpace
	for _, dir := range possibleDirectories {
		if dir.Total < minValue {
			minValue = dir.Total
			dirToDelete = dir
		}
	}
	log.Printf("Delete: %s %d", dirToDelete.Name, dirToDelete.Total)
}

func parseInput(input []string, rootNode *TreeNode) {
	currentNode := rootNode
	fileRegex, err := regexp.Compile(`(\d+) ([\w.]+)`)
	if err != nil {
		log.Fatal(err)
	}
	cdRegex, err := regexp.Compile(`\$ cd ([\w.]+)`)
	for _, s := range input {
		if fileRegex.MatchString(s) {
			matches := fileRegex.FindStringSubmatch(s)
			currentNode.Size += utils.ToInt(matches[1])
		}
		if cdRegex.MatchString(s) {
			matches := cdRegex.FindStringSubmatch(s)
			if matches[1] == ".." {
				currentNode = currentNode.Parent
			} else {
				newNode := &TreeNode{matches[1], 0, 0, currentNode, make([]*TreeNode, 0)}
				currentNode.Children = append(currentNode.Children, newNode)
				currentNode = newNode
			}
		}
	}
}

func updateTotal(node *TreeNode) {
	for _, child := range node.Children {
		updateTotal(child)
		node.Total += child.Total
	}
	node.Total += node.Size
}

func findAtMost(node *TreeNode, atMost int) int {
	result := 0
	if node.Total <= atMost {
		log.Println(node.Name, node.Total)
		result += node.Total
	}
	for _, child := range node.Children {
		result += findAtMost(child, atMost)
	}
	return result
}

func findDirectoryToDelete(node *TreeNode, minDelete int) []*TreeNode {
	result := make([]*TreeNode, 0)
	if node.Total >= minDelete {
		result = append(result, node)
	}
	for _, child := range node.Children {
		result = append(result, findDirectoryToDelete(child, minDelete)...)
	}
	return result
}
