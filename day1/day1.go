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
	totalCaloriePerElf := getSortedTotalCaloriesPerElf(calories)
	lenElves := totalCaloriePerElf.Len()
	fmt.Printf("The top Elf is carrying %d calories in total\n", totalCaloriePerElf[lenElves-1])
	topThreeCaloriesElves := totalCaloriePerElf[lenElves-1] +
		totalCaloriePerElf[lenElves-2] +
		totalCaloriePerElf[lenElves-3]
	fmt.Printf("The top three Elves are carrying %d calories in total\n", topThreeCaloriesElves)
}

func readInputIntoMemory() (calories [][]int) {
	calories = make([][]int, 0)
	// i is for diff people
	var i int
	inputFile, err := os.Open("day1/input1.txt")
	if err != nil {
		fmt.Println("Error reading the input file for day 1!")
		fmt.Println(err.Error())
		return
	}
outer:
	for {
		calories = append(calories, make([]int, 0))
	inner:
		for {
			var s string
			chars_read, err := fmt.Fscanln(inputFile, &s)
			if err == io.EOF {
				break outer
			}
			if chars_read == 0 {
				i++
				break inner
			}

			calorie, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("Errored while converting to int")
			}
			calories[i] = append(calories[i], calorie)
		}
	}
	return
}

func getSortedTotalCaloriesPerElf(calories [][]int) (totalCaloriePerElf sort.IntSlice) {
	totalCaloriePerElf = make(sort.IntSlice, 0)
	for i := range calories {
		var sum int = 0
		for j := range calories[i] {
			sum += calories[i][j]
		}
		totalCaloriePerElf = append(totalCaloriePerElf, sum)
	}
	totalCaloriePerElf.Sort()
	return
}
