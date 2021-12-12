package main

func calculateBMI() float64 {
	weight, height := getUserMetrics()
	return weight / (height * height)
}
