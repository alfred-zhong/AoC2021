package dive

import (
	"bufio"
	"io"
	"os"
	"reflect"
	"testing"
)

func Test_parseAction(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    Action
		wantErr bool
	}{
		{"t1", args{"forward 10"}, Action{OpForward, 10}, false},
		{"t2", args{"up 5"}, Action{OpUp, 5}, false},
		{"t3", args{"up"}, Action{OpUp, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseAction(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseAction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Diver(t *testing.T) {
	f, err := os.Open("./input.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	d := &Diver{}
	bf := bufio.NewReader(f)
	for {
		s, err := bf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatalf("unexpected error: %v", err)
		}

		d.Consume(s)
	}

	t.Logf("Horizontal: %d", d.Hori)
	t.Logf("Depth: %d", d.Depth)
	t.Logf("Multiplying: %d", d.Hori*d.Depth)
}
