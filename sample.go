package fritzer

type Sample struct {
	Left  float64
	Right float64
}

func (s Sample) Equals(left float64, right float64) bool {
	return s.Left == left && s.Right == right
}
