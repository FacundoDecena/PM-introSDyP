package pathfinder

import "testing"

var cases = []struct {
	o, d int64
	dim  int
	want string
}{
	{0, 0, 0, "0"},
	{0, 7, 3, "000 -> 001 -> 011 -> 111"},
	{2, 4, 3, "010 -> 000 -> 100"},
	{0, 15, 4, "0000 -> 0001 -> 0011 -> 0111 -> 1111"},
	{7, 27, 6, "000111 -> 000011 -> 001011 -> 011011"},
	{27, 7, 6, "011011 -> 011111 -> 010111 -> 000111"},
	{7, 27, 5, "00111 -> 00011 -> 01011 -> 11011"},
	{27, 7, 5, "11011 -> 11111 -> 10111 -> 00111"},
}

func TestPathFinder(t *testing.T) {

	for _, c := range cases {
		got := PathFinder(c.o, c.d, c.dim)
		if got != c.want {
			t.Errorf("PathFinder(%d,%d,%d) == %q, want %q", c.o, c.d, c.dim, got, c.want)
		}
	}
}

func BenchmarkPathFinder(b *testing.B) {
	for i := 0; i < b.N; i++ {

		for _, c := range cases {
			PathFinder(c.o, c.d, c.dim)
		}

	}
}
