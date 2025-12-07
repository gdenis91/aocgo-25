package days

import (
	"github.com/gdenis91/aocgo-25/day1"
	"github.com/gdenis91/aocgo-25/day2"
)

func init() {
	inputByDay = map[int]map[InputType][]byte{
		1: {
			InputSample: day1.InputSample,
			InputReal:   day1.InputReal,
		},
		2: {
			InputSample: day2.InputSample,
			InputReal:   day2.InputReal,
		},
	}

	solutionsByDay = map[int]map[Part]Solution{
		1: {
			Part1: day1.Part1,
			Part2: day1.Part2,
		},
		2: {
			Part1: day2.Part1,
			Part2: day2.Part2,
		},
	}
}
