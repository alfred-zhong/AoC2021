package dive

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Diver struct {
	Aim   int
	Hori  int
	Depth int
}

type Op string

const (
	OpForward Op = "forward"
	OpDown    Op = "down"
	OpUp      Op = "up"
)

type Action struct {
	Op  Op
	Num int
}

func (d *Diver) Consume(s string) {
	act, err := parseAction(s)
	if err != nil {
		return
	}

	switch act.Op {
	case OpForward:
		d.Hori += act.Num
		d.Depth += act.Num * d.Aim
	case OpDown:
		d.Aim += act.Num
	case OpUp:
		d.Aim -= act.Num
	}
}

func parseAction(s string) (Action, error) {
	var act Action
	n, err := fmt.Sscanf(strings.TrimSpace(s), "%s %d", &act.Op, &act.Num)
	if err != nil {
		return act, err
	}
	if n != 2 {
		return act, errors.Errorf("%s can't be parsed", s)
	}
	return act, nil
}
