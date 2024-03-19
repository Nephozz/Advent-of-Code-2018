package days

import (
	"fmt"
	"os"
	"strings"
)

func Day2() {
	input, err := os.ReadFile("inputs/day2")
	if err != nil {
		fmt.Println("Error, couldn't read the file!: ", err)
		return
	}

	lines := strings.Split(string(input), "\n")
	var twoTimes, threeTimes int

	for _, line := range lines {
		letterCounts := make(map[rune]int)

		for _, letter := range line {
			letterCounts[letter]++
		}

		twoFound, threeFound := false, false

		for _, count := range letterCounts {
			if count == 2 && !twoFound {
				twoTimes++
				twoFound = true
			}
			if count == 3 && !threeFound {
				threeTimes++
				threeFound = true
			}
		}
	}

	checksum := twoTimes * threeTimes

	fmt.Println("Checksum is:", checksum)
}
