package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseInputFrames(frames string) [][]string {
	parts := strings.Split(frames, " ")
	output := [][]string{}

	for _, part := range parts {
		frame := strings.Split(part, "")
		output = append(output, frame)
	}
	return output
}

func GetRollScoreAndBonus(roll string, frame_score int) (int, int) {
	roll_score := 0
	bonus_rolls := 0
	if roll == "X" {
		roll_score = 10
		bonus_rolls = 2
	} else if roll == "/" {
		roll_score = 10 - frame_score
		bonus_rolls = 1
	} else if roll == "-" {
		roll_score = 0
	} else {
		roll_score, _ = strconv.Atoi(roll)
	}

	return roll_score, bonus_rolls
}

func ApplyBonus(bonus []int, roll_score int) int {
	addt := 0
	for i := range bonus {
		if bonus[i] > 0 {
			addt += roll_score
			bonus[i]--
		}
	}
	return addt
}

func GetScore(score_report string) int {
	frames := ParseInputFrames(score_report)

	if len(frames) != 10 {
		fmt.Println("INVALID SCORE")
		return -1
	}

	bonus := []int{}
	score := 0

	for _, frame := range frames[:9] {
		frame_score := 0
		for _, roll := range frame {
			roll_score, bonus_rolls := GetRollScoreAndBonus(roll, frame_score)

			frame_score += roll_score

			score += ApplyBonus(bonus, roll_score)
			bonus = append(bonus, bonus_rolls)
		}
		score += frame_score
	}

	// Do complicated 10th frame
	frame_score := 0
	for r, roll := range frames[9] {
		roll_score, bonus_rolls := GetRollScoreAndBonus(roll, frame_score)
		if r == 0 {
			// Apply bonuses as normal and apply bonus for the 10th roll
			score += ApplyBonus(bonus, roll_score)
			bonus = []int{bonus_rolls}
		} else if r == 1 {
			// Apply bonus but only to previous roll
			score += ApplyBonus(bonus, roll_score)
		}
		frame_score += roll_score
		score += roll_score
	}

	return score
}

func main() {
	inputs := []string{
		"X X X X X X X X X XXX",
		"X -/ X 5- 8/ 9- X 81 1- 4/X",
		"62 71 X 9- 8/ X X 35 72 5/8",
	}

	for i := range inputs {
		fmt.Println(GetScore(inputs[i]))
	}
}
