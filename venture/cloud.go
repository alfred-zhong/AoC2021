package venture

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

type Point struct {
	X, Y int
}

type Cloud struct {
	once sync.Once
	m    map[Point]int
}

func (c *Cloud) Consume(l Line) {
	c.once.Do(func() {
		c.m = make(map[Point]int, 0)
	})

	ps := l.Points()
	for _, p := range ps {
		c.m[p] += 1
	}
}

func (c *Cloud) OverlapPoints(leastCnt int) (count int) {
	c.once.Do(func() {
		c.m = make(map[Point]int, 0)
	})

	for _, cnt := range c.m {
		if cnt >= leastCnt {
			count++
		}
	}
	return
}

type Line struct {
	Start, End Point
}

func (l Line) IsHorizontal() bool {
	return l.Start.Y == l.End.Y
}

func (l Line) IsVertical() bool {
	return l.Start.X == l.End.X
}

func (l Line) IsDiagonal() bool {
	return abs(l.End.X-l.Start.X) == abs(l.End.Y-l.Start.Y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (l Line) Points() []Point {
	cnt := max(abs(l.End.Y-l.Start.Y+1), abs(l.End.X-l.Start.X+1))
	ps := make([]Point, 0, cnt)
	for x, y := l.Start.X, l.Start.Y; ; {
		ps = append(ps, Point{X: x, Y: y})

		if l.End.X > l.Start.X {
			x++
			if x > l.End.X {
				break
			}
		} else if l.End.X < l.Start.X {
			x--
			if x < l.End.X {
				break
			}
		}
		if l.End.Y > l.Start.Y {
			y++
			if y > l.End.Y {
				break
			}
		} else if l.End.Y < l.Start.Y {
			y--
			if y < l.End.Y {
				break
			}
		}
	}
	return ps
}

func ParseLine(s string) (l Line, err error) {
	n, err := fmt.Sscanf(s, "%d,%d -> %d,%d", &l.Start.X, &l.Start.Y, &l.End.X, &l.End.Y)
	if err != nil {
		return l, err
	}
	if n != 4 {
		return l, errors.Errorf("count[%d] scanned not 4", n)
	}
	return l, nil
}
