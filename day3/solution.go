package day3

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input []byte) (string, error) {
	var sln int
	for l := range strings.SplitSeq(string(input), "\n") {
		leftValIdx, rightValIdx := len(l)-2, len(l)-1
		for i := leftValIdx - 1; i >= 0; i-- {
			if l[i] >= l[leftValIdx] {
				if l[leftValIdx] > l[rightValIdx] {
					rightValIdx = leftValIdx
				}
				leftValIdx = i
			}
		}
		v, err := strconv.Atoi(
			string(l[leftValIdx]) + string(l[rightValIdx]),
		)
		if err != nil {
			return "", fmt.Errorf("atoi: %w", err)
		}
		sln += v
	}
	return fmt.Sprint(sln), nil
}

func Part2(input []byte) (string, error) {
	var sln int
	for l := range strings.SplitSeq(string(input), "\n") {
		bank12 := len(l) - 1
		bank11 := len(l) - 2
		bank10 := len(l) - 3
		bank9 := len(l) - 4
		bank8 := len(l) - 5
		bank7 := len(l) - 6
		bank6 := len(l) - 7
		bank5 := len(l) - 8
		bank4 := len(l) - 9
		bank3 := len(l) - 10
		bank2 := len(l) - 11
		bank1 := len(l) - 12
		for i := bank1 - 1; i >= 0; i-- {
			if l[i] >= l[bank1] {
				if l[bank1] >= l[bank2] {
					if l[bank2] >= l[bank3] {
						if l[bank3] >= l[bank4] {
							if l[bank4] >= l[bank5] {
								if l[bank5] >= l[bank6] {
									if l[bank6] >= l[bank7] {
										if l[bank7] >= l[bank8] {
											if l[bank8] >= l[bank9] {
												if l[bank9] >= l[bank10] {
													if l[bank10] >= l[bank11] {
														if l[bank11] > l[bank12] {
															bank12 = bank11
														}
														bank11 = bank10
													}
													bank10 = bank9
												}
												bank9 = bank8
											}
											bank8 = bank7
										}
										bank7 = bank6
									}
									bank6 = bank5
								}
								bank5 = bank4
							}
							bank4 = bank3
						}
						bank3 = bank2
					}
					bank2 = bank1
				}
				bank1 = i
			}
		}
		v, err := strconv.Atoi(
			string(l[bank1]) +
				string(l[bank2]) +
				string(l[bank3]) +
				string(l[bank4]) +
				string(l[bank5]) +
				string(l[bank6]) +
				string(l[bank7]) +
				string(l[bank8]) +
				string(l[bank9]) +
				string(l[bank10]) +
				string(l[bank11]) +
				string(l[bank12]),
		)
		if err != nil {
			return "", fmt.Errorf("atoi: %w", err)
		}
		sln += v
	}
	return fmt.Sprint(sln), nil
}
