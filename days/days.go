package days

import (
	"github.com/gdenis91/aocgo-25/day1"
	"github.com/gdenis91/aocgo-25/day2"
	"github.com/gdenis91/aocgo-25/day3"
	"github.com/gdenis91/aocgo-25/day4"
	"github.com/gdenis91/aocgo-25/day5"
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
		3: {
			InputSample: day3.InputSample,
			InputReal:   day3.InputReal,
		},
		4: {
			InputSample: day4.InputSample,
			InputReal:   day4.InputReal,
		},
		5: {
			InputSample: day5.InputSample,
			InputReal:   day5.InputReal,
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
		3: {
			Part1: day3.Part1,
			Part2: day3.Part2,
		},
		4: {
			Part1: day4.Part1,
			Part2: day4.Part2,
		},
		5: {
			Part1: day5.Part1,
			Part2: day5.Part2,
		},
	}
}
