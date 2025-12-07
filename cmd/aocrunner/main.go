package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gdenis91/aocgo-25/days"
)

func main() {
	dayFlag := flag.Int("day", 0, "The day to run")
	partFlag := flag.Int("part", 0, "The part to run")
	sampleData := flag.Bool(
		"sample",
		false,
		"When the sample flag is provided the sample input is used",
	)
	flag.Parse()

	if *dayFlag == 0 {
		fmt.Println("No day specified")
		os.Exit(1)
	}
	day := *dayFlag

	part := days.Part(*partFlag)
	switch part {
	case 1, 2:
	default:
		fmt.Println("Part flag must be 1 or 2")
		os.Exit(1)
	}

	inputType := days.InputReal
	if *sampleData {
		inputType = days.InputSample
	}

	input, ok := days.LookupInput(day, inputType)
	if !ok {
		fmt.Printf("No input for day %d input %d\n", day, inputType)
		os.Exit(1)
	}

	sln, ok := days.LookupSolution(day, part)
	if !ok {
		fmt.Printf("No solution for day %d part %d\n", day, part)
		os.Exit(1)
	}

	out, err := sln(input)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Solution:", out)
}
