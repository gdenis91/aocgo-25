package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input []byte) (string, error) {
	dialPos := 50
	var pwd int
	for l := range strings.SplitSeq(string(input), "\n") {
		newDialPos, _, err := rotate(dialPos, l)
		if err != nil {
			return "", fmt.Errorf("rotate: %w", err)
		}
		if newDialPos == 0 {
			pwd++
		}
		dialPos = newDialPos
	}
	return fmt.Sprint(pwd), nil
}

func Part2(input []byte) (string, error) {
	dialPos := 50
	var pwd int
	for l := range strings.SplitSeq(string(input), "\n") {
		newDialPos, passedZero, err := rotate(dialPos, l)
		if err != nil {
			return "", fmt.Errorf("rotate: %w", err)
		}
		fmt.Printf("rot: %s, pos: %d->%d, passed: %d\n", l, dialPos, newDialPos, passedZero)

		pwd += passedZero
		if newDialPos == 0 {
			pwd++
		}
		dialPos = newDialPos
		// fmt.Println(dialPos)
	}
	return fmt.Sprint(pwd), nil
}

func rotate(pos int, rotation string) (int, int, error) {
	initialPos := pos
	v, err := strconv.Atoi(rotation[1:])
	if err != nil {
		return 0, 0, fmt.Errorf("strconv rotation value %s: %w", rotation, err)
	}
	fullRots := v / 100
	rot := v % 100
	passedZero := fullRots
	if strings.HasPrefix(rotation, "L") {
		pos -= rot
		if pos < 0 {
			pos = 100 + pos
			if initialPos != 0 {
				passedZero++
			}
		}
	} else {
		pos += rot
		if pos > 100 {
			pos = pos - 100
			if initialPos != 0 {
				passedZero++
			}
		}
	}
	return pos % 100, passedZero, nil
}

// func rotateAlt(startingPos int, rotation string) (int, int, error) {
// 	rotAmount, err := strconv.Atoi(rotation[1:])
// 	if err != nil {
// 		return 0, 0, fmt.Errorf("strconv %s: %w", rotation, err)
// 	}

// 	pos := startingPos
// 	if strings.HasPrefix(rotation, "L") {
// 		pos -= rotAmount
// 	} else {
// 		pos += rotAmount
// 	}
// 	var passedZero int
// 	// If we rotated more than 100 we had to pass 0
// 	if rotAmount >= 100 {
// 		passedZero = rotAmount / 100
// 		// If we started and finished at 0, we don't want to count the final
// 		// rotation as passing 0, it's on 0
// 		if startingPos == 0 && mod(pos, 100) == 0 {
// 			passedZero--
// 		}
// 	} else if startingPos != 0 && (pos < 0 || pos > 100) {
// 		fmt.Println("pass!", pos)
// 		passedZero++
// 	}
// 	pos = mod(pos, 100)

// 	return pos, passedZero, nil

// 	// TODO: Break into full rotations and partial rotations
// 	// From there use the two pieces to figure out  how many times passed zero
// 	// Always return positive pos

// 	// If we're rotating less than 100 we need to check if we crossed zero
// 	if rotAmount < 100 &&
// 		startingPos != 0 &&
// 		(pos < 0 ||
// 			pos > 100) {
// 		fmt.Printf("Passing zero! (%d, %d)\n", pos, rotAmount)
// 		passedZero++
// 	}
// 	pos = mod(pos, 100)

// 	return pos, passedZero, nil
// }

// func mod(a, b int) int {
// 	return (a%b + b) % b
// }
