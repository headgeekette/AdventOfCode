package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const x int = 0
const y int = 1

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// scan first line
	scanner.Scan()
	wire1 := strings.Split(scanner.Text(), ",")
	// scan second line
	scanner.Scan()
	wire2 := strings.Split(scanner.Text(), ",")

	// get wire1 path
	wirepath1 := GetPath(wire1)
	wirepath2 := GetPath(wire2)
	fmt.Println(wirepath1)
	fmt.Println(wirepath2)

	// Test for boundary.
	// https://martin-thoma.com/how-to-check-if-two-line-segments-intersect/
	// https://www.geeksforgeeks.org/check-if-two-given-line-segments-intersect/
	// https://blogs.sas.com/content/iml/2018/07/09/intersection-line-segments.html

	// If withing bounds, find intersection
	// https://www.geeksforgeeks.org/program-for-point-of-intersection-of-two-lines/

	// Find manhattan distance for closest intersection
}

// GetPath gets the actual coordinates of the wire path
func GetPath(in []string) [][]int {
	result := make([][]int, len(in)+1)
	result[0] = []int{0, 0}

	for i := 0; i < len(in); i++ {
		coor := GetCoor(in[i])
		xval := result[i][x] + coor[x]
		yval := result[i][y] + coor[y]
		result[i+1] = []int{xval, yval}
	}
	return result
}

// GetCoor converts a string (e.g., R75) to a coordinate (e.g., [75,0])
func GetCoor(input string) []int {

	result := []int{0, 0}
	// split the string
	in := strings.SplitAfterN(input, "", 2)
	v, _ := strconv.Atoi(in[1])

	switch in[0] {
	case "R":
		result[x] = v
		break
	case "L":
		result[x] = -v
		break
	case "U":
		result[y] = v
		break
	case "D":
		result[y] = -v
	default:
		break
	}
	return result
}
