package days

type InputType int

const (
	InputSample InputType = iota
	InputReal
)

type Part int

const (
	_ Part = iota
	Part1
	Part2
)

var (
	inputByDay     map[int]map[InputType][]byte
	solutionsByDay map[int]map[Part]Solution
)

func LookupInput(day int, inputType InputType) ([]byte, bool) {
	b, ok := inputByDay[day][inputType]
	return b, ok
}

type Solution func(input []byte) (string, error)

func LookupSolution(day int, part Part) (Solution, bool) {
	sln, ok := solutionsByDay[day][part]
	return sln, ok
}
