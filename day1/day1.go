package day1

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func SolveDay1() {
	calories := readInputIntoMemory()
	calories.Sort()
	lenElves := calories.Len()
	fmt.Printf("The top Elf is carrying %d calories in total\n", calories[lenElves-1])
	topThreeCaloriesElves := calories[lenElves-1] +
		calories[lenElves-2] +
		calories[lenElves-3]
	fmt.Printf("The top three Elves are carrying %d calories in total\n", topThreeCaloriesElves)
}

func readInputIntoMemory() (calories sort.IntSlice) {
	calories = make(sort.IntSlice, 0)
	inputFile, err := os.Open("day1/input1.txt")
	if err != nil {
		fmt.Println("Error reading the input file for day 1 with error: ", err)
		return
	}
outer:
	for {
		var sumCalories int
	inner:
		for {
			var s string
			chars_read, err := fmt.Fscanln(inputFile, &s)
			if err == io.EOF {
				break outer
			}
			if chars_read == 0 {
				break inner
			}

			calorie, err := strconv.Atoi(s)
			sumCalories += calorie
			if err != nil {
				fmt.Println("Errored while converting input to int")
			}
		}
		calories = append(calories, sumCalories)
	}
	return
}
