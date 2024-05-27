package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	data, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	sum := 0

	fileLines := regexp.MustCompile(`\r?\n`).Split(string(data), -1)

	for _, line := range fileLines {
		re := regexp.MustCompile("[0-9]+")
		arr := re.FindAllString(line, -1)

		fmt.Println(arr)
		int1 := arr[0]

		if len(int1) > 1 {
			temp := string(int1)
			int1 = string(temp[0])
		}
		int2 := arr[len(arr)-1]
		if len(int2) > 1 {
			temp := string(int2)
			int2 = string(temp[len(temp)-1])
		}

		resInt, err := strconv.Atoi(int1 + int2)

		if err != nil {
			return
		}
		sum = sum + resInt
	}

	println(sum)
}
