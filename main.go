package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Creating a new scanner to read from console input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter a temperature (32 F or 100 C): or type q to quit ")
		scanner.Scan() //read user input

		temperature := scanner.Text() // user input is stored as a string

		//strings.Fields splits the user input into temperature value and unit of temperatur based on white space
		splitInput := strings.Fields(temperature)
		if strings.ToLower(temperature) == "q" {
			break
		}

		// fmt.Println(splitInput)
		if len(splitInput) != 2 {
			fmt.Println("Invalid input format. Please enter temperture unit (F or C).")
			continue
		}

		// get the unit of temperture
		unitOfTemperture := strings.ToUpper(splitInput[1])

		//convert the string (temperature value) to a floating-point number
		tempertureToConvert, err := strconv.ParseFloat(splitInput[0], 64)

		if err != nil {
			fmt.Println("Invalid number. Please enter a valid numeric value.")
			continue
		}

		// using if statement to check the unit of the temperture
		if unitOfTemperture == "F" {
			// Calling FahrenheitToCelsius to convert fahrenheit into celsius
			celsius := FahrenheitToCelsius(tempertureToConvert)
			fmt.Printf("Your Celsius value is %.2f C\n", celsius)

		} else if unitOfTemperture == "C" {
			fahrenheit := CelsiusToFahrenheit(tempertureToConvert)
			// Calling FahrenheitToCelsius to convert celsius into fahrenheit
			fmt.Printf("Your Fahrenheit value is %.2f F\n", fahrenheit)

		} else {
			fmt.Println("Invalid unit. Please use 'F' for Fahrenheit or 'C' for Celsius.")
		}

	}
}

// Function to convert fahrenheit into celsius
func FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

// Function to convert celsius into fahrenheit
func CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}
