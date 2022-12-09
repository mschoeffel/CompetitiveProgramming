package main

import (
	"AdventOfCode-2022/Day01"
	"AdventOfCode-2022/Day02"
	"AdventOfCode-2022/Day03"
	"AdventOfCode-2022/Day04"
	"AdventOfCode-2022/Day05"
	"AdventOfCode-2022/Day06"
	"AdventOfCode-2022/Day07"
	"AdventOfCode-2022/Day08"
	"AdventOfCode-2022/Day09"
	"fmt"
	"log"
)

func main() {
	fmt.Print("Enter number of day: ")
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Invalid input.")
		log.Fatal(err)
	} else {
		switch number {
		case 1:
			Day01.Main()
			break
		case 2:
			Day02.Main()
			break
		case 3:
			Day03.Main()
			break
		case 4:
			Day04.Main()
			break
		case 5:
			Day05.Main()
			break
		case 6:
			Day06.Main()
			break
		case 7:
			Day07.Main()
			break
		case 8:
			Day08.Main()
			break
		case 9:
			Day09.Main()
			break
		}
	}
}
