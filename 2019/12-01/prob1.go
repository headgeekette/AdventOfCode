package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var subtotal int
	var grandtotal int
	for scanner.Scan() {
		input, _ := strconv.Atoi(scanner.Text())
		fuel := Fuel(input)
		subtotal += fuel
		grandtotal += ExtFuel(fuel)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Initial fuel needed ::: %d \n", subtotal)
	fmt.Printf("Total fuel needed   ::: %d \n", grandtotal)
}

// Fuel computes for the initial fuel needed according to problem 1.
func Fuel(mass int) int {
	result := mass/3 - 2
	return result
}

// ExtFuel is short for "Extended Fuel". This provides the answer to problem 2.
func ExtFuel(fuel int) int {
	final := fuel
	result := fuel
	for result > 0 {
		result = Fuel(result)
		if result > 0 {
			final += result
		}
	}
	return final
}
