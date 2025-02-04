package main

import "fmt"

func main(){
	var a, b float64
	var operation string
	
	fmt.Print("Enter first number:")
	fmt.Scan(&a)
	fmt.Print("Enter second number:")
	fmt.Scan(&b)
	fmt.Print("Enter operation (+, - , /, *):")
	fmt.Print(&operation)
}