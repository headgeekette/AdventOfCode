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

type Points struct {
	P    Point
}

func main() {
	wire1, wire2 := readInput()

	wirepath1 := GetPath(wire1)
	wirepath2 := GetPath(wire2)

	md := 0 // The Manhattan Distance
	for i := 0; i < len(wirepath1); i++ {
		for j := 0; j < len(wirepath2); j++ {
			if wirepath1[i].P.X == wirepath2[j].P.X && wirepath1[i].P.Y == wirepath2[j].P.Y {
				currMd := abs(wirepath1[i].P.X) + abs(wirepath1[i].P.Y)
				if md == 0||currMd < md {
					md = currMd
				}
			}
		}
	}
	fmt.Printf("Manhattan Distance: %v \n", md)
}

// GetPath all the coordinates of the wire path
func GetPath(in []string) []Points {
	pp := make([]Points, 0)

	prev := Point{0, 0}

	for _, v := range in {
		curr := GetCoor(v)
		// fmt.Printf("curr %v \n", curr)
		if curr.X == 0 {
			actual := 0
			if curr.Y < 0 {
				actual = curr.Y - 1
			}
			for i := 1; i <= abs(curr.Y); i++ {
				p := Point{prev.X, prev.Y + i + actual}
				// fmt.Println(p)
				pp = append(pp, Points{p})
			}
		} else if curr.Y == 0 {
			actual := 0
			if curr.X < 0 {
				actual = curr.X - 1
			}
			for i := 1; i <= abs(curr.X); i++ {
				p := Point{prev.X + i + actual, prev.Y}
				// fmt.Println(p)
				pp = append(pp, Points{p})
			}
		}
		prev = Point{prev.X + curr.X, prev.Y + curr.Y}
		// fmt.Printf("New Prev::: %v \n", prev)
	}
	return pp
}

// GetCoor converts a string (e.g., R75) to a coordinate (e.g., [75,0])
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func readInput() ([]string, []string) {
	file, err := os.Open("../input.txt")
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
