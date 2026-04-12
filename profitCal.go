package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	fmt.Print("Enter how much revenue you make: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter your expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter your present tax rate: ")
	fmt.Scan(&taxRate)
	EBT := revenue - expenses
	Profit := EBT * (1 - taxRate)
	Ratio := EBT / Profit
	fmt.Println(EBT)
	fmt.Println(Profit)
	fmt.Println(Ratio)
}
