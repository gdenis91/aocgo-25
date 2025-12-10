package day5

import (
	"fmt"
	"strconv"
	"strings"
)

type idRange struct {
	lower int
	upper int
}

func Part1(input []byte) (string, error) {
	readIDRanges := true
	idRanges := make([]idRange, 0)
	var sln int
	for l := range strings.SplitSeq(string(input), "\n") {
		if l == "" {
			readIDRanges = false
			continue
		}
		if readIDRanges {
			parts := strings.Split(l, "-")
			lower, err := strconv.Atoi(parts[0])
			if err != nil {
				return "", fmt.Errorf("atoi lower: %w", err)
			}
			upper, err := strconv.Atoi(parts[1])
			if err != nil {
				return "", fmt.Errorf("atoi upper: %w", err)
			}
			idRanges = append(idRanges, idRange{lower, upper})
			continue
		}
		v, err := strconv.Atoi(l)
		if err != nil {
			return "", fmt.Errorf("atoi id: %w", err)
		}
		for _, r := range idRanges {
			if v >= r.lower &&
				v <= r.upper {
				sln++
				break
			}
		}
	}
	return fmt.Sprint(sln), nil
}

func Part2(input []byte) (string, error) {
	idRanges := make([]idRange, 0)
	for l := range strings.SplitSeq(string(input), "\n") {
		if l == "" {
			break
		}

		parts := strings.Split(l, "-")
		lower, err := strconv.Atoi(parts[0])
		if err != nil {
			return "", fmt.Errorf("atoi lower: %w", err)
		}
		upper, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", fmt.Errorf("atoi upper: %w", err)
		}
		idRanges = append(idRanges, idRange{lower, upper})
	}

	for {
		var hadMerge bool
		idRanges, hadMerge = mergeIDRanges(idRanges)
		if !hadMerge {
			break
		}
	}

	var sln int
	for _, r := range idRanges {
		sln += (r.upper - r.lower) + 1
	}
	return fmt.Sprint(sln), nil
}

// naively merge idRanges by iterating, starting with the first
// range and merging in the first possible range into the first
// range that can be merged and return. If nothing can be merged
// the same ranges are returned. A bool is returned indicating
// if any merges happened
func mergeIDRanges(idRanges []idRange) ([]idRange, bool) {
	var newIDRanges []idRange
	var hadMerge bool
	var mergedRange idRange
	mergedIndexes := make(map[int]struct{})
	for i, r := range idRanges {
		if !hadMerge {
			for j, r2 := range idRanges[i+1:] {
				if hadMerge {
					if max(mergedRange.lower, r2.lower) <= min(mergedRange.upper, r2.upper) {
						mergedRange = idRange{
							min(mergedRange.lower, r2.lower),
							max(mergedRange.upper, r2.upper),
						}
						mergedIndexes[i+j+1] = struct{}{}
					}
					continue
				}

				if max(r.lower, r2.lower) <= min(r.upper, r2.upper) {
					hadMerge = true
					mergedRange = idRange{
						min(r.lower, r2.lower),
						max(r.upper, r2.upper),
					}
					mergedIndexes[i+j+1] = struct{}{}
					mergedIndexes[i] = struct{}{}
				}
			}
		}

		if hadMerge {
			if _, ok := mergedIndexes[i]; !ok {
				newIDRanges = append(newIDRanges, r)
			}
		} else {
			newIDRanges = append(newIDRanges, r)
		}
	}
	if hadMerge {
		newIDRanges = append(newIDRanges, mergedRange)
	}
	return newIDRanges, hadMerge
}
