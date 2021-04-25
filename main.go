package main

import (
	"fmt"
	"strconv"
)

//Выяснить, входить ли цифра 3 в запись числа n^2

func Reader() string {
	fmt.Print("Enter integer: ")
	var input string
	fmt.Scan(&input)
	return input
}

func task88a(n string) bool {
	var counter = 0
	new, _ := strconv.Atoi(n)
	new *= new

	s := strconv.Itoa(new)
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if string([]rune(s)[i]) == "3" {
			counter += 1
			break
		}
	}
	if counter == 1 {
		fmt.Println("Yes, we have 3.")
		return true
	} else {
		fmt.Println("No, we haven't 3.")
		return false
	}
}

func main() {
	g := Reader()
	task88a(g)
}
