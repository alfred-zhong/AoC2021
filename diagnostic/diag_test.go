package diagnostic

import (
	"bufio"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Reporter_1(t *testing.T) {
	data := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	var r = &Reporter{}
	for _, s := range data {
		r.Consume(s)
	}
	require.Equal(t, int64(22), r.GammaRate())
	require.Equal(t, int64(9), r.EpsilonRate())
}

func Test_Reporter_2(t *testing.T) {
	f, err := os.Open("./input.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var r = &Reporter{}
	bf := bufio.NewReader(f)
	for {
		s, err := bf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatalf("unexpected error: %v", err)
		}

		r.Consume(s)
	}

	t.Logf("Gamma Rate: %d", r.GammaRate())
	t.Logf("Epsilon Rate: %d", r.EpsilonRate())
	t.Logf("Multiply: %d", r.GammaRate()*r.EpsilonRate())
}
