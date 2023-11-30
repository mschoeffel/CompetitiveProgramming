package main

import (
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
		default:
			fmt.Println("Not implemented yet.")
			break
		}
	}
}
