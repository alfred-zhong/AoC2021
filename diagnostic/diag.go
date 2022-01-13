package diagnostic

import (
	"strconv"
	"strings"
)

type Reporter struct {
	columns []Column
}

type Column struct {
	ZeroCount int
	OneCount  int
}

func (r *Reporter) Consume(s string) {
	for i, char := range []rune(strings.TrimSpace(s)) {
		if len(r.columns) <= i {
			r.columns = append(r.columns, Column{})
		}

		switch char {
		case '0':
			r.columns[i].ZeroCount++
		case '1':
			r.columns[i].OneCount++
		}
	}
}

func (r *Reporter) GammaRate() int64 {
	var chars []rune
	for _, col := range r.columns {
		if col.ZeroCount >= col.OneCount {
			chars = append(chars, '0')
		} else {
			chars = append(chars, '1')
		}
	}

	i, _ := strconv.ParseInt(string(chars), 2, 64)
	return i
}

func (r *Reporter) EpsilonRate() int64 {
	var chars []rune
	for _, col := range r.columns {
		if col.ZeroCount >= col.OneCount {
			chars = append(chars, '1')
		} else {
			chars = append(chars, '0')
		}
	}

	i, _ := strconv.ParseInt(string(chars), 2, 64)
	return i
}
