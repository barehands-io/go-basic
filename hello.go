package main

import (
	"fmt"
	"go-basic/utils"
)

func main() {
	var (
		a int
		b int    = 1
		c string = "hello"
	)

	var listOfStudents = []string{"John", "Jane", "Mary"}

	var accountBalance = 100.50

	var price = 300

	if accountBalance > float64(price) {
		fmt.Println("You can buy this item")
	} else {
		fmt.Println("You don't have enough money")
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(listOfStudents)

	utils.CheckUsers()

}
