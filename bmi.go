package main

import (
	"github.com/redmicelles/goBMICalculator/info"
)

func main() {
	info.PrintWelcome()
	printBMI(calculateBMI())
}
