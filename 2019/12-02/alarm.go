package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	original := convert(strings.Split(scanner.Text(), ","))
	fmt.Println(original)

	for noun := 0; noun <= 100; noun++ {
		for verb := 0; verb <= 100; verb++ {
			s := make([]int, len(original))
			arr := copy(s, original)
			s[1] = noun
			s[2] = verb
			fmt.Println(arr)
			result := Compute(s)
			if result[0] == 19690720 {
				fmt.Println("Got result 19690720")
				fmt.Printf("Noun: %d, Verb %d \n", noun, verb)
				fmt.Printf("100 * noun + verb? %d", 100*noun+verb)
				os.Exit(1)
			}
		}
	}

}

func Compute(arr []int) []int {
	begin := 0

	for {

		inputA := arr[begin+1]
		inputB := arr[begin+2]
		resultI := arr[begin+3]

		if arr[begin] == 1 {
			arr[resultI] = arr[inputA] + arr[inputB]
		} else if arr[begin] == 2 {
			arr[resultI] = arr[inputA] * arr[inputB]
		} else {
			fmt.Printf("What is this? %d at position %d \n", arr[begin], begin)
			break
		}
		begin += 4
		if arr[begin] == 99 {
			break
		}
	}

	return arr
}

func convert(arr []string) []int {
	var r = []int{}

	for _, i := range arr {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		r = append(r, j)
	}
	return r
}
