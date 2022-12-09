package Day09

import (
	"AdventOfCode-2022/Utils"
	"fmt"
	"regexp"
	"strconv"
)

func Main() {
	commands := Utils.ReadFileByLinesString("\\Day09\\commands.txt")
	fmt.Println("Day 9: Result part one: " + strconv.Itoa(Day9Part1(commands)))
	fmt.Println("Day 9: Result part two: " + strconv.Itoa(Day9Part2(commands)))
}

func Day9Part1(commands []string) int {
	tailPositions := make(map[string]bool)
	headPosX := 0
	headPosY := 0
	headPrevPosX := 0
	headPrevPosY := 0
	tailPosX := 0
	tailPosY := 0

	for _, command := range commands {
		c := rune(command[0])
		reg := regexp.MustCompile(`\d+`)
		v, _ := strconv.Atoi(reg.FindAllString(command[2:], 1)[0])
		for i := 0; i < v; i++ {
			headPrevPosX = headPosX
			headPrevPosY = headPosY
			switch c {
			case 'R':
				headPosX++
				break
			case 'L':
				headPosX--
				break
			case 'U':
				headPosY--
				break
			case 'D':
				headPosY++
				break
			}
			if headPosX > tailPosX+1 || headPosX < tailPosX-1 || headPosY > tailPosY+1 || headPosY < tailPosY-1 {
				tailPosX = headPrevPosX
				tailPosY = headPrevPosY
			}
			tailPositions[strconv.Itoa(tailPosX)+"|"+strconv.Itoa(tailPosY)] = true
		}
	}

	return len(tailPositions)
}

func Day9Part2(commands []string) int {
	tailPositions := make(map[string]bool)
	nodePositions := [10][2]int{}

	for i := 0; i < len(nodePositions); i++ {
		nodePositions[i][0] = 0
		nodePositions[i][1] = 0
	}

	for _, command := range commands {
		c := rune(command[0])
		reg := regexp.MustCompile(`\d+`)
		v, _ := strconv.Atoi(reg.FindAllString(command[2:], 1)[0])
		for i := 0; i < v; i++ {
			switch c {
			case 'R':
				nodePositions[0][0]++
				break
			case 'L':
				nodePositions[0][0]--
				break
			case 'U':
				nodePositions[0][1]--
				break
			case 'D':
				nodePositions[0][1]++
				break
			}
			for nodeIdx := 0; nodeIdx < len(nodePositions)-1; nodeIdx++ {
				nodePositions[nodeIdx+1][0], nodePositions[nodeIdx+1][1] = moveNode(nodePositions[nodeIdx][0], nodePositions[nodeIdx][1], nodePositions[nodeIdx+1][0], nodePositions[nodeIdx+1][1])
			}
			tailPositions[strconv.Itoa(nodePositions[9][0])+"|"+strconv.Itoa(nodePositions[9][1])] = true
		}
	}

	return len(tailPositions)
}

func moveNode(headPosX int, headPosY int, nodePosX int, nodePosY int) (int, int) {
	if headPosX > nodePosX+1 || headPosX < nodePosX-1 || headPosY > nodePosY+1 || headPosY < nodePosY-1 {
		if headPosX > nodePosX {
			nodePosX += 1
		}
		if headPosY > nodePosY {
			nodePosY += 1
		}
		if headPosX < nodePosX {
			nodePosX -= 1
		}
		if headPosY < nodePosY {
			nodePosY -= 1
		}
	}
	return nodePosX, nodePosY

}
