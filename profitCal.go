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
	Profit := EBT * (1 - taxRate/100)
	Ratio := EBT / Profit
	fmt.Printf("EBT :%.2f\n", EBT)
	fmt.Printf("The Profit is %.2f\n", Profit)
	fmt.Printf("The Ratio is %.2f\n", Ratio)
}
