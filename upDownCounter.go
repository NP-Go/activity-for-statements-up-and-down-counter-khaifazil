package main

import "fmt"

func main() {
	//Insert code here
	var num1 int
	var num2 int

	fmt.Println("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&num2)

	if num1 < num2 {
		for i := num1; i <= num2; i++ {
			fmt.Println(i)
		}

		for i := num2; i >= num1; i-- {
			fmt.Println(i)
		}
	} else if num1 > num2 {
		for i := num2; i <= num1; i++ {
			fmt.Println(i)
		}

		for i := num1; i >= num2; i-- {
			fmt.Println(i)
		}
	} else {
		fmt.Println("Error")
	}
}
