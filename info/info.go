package info

import "fmt"

const mainTitle = "BMI Calculator"
const separator = "------------------------------------------"
const WeightPrompt = "Please enter your Weight(kg): "
const HeightPrompt = "Please enter your height(m): "

func PrintWelcome() {
	fmt.Println(mainTitle)
	fmt.Println(mainTitle)
	fmt.Println(separator)
}
