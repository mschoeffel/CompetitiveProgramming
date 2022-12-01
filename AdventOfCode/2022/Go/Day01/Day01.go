package Day01

import (
	"AdventOfCode-2022/Utils"
	"fmt"
	"log"
	"strconv"
)

func Main() {
	inventory := Utils.ReadFileByLinesString("\\Day01\\inventory.txt")
	fmt.Println("Day 1: Result part one: " + strconv.Itoa(Day1Part1(inventory)))
	fmt.Println("Day 1: Result part two: " + strconv.Itoa(Day1Part2(inventory)))
}

func Day1Part1(inventory []string) int {
	maxCalories := 0
	currentCalories := 0
	for i := 0; i < len(inventory); i++ {
		if inventory[i] == "" {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
		} else {
			addCalories, err := strconv.Atoi(inventory[i])
			if err != nil {
				log.Fatal(err)
			}
			currentCalories += addCalories
		}
	}
	return maxCalories
}

func Day1Part2(inventory []string) int {
	maxThreeCalories := []int{0, 0, 0}
	currentCalories := 0
	for i := 0; i < len(inventory); i++ {
		if inventory[i] == "" {
			if currentCalories > maxThreeCalories[0] {
				maxThreeCalories[2] = maxThreeCalories[1]
				maxThreeCalories[1] = maxThreeCalories[0]
				maxThreeCalories[0] = currentCalories
			} else if currentCalories > maxThreeCalories[1] {
				maxThreeCalories[2] = maxThreeCalories[1]
				maxThreeCalories[1] = currentCalories
			} else if currentCalories > maxThreeCalories[2] {
				maxThreeCalories[2] = currentCalories
			}
			currentCalories = 0
		} else {
			addCalories, err := strconv.Atoi(inventory[i])
			if err != nil {
				log.Fatal(err)
			}
			currentCalories += addCalories
		}
	}
	sum := 0
	for _, calory := range maxThreeCalories {
		sum += calory
	}
	return sum
}
