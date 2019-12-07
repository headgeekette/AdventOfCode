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
	answer1 := 0
	answer2 := 0
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		answer1 += Fuel(mass)
		answer2 += ExtFuel(mass)
	}

	fmt.Printf("Answer to first problem : %d \n", answer1)
	fmt.Printf("Answer to second problem : %d \n", answer2)

}

// Fuel computes for the fuel needed per single mass
func Fuel(mass int) int {
	return mass/3 - 2
}

// ExtFuel computes for the fuel needed per module
func ExtFuel(mass int) int {
	total := 0
	fuel := mass
	for fuel > 0 {
		fuel = Fuel(fuel)
		if fuel > 0 {
			total += fuel
		}
	}
	return total
}
