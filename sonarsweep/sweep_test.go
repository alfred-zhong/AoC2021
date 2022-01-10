package sonarsweep

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test_Sweeper(t *testing.T) {
	f, err := os.Open("./input.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	swp := &Sweeper{}
	bf := bufio.NewReader(f)
	for {
		s, err := bf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatalf("unexpected error: %v", err)
		}

		// t.Log(s)
		num, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			continue
		}

		swp.Consume(num)
	}

	t.Logf("Increased: %d", swp.Increased)
	t.Logf("Decreased: %d", swp.Decreased)
}

func Test_SumSweeper(t *testing.T) {
	f, err := os.Open("./input.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	swp := &SumSweeper{}
	bf := bufio.NewReader(f)
	for {
		s, err := bf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatalf("unexpected error: %v", err)
		}

		// t.Log(s)
		num, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			continue
		}

		swp.Consume(num)
	}

	t.Logf("Increased: %d", swp.Increased)
}
