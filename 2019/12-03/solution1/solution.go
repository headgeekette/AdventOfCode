package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func main() {
	wire1, wire2 := readInput("../input.txt")
	wirepath1 := GetPath(wire1)
	wirepath2 := GetPath(wire2)

	omh := 0
	for p1 := 1; p1 < len(wirepath1); p1++ {
		for p2 := 1; p2 < len(wirepath2); p2++ {
			// Figure out first if the wires will intersect
			if DoCross(wirepath1[p1-1], wirepath1[p1], wirepath2[p2-1], wirepath2[p2]) {
				// Then get the point of intersection
				intersect := GetCross(wirepath1[p1-1], wirepath1[p1], wirepath2[p2-1], wirepath2[p2])
				// And compute for the Manhattan distance.
				nmh := abs(intersect.X) + abs(intersect.Y)
				if omh == 0 || nmh < omh{
					omh = nmh
				}
			}
		}
	}
	fmt.Printf("Manhattan Distance %v \n", omh)
}

func abs(x int) int {
	// yeah. Go doesn't have a function that would get the absolute value of an int.
	if x < 0 {
		return -x
	}
	return x
}

func readInput(name string) ([]string, []string) {
	file, err := os.Open(name)
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

	return wire1, wire2
}

// GetPath gets the actual coordinates of the wire path
func GetPath(in []string) []Point {
	pp := make([]Point, len(in)+1)

	for i := 0; i < len(in); i++ {
		p := GetCoor(in[i])
		pp[i+1].X = pp[i].X + p.X
		pp[i+1].Y = pp[i].Y + p.Y
	}
	return pp
}

// GetCoor converts a string (e.g., R75) to a coordinate (e.g., [75,0])
// I'm sure there's a faster way to do this part
func GetCoor(input string) Point {

	var p Point
	// split the string
	in := strings.SplitAfterN(input, "", 2)
	v, _ := strconv.Atoi(in[1])

	switch in[0] {
	case "R":
		p.X = v
		break
	case "L":
		p.X = -v
		break
	case "U":
		p.Y = v
		break
	case "D":
		p.Y = -v
	default:
		break
	}
	return p
}

// DoCross is used to determine if the wires do intersect based on their orientation to each other.
// What we want is that the wires aren't colinear.
func DoCross(p1 Point, q1 Point, p2 Point, q2 Point) bool {

	// Are the ends of wire B to the left and right of wire A?
	o1 := getOrient(p1, q1, p2)
	o2 := getOrient(p1, q1, q2)
	// Are the ends of wire A to the left and right of wire B?
	o3 := getOrient(p2, q2, p1)
	o4 := getOrient(p2, q2, q1)

	if o1 != o2 && o3 != o4 {
		return true
	}

	return false
}

func getOrient(p1 Point, q1 Point, v Point) int {

	r := (q1.Y-p1.Y)*(v.X-q1.X) - (q1.X-p1.X)*(v.Y-q1.Y)

	if r == 0 {
		return 0 // colinear
	} else if r > 0 {
		return 1 // point is to the right of the line
	} else {
		return 2 // point is to the left of the line
	}

}

// GetCross returns the point of intersection of the two wires.
func GetCross(p1 Point, q1 Point, p2 Point, q2 Point) Point {

	// wire1
	a1 := q1.Y - p1.Y
	b1 := p1.X - q1.X
	c1 := a1*(p1.X) + b1*(p1.Y)

	// wire2
	a2 := q2.Y - p2.Y
	b2 := p2.X - q2.X
	c2 := a2*(p2.X) + b2*(p2.Y)

	d := a1*b2 - a2*b1

	x := (b2*c1 - b1*c2) / d
	y := (a1*c2 - a2*c1) / d

	return Point{x, y}
}
