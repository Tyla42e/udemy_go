package main

import (
	"fmt"
	"math"
)

func main() {

	// var investmentAmount,years float64 = 1000,10
	// investmentAmount,years := 1000.00, 10.0
	// var a,b = 1000.00,"10"
	const INFLATION_RATE = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years int

	fmt.Print("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/ 100,float64(years))
	futureRealValue := futureValue / math.Pow(1 + INFLATION_RATE /100, float64(years))
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}