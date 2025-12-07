package day2

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Part1(input []byte) (string, error) {
	var sln int
	for idRange := range strings.SplitSeq(string(input), ",") {
		idRangeVals := strings.Split(idRange, "-")
		startID, err := strconv.Atoi(idRangeVals[0])
		if err != nil {
			return "", fmt.Errorf("atoi start id: %w", err)
		}
		endID, err := strconv.Atoi(idRangeVals[1])
		if err != nil {
			return "", fmt.Errorf("atoi end id: %w", err)
		}

		for id := startID; id <= endID; id++ {
			idStr := fmt.Sprint(id)
			if idStr[:len(idStr)/2] == idStr[len(idStr)/2:] {
				sln += id
			}
		}
	}
	return fmt.Sprint(sln), nil
}

func Part2(input []byte) (string, error) {
	var sln int
	for idRange := range strings.SplitSeq(string(input), ",") {
		idRangeVals := strings.Split(idRange, "-")
		startID, err := strconv.Atoi(idRangeVals[0])
		if err != nil {
			return "", fmt.Errorf("atoi start id: %w", err)
		}

		endID, err := strconv.Atoi(idRangeVals[1])
		if err != nil {
			return "", fmt.Errorf("atoi end id: %w", err)
		}

		for id := startID; id <= endID; id++ {
			// single ids can't repeat
			if id < 10 {
				continue
			}
			idRunes := []rune(fmt.Sprint(id))
			seq := idRunes[:1]
			// only check seq up to half the len since sequences
			// longer can't repeat
			for len(seq) <= len(idRunes)/2 {
				// if the seq can't fit nicely into the id we can skip
				if len(idRunes)%len(seq) == 0 {
					invalid := true
					// compare seq slice to each chunk of idRunes
					for i := range len(idRunes) / len(seq) {
						if slices.Equal(seq, idRunes[i*len(seq):(i*len(seq))+len(seq)]) {
							continue
						}
						invalid = false
					}
					if invalid {
						sln += id
						break
					}
				}
				seq = idRunes[:len(seq)+1]
			}
		}
	}
	return fmt.Sprint(sln), nil
}
