package venture

import (
	"bufio"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLine_Points(t *testing.T) {
	type fields struct {
		Start Point
		End   Point
	}
	tests := []struct {
		name   string
		fields fields
		want   []Point
	}{
		{
			"t1",
			fields{
				Start: Point{1, 1}, End: Point{1, 3},
			},
			[]Point{
				Point{1, 1},
				Point{1, 2},
				Point{1, 3},
			},
		},
		{
			"t2",
			fields{
				Start: Point{1, 1}, End: Point{2, 2},
			},
			[]Point{
				Point{1, 1},
				Point{2, 2},
			},
		},
		{
			"t3",
			fields{
				Start: Point{9, 7}, End: Point{7, 9},
			},
			[]Point{
				Point{9, 7},
				Point{8, 8},
				Point{7, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Line{
				Start: tt.fields.Start,
				End:   tt.fields.End,
			}
			if got := l.Points(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Line.Points() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Cloud_sample_1(t *testing.T) {
	c, err := parseInput("./sample.txt")
	require.NoError(t, err)

	cloud := &Cloud{}
	for l := range c {
		if l.IsHorizontal() || l.IsVertical() {
			cloud.Consume(l)
		}
	}
	cnt := cloud.OverlapPoints(2)
	require.Equal(t, 5, cnt)
}

func Test_Cloud_input_1(t *testing.T) {
	c, err := parseInput("./input.txt")
	require.NoError(t, err)

	cloud := &Cloud{}
	for l := range c {
		if l.IsHorizontal() || l.IsVertical() {
			cloud.Consume(l)
		}
	}
	cnt := cloud.OverlapPoints(2)
	t.Logf("count: %d", cnt)
}

func Test_Cloud_sample_2(t *testing.T) {
	c, err := parseInput("./sample.txt")
	require.NoError(t, err)

	cloud := &Cloud{}
	for l := range c {
		cloud.Consume(l)
	}
	cnt := cloud.OverlapPoints(2)
	require.Equal(t, 12, cnt)
}

func Test_Cloud_input_2(t *testing.T) {
	c, err := parseInput("./input.txt")
	require.NoError(t, err)

	cloud := &Cloud{}
	for l := range c {
		cloud.Consume(l)
	}
	cnt := cloud.OverlapPoints(2)
	t.Logf("count: %d", cnt)
}

func parseInput(path string) (chan Line, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	c := make(chan Line)
	go func() {
		defer f.Close()
		defer close(c)

		br := bufio.NewReader(f)
		for {
			s, err := br.ReadString('\n')
			if err != nil {
				break
			}

			if l, err := ParseLine(strings.TrimSpace(s)); err == nil {
				c <- l
			}
		}
	}()

	return c, nil
}

func TestParseLine(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantL   Line
		wantErr bool
	}{
		{"t1", args{"0,9 -> 5,9"}, Line{Start: Point{X: 0, Y: 9}, End: Point{X: 5, Y: 9}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotL, err := ParseLine(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotL, tt.wantL) {
				t.Errorf("ParseLine() = %v, want %v", gotL, tt.wantL)
			}
		})
	}
}
