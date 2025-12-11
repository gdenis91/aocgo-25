package day6

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input []byte) (string, error) {
	var inputRows [][]int
	var operators []string
	var linePos int
	for l := range strings.SplitSeq(string(input), "\n") {
		if linePos == 4 {
			operators = strings.Fields(l)
			break
		}

		var row []int
		for vStr := range strings.FieldsSeq(l) {
			v, err := strconv.Atoi(vStr)
			if err != nil {
				return "", fmt.Errorf("atoi: %w", err)
			}
			row = append(row, v)
		}
		inputRows = append(inputRows, row)
		linePos++
	}

	answers := make([]int, len(operators))
	for i, op := range operators {
		for j, row := range inputRows {
			if j == 0 {
				answers[i] = row[i]
				continue
			}
			switch op {
			case "+":
				answers[i] = row[i] + answers[i]
			case "*":
				answers[i] = row[i] * answers[i]
			default:
				panic(fmt.Errorf("unknown operator: %s", op))
			}
		}
	}

	var sln int
	for _, v := range answers {
		sln += v
	}

	return fmt.Sprint(sln), nil
}

type question struct {
	inputRows [][]rune
	width     int
	op        rune
}

func (q question) solve() (int, error) {
	var inputs []int
	for i := range q.width {
		var parts []rune
		for _, row := range q.inputRows {
			if row[i] == ' ' {
				continue
			}
			parts = append(parts, row[i])
		}
		v, err := strconv.Atoi(string(parts))
		if err != nil {
			return 0, fmt.Errorf("atoi %s: %w", string(parts), err)
		}
		inputs = append(inputs, v)
	}

	var answer int
	for i, v := range inputs {
		if i == 0 {
			answer = v
			continue
		}
		switch q.op {
		case '+':
			answer = v + answer
		case '*':
			answer = v * answer
		default:
			panic(fmt.Errorf("unknown operator: %s", string(q.op)))
		}
	}

	return answer, nil
}

func Part2(input []byte) (string, error) {
	var questions []question
	var q question
	var questionStart int
	lines := strings.Split(string(input), "\n")
	for i, v := range lines[len(lines)-1] {
		if v != ' ' {
			if q.op == 0 {
				// First question
				questionStart = i
				q.op = v
				continue
			}
			// End of question
			q.inputRows = make([][]rune, len(lines)-1)
			for j := range q.inputRows {
				q.inputRows[j] = make([]rune, q.width)
			}

			for j, l := range lines[:len(lines)-1] {
				for k := range q.width {
					q.inputRows[j][k] = []rune(l)[questionStart+k]
				}
			}
			questions = append(questions, q)
			questionStart = i
			q = question{
				op: v,
			}
			continue
		}
		q.width++
	}
	// Final question
	q.width++
	q.inputRows = make([][]rune, len(lines)-1)
	for j := range q.inputRows {
		q.inputRows[j] = make([]rune, q.width)
	}

	for j, l := range lines[:len(lines)-1] {
		for k := range q.width {
			q.inputRows[j][k] = []rune(l)[questionStart+k]
		}
	}
	questions = append(questions, q)

	var sln int
	for _, q := range questions {
		v, err := q.solve()
		if err != nil {
			return "", fmt.Errorf("solve question %v: %w", q, err)
		}
		sln += v
	}

	return fmt.Sprint(sln), nil
}
