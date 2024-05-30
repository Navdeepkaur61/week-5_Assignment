package main

import (
	"bufio"
	"fmt"
	"os"
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

		fmt.Println(splitInput)
		if len(splitInput) != 2 {
			fmt.Println("Invalid input format. Please enter temperture unit (F or C).")
			continue
		}

	}
}
