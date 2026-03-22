package main

import "fmt"

/*func createTemperatureAdjuster() (func(change float64) float64, float64) {
	baseTemperature := 90.0

	adjustTemperature := func(change float64) float64 {
		baseTemperature = baseTemperature + change
		return baseTemperature
	}

	return adjustTemperature, baseTemperature
}

func main() {
	adjustTemp, originalTemp := createTemperatureAdjuster()

	fmt.Printf("Original temperature is %.1f\n", originalTemp)

	fmt.Printf("Adjusted Temp +1.5: %.1f grad C\n", adjustTemp(1.5))
	fmt.Printf("Adjusted Temp -3.0: %.1f grad C\n", adjustTemp(-3.0))
} */

/*
func counter(start int) func() int{
	count := start
	currentCount := func() int {
		count = count +1

		return count
	}

	return currentCount
}


func main(){
c := counter(10)
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

}*/
/*
func temperatureAdjuster(start float64) func(change float64) float64 {

	// Give me a function that remembers a temperature, and every time I call it, it updates that temperature.”

	baseTemprature := start

	adjust := func(change float64) float64 {
		baseTemprature = baseTemprature + change

		return float64(baseTemprature)
	}
	return adjust

}

func main() {
	adjust := temperatureAdjuster(50)

	fmt.Println(adjust(5))   // 55
	fmt.Println(adjust(-10)) // 45
}*/


