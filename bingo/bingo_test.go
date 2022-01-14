package bingo

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Bingo_Win_1(t *testing.T) {
	nums, boards, err := parseInput("./sample.txt")
	require.NoError(t, err)

	score := Win(boards, nums)
	require.Equal(t, 4512, score)
}

func Test_Bingo_Win_2(t *testing.T) {
	nums, boards, err := parseInput("./input.txt")
	require.NoError(t, err)

	score := Win(boards, nums)
	t.Logf("Score: %d", score)
}

func Test_Bingo_WinLast_1(t *testing.T) {
	nums, boards, err := parseInput("./sample.txt")
	require.NoError(t, err)

	score := WinLast(boards, nums)
	require.Equal(t, 1924, score)
}

func Test_Bingo_WinLast_2(t *testing.T) {
	nums, boards, err := parseInput("./input.txt")
	require.NoError(t, err)

	score := WinLast(boards, nums)
	t.Logf("Score: %d", score)
}

func parseInput(path string) (nums []int, boards []Board, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	br := bufio.NewReader(f)
	// nums
	s, err := br.ReadString('\n')
	if err != nil {
		return nil, nil, err
	}
	ss := strings.Split(strings.TrimSpace(s), ",")
	nums = make([]int, 0, len(ss))
	for _, s := range ss {
		i, _ := strconv.Atoi(s)
		nums = append(nums, i)
	}

	// boards
	boards = make([]Board, 0)
	var board = Board{
		Data: make([][]int, 0, 5),
	}
	for {
		s, err := br.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, nil, err
		}

		// empty line
		if strings.TrimSpace(s) == "" {
			continue
		}

		ss := strings.Split(strings.TrimSpace(s), " ")
		nn := make([]int, 0, len(ss))
		for _, s := range ss {
			if i, err := strconv.Atoi(s); err == nil {
				nn = append(nn, i)
			}
		}

		board.Data = append(board.Data, nn)
		if len(board.Data) == 5 {
			boards = append(boards, board)
			board = Board{
				Data: make([][]int, 0, 5),
			}
		}
	}
	return nums, boards, nil
}

func Test_calcMinIndex(t *testing.T) {
	nums := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	samples := []int{0, 13, 7, 10, 16}

	require.Equal(t, 14, calcMinIndex(samples, nums))
}
