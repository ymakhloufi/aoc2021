package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	bitStreams, err := readLines("assets/input")
	if err != nil {
		fmt.Printf("error reading file: %s\n", err.Error())
		return
	}

	mostCommonBits, err := getMostCommonBits(bitStreams)
	if err != nil {
		fmt.Printf("failed determining most common bits: %s\n", err.Error())
		return
	}
	decimal, err := boolSliceToInt(mostCommonBits)
	if err != nil {
		fmt.Printf("failed to convert boolSlice to int: %s\n", err.Error())
	}
	fmt.Printf("The most common bits are %v, which is in decimals: %d\n", mostCommonBits, decimal)

	leastCommonBits := invertBits(mostCommonBits)
	decimal, err = boolSliceToInt(leastCommonBits)
	if err != nil {
		fmt.Printf("failed to convert boolSlice to int: %s\n", err.Error())
	}
	fmt.Printf("The lease common bits are %v, which is in decimals: %d\n", leastCommonBits, decimal)

	oxigen, co2scrubber, err := filterByCommonality(bitStreams)
	if err != nil {
		fmt.Printf("failed to calculate life support metrics: %s\n", err.Error())
	}
	fmt.Printf("Commonality:\n%v\n%v\n", oxigen, co2scrubber)
}

func filterByCommonality(streams [][]bool) (oxigen [][]bool, co2Scrubber [][]bool, err error) {
	remainingOxygenStreams := make([][]bool, len(streams))
	copy(remainingOxygenStreams, streams)

	remainingCo2ScrubberStreams := make([][]bool, len(streams))
	copy(remainingCo2ScrubberStreams, streams)

	streamLength := len(streams[0])
	for i := 0; i < streamLength; i++ {
		mostCommonBits, err := getMostCommonBits(remainingOxygenStreams)
		fmt.Println(mostCommonBits)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to determine most common bits for %v: %s", streams, err)
		}

		remainingOxygenStreams = filterByPositionValue(remainingOxygenStreams, i, mostCommonBits[i])
		fmt.Println(remainingOxygenStreams)
	}
	return remainingOxygenStreams, remainingCo2ScrubberStreams, nil
}

func filterByPositionValue(streams [][]bool, pos int, value bool) [][]bool {
	keptStreams := [][]bool{}
	for _, stream := range streams {
		if stream[pos] == value {
			keptStreams = append(keptStreams, stream)
		}
	}
	return keptStreams
}

func getMostCommonBits(streams [][]bool) ([]bool, error) {
	positionSums := make([]uint, len(streams[0]))
	for _, bitStream := range streams {
		for idx, bit := range bitStream {
			if bit {
				if len(positionSums) > idx {
					positionSums[idx]++
				}
			}
		}
	}

	streamCount := len(streams)
	halfStreamCount := uint(streamCount / 2)
	mostCommonBits := make([]bool, len(positionSums))
	for idx, sum := range positionSums {
		mostCommonBits[idx] = sum >= halfStreamCount
	}

	return mostCommonBits, nil
}

func invertBits(bits []bool) []bool {
	invertedBits := make([]bool, len(bits))
	for idx, bit := range bits {
		invertedBits[idx] = !bit
	}
	return invertedBits
}

func boolSliceToInt(bits []bool) (int64, error) {
	str := ""
	for _, bit := range bits {
		if bit {
			str += "1"
		} else {
			str += "0"
		}
	}
	return strconv.ParseInt(str, 2, 64)
}

func readLines(path string) ([][]bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()

	var lines [][]bool
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		i++
		bits := []bool{}
		for _, rune := range line {
			switch rune {
			case '1':
				bits = append(bits, true)
			case '0':
				bits = append(bits, false)
			default:
				return nil, fmt.Errorf("could not convert the tune %s into bit", rune)
			}
		}
		lines = append(lines, bits)
	}
	return lines, scanner.Err()
}
