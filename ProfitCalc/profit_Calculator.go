package main

import "fmt"


func main() {
	revenue := userInput("Enter revenue amount: ")
	expense := userInput("Enter expense amount: ")
	taxRate := userInput("Enter tax rate: ")

	ebt := revenue - expense
	profit := ebt * (1 - taxRate/100)
	ratio := ebt/profit

	fmt.Printf("EBT: %.2f \nProfit: %.2f \nRatio: %.2f\n", ebt ,profit,ratio )


}

func userInput(message string) float64 {
	var userInput float64
	fmt.Print(message)
	fmt.Scan(&userInput)

	return userInput
}