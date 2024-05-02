package main

import (
	"fmt"
	"os"
	"strconv"
)

func factorial(n int) int{
	if n== 1  {
		return 1
	}
	return n*factorial(n-1)
}

func main(){
	fmt.Print("\t\t****FACTORIAL-CALCULATOR****\n\n\t\tNumber: ")
	var inputString string 
	fmt.Scanln(&inputString)

	number, err := strconv.Atoi(inputString)
	if err != nil {
		fmt.Println("An error occured!")
		os.Exit(1)
	}
	
	if number<1 {
		fmt.Println("Invalid number for Factorial!")
		os.Exit(1)
	}

	result := factorial(number)
	fmt.Println("\n\t\tThe Factorial of ",inputString," is ",result)

}
