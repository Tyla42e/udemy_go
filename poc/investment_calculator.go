package main

import (
	"fmt"
	"math"
)

const INFLATION_RATE = 2.5

func main() {
	investmentAmount := userInput("Enter Investment Amount: ")
	expectedReturnRate := userInput("Enter Expected Return Rate: ")
	years := int(userInput("Enter number of years: "))

	futureValue, futureRealValue := calculateFinancials(investmentAmount, expectedReturnRate, years)

	fmt.Printf("Future value %.2f\nFuture Real Value: %.2f\n", futureValue, futureRealValue)
}

func userInput(message string) float64 {
	var userInput float64
	fmt.Print(message)
	fmt.Scan(&userInput)

	return userInput
}

func calculateFinancials(investmentAmount, expectedReturnRate float64, years int) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
	frv = fv / math.Pow(1+INFLATION_RATE/100, float64(years))

	return
}
