package day2

import (
	"fmt"
	"io"
	"os"
)

func SolveDay2() {
	fmt.Println()
	fmt.Println("Solving Day 2!")
	score := calculateScoreForAssumedStrategy()
	fmt.Println("Part 1: The score with the assumed strategy is: ", score)
	score = calculateScoreFromActualStrategy()
	fmt.Println("Part 1: The score with the actual strategy is: ", score)
	fmt.Println()
}

func calculateScoreForAssumedStrategy() (score int) {
	inputFile, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Println("Error opening the input file for day 2 (assumed strategy) with error: ", err)
	}
	for {
		var opponentInput, playerInput string
		_, err := fmt.Fscanln(inputFile, &opponentInput, &playerInput)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("Error while reading the input file (assumed strategy) with error: ", err)
		}
		var opponent, player = int(opponentInput[0] - 'A'), int(playerInput[0] - 'X')
		score += ((player - opponent + 4) % 3) * 3 // outcome of the round
		score += player + 1                        // score from the move
	}
}

func calculateScoreFromActualStrategy() (score int) {
	inputFile, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Println("Error opening the input file for day 2 (actual strategy) with error: ", err)
	}
	for {
		var opponentInput, playerInput string
		_, err := fmt.Fscanln(inputFile, &opponentInput, &playerInput)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("Error while reading the input file (actual strategy) with error: ", err)
		}
		var opponent, outcome = int(opponentInput[0] - 'A' + 3), int(playerInput[0] - 'X')
		score += outcome * 3                        // outcome of the round
		score += ((opponent + outcome - 1) % 3) + 1 // score from the move
	}
}
