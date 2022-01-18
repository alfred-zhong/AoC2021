package lanternfish

type Fisher struct {
	Data []int
}

func (f *Fisher) OneDayPassed() {
	n := len(f.Data)
	for i := 0; i < n; i++ {
		f.Data[i] -= 1

		if f.Data[i] < 0 {
			f.Data[i] = 6
			f.Data = append(f.Data, 8)
		}
	}
}

type Fisher2 struct {
	Data map[int]int
}

func (f *Fisher2) OneDayPassed() {
	newData := make(map[int]int, 10)
	for clock, cnt := range f.Data {
		if clock == 0 {
			newData[6] += cnt
			newData[8] += cnt
		} else {
			newData[clock-1] += cnt
		}
	}
	f.Data = newData
}

func (f *Fisher2) Count() int {
	sum := 0
	for _, cnt := range f.Data {
		sum += cnt
	}
	return sum
}

func NewFisher2(arr []int) *Fisher2 {
	d := make(map[int]int, 10)
	for _, a := range arr {
		d[a]++
	}
	return &Fisher2{
		Data: d,
	}
}
