package bingo

import (
	"math"
)

type Board struct {
	Data [][]int
}

type MarkType int

const (
	Row    MarkType = 0
	Column MarkType = 1
)

type WinMark struct {
	MarkType
	Index int
}

func (b *Board) Wins(nums []int) (WinMark, int) {
	var cnt = math.MaxInt
	var mark WinMark

	n := len(b.Data)
	// by rows
	for i := 0; i < n; i++ {
		samples := b.Data[i]

		if idx := calcMinIndex(samples, nums); idx < cnt {
			cnt = idx
			mark.MarkType = Row
			mark.Index = i
		}
	}
	// by columns
	for i := 0; i < n; i++ {
		samples := make([]int, 0, n)
		for ii := 0; ii < n; ii++ {
			samples = append(samples, b.Data[ii][i])
		}

		if idx := calcMinIndex(samples, nums); idx < cnt {
			cnt = idx
			mark.MarkType = Column
			mark.Index = i
		}
	}
	return mark, cnt
}

func (b *Board) UnmarkedSum(nums []int) int {
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		m[num] = true
	}

	n := len(b.Data)
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !m[b.Data[i][j]] {
				sum += b.Data[i][j]
			}
		}
	}
	return sum
}

func calcMinIndex(samples []int, nums []int) int {
	var index int
	for i := 0; i < len(samples); i++ {
		var idx = -1
		for j := 0; j < len(nums); j++ {
			if nums[j] == samples[i] {
				idx = j
				break
			}
		}

		if idx == -1 {
			return math.MaxInt
		}

		if idx > index {
			index = idx
		}
	}
	return index
}

func Win(boards []Board, nums []int) (score int) {
	var (
		winCnt     = math.MaxInt
		boardIndex int
	)
	for i := range boards {
		if _, cnt := boards[i].Wins(nums); cnt < winCnt {
			winCnt = cnt
			boardIndex = i
		}
	}

	unmarkedSum := boards[boardIndex].UnmarkedSum(nums[:winCnt+1])
	return unmarkedSum * nums[winCnt]
}

func WinLast(boards []Board, nums []int) (score int) {
	var (
		winCnt     = -1
		boardIndex int
	)
	for i := range boards {
		if _, cnt := boards[i].Wins(nums); cnt > winCnt {
			winCnt = cnt
			boardIndex = i
		}
	}

	unmarkedSum := boards[boardIndex].UnmarkedSum(nums[:winCnt+1])
	return unmarkedSum * nums[winCnt]
}
