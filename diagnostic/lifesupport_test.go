package diagnostic

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_LifeSupport_1(t *testing.T) {
	data := [][]rune{
		[]rune("00100"),
		[]rune("11110"),
		[]rune("10110"),
		[]rune("10111"),
		[]rune("10101"),
		[]rune("01111"),
		[]rune("00111"),
		[]rune("11100"),
		[]rune("10000"),
		[]rune("11001"),
		[]rune("00010"),
		[]rune("01010"),
	}

	oxygen, co2 := OxygenCalc(data), CO2Calc(data)
	require.Equal(t, int64(23), oxygen)
	require.Equal(t, int64(10), co2)
}

func Test_LifeSupport_2(t *testing.T) {
	f, err := os.Open("./input.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var data = make([][]rune, 0)
	bf := bufio.NewReader(f)
	for {
		s, err := bf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatalf("unexpected error: %v", err)
		}

		data = append(data, []rune(strings.TrimSpace(s)))
	}

	oxygen, co2 := OxygenCalc(data), CO2Calc(data)
	t.Logf("Oxygen Rate: %d", oxygen)
	t.Logf("CO2 Rate: %d", co2)
	t.Logf("Multiply: %d", oxygen*co2)
}
