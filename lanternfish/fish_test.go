package lanternfish

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Fisher_sample_1(t *testing.T) {
	initial := []int{3, 4, 3, 1, 2}
	f := &Fisher{
		Data: initial,
	}
	for i := 0; i < 80; i++ {
		f.OneDayPassed()
	}
	require.Len(t, f.Data, 5934)
}

var input = []int{4, 5, 3, 2, 3, 3, 2, 4, 2, 1, 2, 4, 5, 2, 2, 2, 4, 1, 1, 1, 5, 1, 1, 2, 5, 2, 1, 1, 4, 4, 5, 5, 1, 2, 1, 1, 5, 3, 5, 2, 4, 3, 2, 4, 5, 3, 2, 1, 4, 1, 3, 1, 2, 4, 1, 1, 4, 1, 4, 2, 5, 1, 4, 3, 5, 2, 4, 5, 4, 2, 2, 5, 1, 1, 2, 4, 1, 4, 4, 1, 1, 3, 1, 2, 3, 2, 5, 5, 1, 1, 5, 2, 4, 2, 2, 4, 1, 1, 1, 4, 2, 2, 3, 1, 2, 4, 5, 4, 5, 4, 2, 3, 1, 4, 1, 3, 1, 2, 3, 3, 2, 4, 3, 3, 3, 1, 4, 2, 3, 4, 2, 1, 5, 4, 2, 4, 4, 3, 2, 1, 5, 3, 1, 4, 1, 1, 5, 4, 2, 4, 2, 2, 4, 4, 4, 1, 4, 2, 4, 1, 1, 3, 5, 1, 5, 5, 1, 3, 2, 2, 3, 5, 3, 1, 1, 4, 4, 1, 3, 3, 3, 5, 1, 1, 2, 5, 5, 5, 2, 4, 1, 5, 1, 2, 1, 1, 1, 4, 3, 1, 5, 2, 3, 1, 3, 1, 4, 1, 3, 5, 4, 5, 1, 3, 4, 2, 1, 5, 1, 3, 4, 5, 5, 2, 1, 2, 1, 1, 1, 4, 3, 1, 4, 2, 3, 1, 3, 5, 1, 4, 5, 3, 1, 3, 3, 2, 2, 1, 5, 5, 4, 3, 2, 1, 5, 1, 3, 1, 3, 5, 1, 1, 2, 1, 1, 1, 5, 2, 1, 1, 3, 2, 1, 5, 5, 5, 1, 1, 5, 1, 4, 1, 5, 4, 2, 4, 5, 2, 4, 3, 2, 5, 4, 1, 1, 2, 4, 3, 2, 1}

func Test_Fisher_input_1(t *testing.T) {
	f := &Fisher{
		Data: input,
	}
	for i := 0; i < 80; i++ {
		f.OneDayPassed()
	}
	t.Logf("count: %d", len(f.Data))
}

func Test_Fisher2_sample_2(t *testing.T) {
	initial := []int{3, 4, 3, 1, 2}
	f := NewFisher2(initial)
	for i := 0; i < 256; i++ {
		f.OneDayPassed()
	}
	require.Equal(t, 26984457539, f.Count())
}

func Test_Fisher2_input_2(t *testing.T) {
	f := NewFisher2(input)
	for i := 0; i < 256; i++ {
		f.OneDayPassed()
	}
	t.Logf("count: %d", f.Count())
}
