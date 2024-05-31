package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter a temperature (32 F or 100 C): or type y to quit ")
		scanner.Scan()

		temperature := scanner.Text()

		splitInput := strings.Fields(temperature)
		if strings.ToLower(temperature) == "y" {
			break
		}

		// fmt.Println(splitInput)
		if len(splitInput) != 2 {
			fmt.Println("Invalid input format. Please enter temperture unit (F or C).")
			continue
		}

		unitOfTemperture := strings.ToUpper(splitInput[1])

		tempertureToConvert, err := strconv.ParseFloat(splitInput[0], 64)

		if err != nil {
			fmt.Println("Invalid number. Please enter a valid numeric value.")
			continue
		}

		if unitOfTemperture == "F" {
			celsius := FahrenheitToCelsius(tempertureToConvert)
			fmt.Printf("Your Celsius value is %.2f C\n", celsius)
		} else if unitOfTemperture == "C" {
			fahrenheit := CelsiusToFahrenheit(tempertureToConvert)
			fmt.Printf("Your Fahrenheit value is %.2f F\n", fahrenheit)
		} else {
			fmt.Println("Invalid unit. Please use 'F' for Fahrenheit or 'C' for Celsius.")
		}

	}
}

func FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}
