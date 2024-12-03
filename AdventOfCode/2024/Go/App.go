package main

import (
	"AdventOfCode-2024/Day01"
	"AdventOfCode-2024/Day02"
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
		default:
			fmt.Println("Not implemented yet.")
			break
		}
	}
}
