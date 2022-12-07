package Day07

import (
	"AdventOfCode-2022/Utils"
	"fmt"
	"regexp"
	"strconv"
)

type Node struct {
	Name     string
	Size     int
	Parent   *Node
	Children []*Node
}

func Main() {
	sections := Utils.ReadFileByLinesString("\\Day07\\commands.txt")
	fmt.Println("Day 7: Result part one: " + strconv.Itoa(Day7Part1(sections)))
	fmt.Println("Day 7: Result part two: " + strconv.Itoa(Day7Part2(sections)))
}

func Day7Part1(commands []string) int {
	root := createFileTree(commands)
	calcSizeOfNodes(root)

	return sumNodesLimit(root)
}

func Day7Part2(commands []string) int {
	root := createFileTree(commands)
	calcSizeOfNodes(root)
	freeSpace := 70000000 - root.Size
	spaceNeeded := 30000000 - freeSpace
	nodes := findDirsBiggerThan(spaceNeeded, root)

	smallest := &Node{Size: 70000000}
	for _, node := range nodes {
		if node.Size < smallest.Size {
			smallest = node
		}
	}
	return smallest.Size
}

func createFileTree(commands []string) *Node {
	root := &Node{Name: "/", Size: 0, Parent: nil, Children: []*Node{}}
	currentNode := root

	for _, command := range commands {
		if command[0] == '$' {
			//command
			switch command[2:4] {
			case "cd":
				//Move
				if len(command) == 6 && command[5] == '/' {
					//Move to root
					currentNode = root
				} else if len(command) == 7 && command[5:7] == ".." {
					//Move up
					currentNode = currentNode.Parent
				} else {
					//Move to directory or create directory
					directoryName := command[5:]
					exists := false
					for _, child := range currentNode.Children {
						if child.Name == directoryName {
							exists = true
							currentNode = child
							break
						}
					}
					if !exists {
						dirNode := &Node{Name: directoryName, Size: 0, Parent: currentNode, Children: []*Node{}}
						currentNode.Children = append(currentNode.Children, dirNode)
						currentNode = dirNode
					}
				}
				break
			case "ls":
				//List -> do nothing
				break
			default:
				//do nothing
			}
		} else {
			//file
			if len(command) >= 3 && command[0:3] == "dir" {

			} else {
				reg := regexp.MustCompile(`\d+`)
				size := reg.FindAllString(command, 1)[0]
				name := command[len(size)+1:]
				sizeN, _ := strconv.Atoi(size)
				currentNode.Children = append(currentNode.Children, &Node{Name: name, Size: sizeN, Parent: currentNode, Children: nil})
			}
		}
	}
	return root
}

func calcSizeOfNodes(root *Node) int {
	size := 0
	if root.Children != nil {
		for _, child := range root.Children {
			if child != nil {
				size += calcSizeOfNodes(child)
			}
		}
	}
	root.Size = root.Size + size
	return root.Size
}

func sumNodesLimit(root *Node) int {
	size := 0
	if root.Children != nil && len(root.Children) > 0 {
		for _, child := range root.Children {
			size += sumNodesLimit(child)
		}
		if root.Size <= 100000 {
			size += root.Size
		}
	}
	return size
}

func findDirsBiggerThan(size int, root *Node) []*Node {
	nodes := []*Node{}
	if root.Children != nil && len(root.Children) > 0 {
		for _, child := range root.Children {
			nodes = append(nodes, findDirsBiggerThan(size, child)...)
		}
		if root.Size >= size {
			nodes = append(nodes, root)
		}
	}
	return nodes
}
