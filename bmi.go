package main

import (
	"fmt"

	"github.com/redmicelles/goBMICalculator/info"
)

func main() {
	info.PrintWelcome()
	weight, height := getUserMetrics()

	bmi := weight / (height * height)

	fmt.Printf("%.2f\n", bmi)
}
