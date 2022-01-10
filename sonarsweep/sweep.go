package sonarsweep

type Sweeper struct {
	last int

	Increased int
	Decreased int
}

func (s *Sweeper) Consume(d int) {
	// first line
	if s.last == 0 {
		s.last = d
		return
	}

	if d > s.last {
		s.Increased++
	} else if d < s.last {
		s.Decreased++
	}
	s.last = d
}

/////////////////////

type SumSweeper struct {
	lastThree []int
	Increased int
}

func (s *SumSweeper) Consume(d int) {
	if s.lastThree == nil {
		s.lastThree = make([]int, 0, 3)
	}

	if len(s.lastThree) < 3 {
		s.lastThree = append(s.lastThree, d)
	} else {
		if d > s.lastThree[0] {
			s.Increased++
		}

		s.lastThree = append(s.lastThree[1:], d)
	}
}
