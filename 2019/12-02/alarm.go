package main

import (
	"os"
	"log"
	"bufio"
	"strings"
	"fmt"
	"strconv"
)

func main()  {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	arr := convert(strings.Split(scanner.Text(), ","))

	res := Compute((arr))
	fmt.Println(res[0])

}

func Compute(arr []int) []int{
	begin := 0

	for {
		
		inputA := arr[begin+1]
		inputB := arr[begin+2]
		resultI := arr[begin+3]

		fmt.Printf("Now reading at postion %d [ %d %d %d %d ] \n", begin, arr[begin], arr[begin+1], arr[begin+2], arr[begin+3])
		if arr[begin] == 1 {
			arr[resultI] = arr[inputA] + arr[inputB]
		} else if arr[begin] == 2 {
			arr[resultI] = arr[inputA] * arr[inputB]
		} else {
			fmt.Printf("What is this? %d at position %d \n", arr[begin], begin)
		}
		begin += 4
		if arr[begin] == 99 { 
			fmt.Printf("Reached 99 at position %d \n", begin)
			break 
		}
	}

	return arr
}

func convert(arr []string) []int  {
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