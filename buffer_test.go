package fritzer_test

import (
	"scott/fritzer"
	"testing"
)

func SampleSliceEqual(a []fritzer.Sample, b []fritzer.Sample) bool {
	alen := len(a)
	blen := len(b)
	if alen != blen {
		return false
	}
	for i := 0; i < alen; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestBuffer(t *testing.T) {
	b := fritzer.NewBuffer(6)
	window := b.Get()
	expected := make([]fritzer.Sample, 6)
	if !SampleSliceEqual(window, expected) || !b.Head().Equals(0, 0) {
		t.Fatal("not equal")
	}
	b.Write(fritzer.Sample{Left: 1, Right: 1})
	window = b.Get()
	expected[5].Left = 1
	expected[5].Right = 1
	if !SampleSliceEqual(window, expected) || !b.Head().Equals(0, 0) {
		t.Fatal("not equal")
	}
	b.Write(fritzer.Sample{Left: 2, Right: 2})
	window = b.Get()
	expected[4].Left = 1
	expected[4].Right = 1
	expected[5].Left = 2
	expected[5].Right = 2
	if !SampleSliceEqual(window, expected) || !b.Head().Equals(0, 0) {
		t.Fatal("not equal")
	}
	b.Write(fritzer.Sample{Left: 3, Right: 3})
	window = b.Get()
	expected[3].Left = 1
	expected[3].Right = 1
	expected[4].Left = 2
	expected[4].Right = 2
	expected[5].Left = 3
	expected[5].Right = 3
	if !SampleSliceEqual(window, expected) || !b.Head().Equals(0, 0) {
		t.Fatal("not equal")
	}
	b.Write(fritzer.Sample{Left: 4, Right: 4})
	window = b.Get()
	expected[2].Left = 1
	expected[2].Right = 1
	expected[3].Left = 2
	expected[3].Right = 2
	expected[4].Left = 3
	expected[4].Right = 3
	expected[5].Left = 4
	expected[5].Right = 4
	if !SampleSliceEqual(window, expected) || !b.Head().Equals(0, 0) {
		t.Fatal("not equal")
	}
	b.Write(fritzer.Sample{Left: 5, Right: 5})
	window = b.Get()
	expected[1].Left = 1
	expected[1].Right = 1
	expected[2].Left = 2
	expected[2].Right = 2
	expected[3].Left = 3
	expected[3].Right = 3
	expected[4].Left = 4
	expected[4].Right = 4
	expected[5].Left = 5
	expected[5].Right = 5
	if !SampleSliceEqual(window, expected) || !b.Head().Equals(0, 0) {
		t.Fatal("not equal")
	}
	b.Write(fritzer.Sample{Left: 6, Right: 6})
	window = b.Get()
	expected[0].Left = 1
	expected[0].Right = 1
	expected[1].Left = 2
	expected[1].Right = 2
	expected[2].Left = 3
	expected[2].Right = 3
	expected[3].Left = 4
	expected[3].Right = 4
	expected[4].Left = 5
	expected[4].Right = 5
	expected[5].Left = 6
	expected[5].Right = 6
	if !SampleSliceEqual(window, expected) || !b.Head().Equals(1, 1) {
		t.Fatal("not equal")
	}
	b.Write(fritzer.Sample{Left: 7, Right: 7})
	window = b.Get()
	expected[0].Left = 2
	expected[0].Right = 2
	expected[1].Left = 3
	expected[1].Right = 3
	expected[2].Left = 4
	expected[2].Right = 4
	expected[3].Left = 5
	expected[3].Right = 5
	expected[4].Left = 6
	expected[4].Right = 6
	expected[5].Left = 7
	expected[5].Right = 7
	if !SampleSliceEqual(window, expected) || !b.Head().Equals(2, 2) {
		t.Fatal("not equal")
	}
}
