package diagnostic

import (
	"strconv"
)

func mostCommon(data [][]rune, index int) [][]rune {
	onePart := make([][]rune, 0, len(data))
	zeroPart := make([][]rune, 0, len(data))

	for _, d := range data {
		switch d[index] {
		case '1':
			onePart = append(onePart, d)
		case '0':
			zeroPart = append(zeroPart, d)
		}
	}

	if len(onePart) >= len(zeroPart) {
		return onePart
	} else {
		return zeroPart
	}
}

func leastCommon(data [][]rune, index int) [][]rune {
	onePart := make([][]rune, 0, len(data))
	zeroPart := make([][]rune, 0, len(data))

	for _, d := range data {
		switch d[index] {
		case '1':
			onePart = append(onePart, d)
		case '0':
			zeroPart = append(zeroPart, d)
		}
	}

	if len(onePart) >= len(zeroPart) {
		return zeroPart
	} else {
		return onePart
	}
}

func OxygenCalc(data [][]rune) int64 {
	if len(data) == 0 {
		return 0
	}

	n := len(data[0])
	for i := 0; i < n; i++ {
		if len(data) == 1 {
			break
		}
		data = mostCommon(data, i)
	}

	i, _ := strconv.ParseInt(string(data[0]), 2, 64)
	return i
}

func CO2Calc(data [][]rune) int64 {
	if len(data) == 0 {
		return 0
	}

	n := len(data[0])
	for i := 0; i < n; i++ {
		if len(data) == 1 {
			break
		}
		data = leastCommon(data, i)
	}

	i, _ := strconv.ParseInt(string(data[0]), 2, 64)
	return i
}
