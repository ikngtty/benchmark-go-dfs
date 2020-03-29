package int

import (
	"fmt"
	"testing"
)

func TestPow(t *testing.T) {
	cases := []struct {
		base     int
		exponent uint
		want     int
	}{
		{5, 0, 1},
		{5, 1, 5},
		{5, 2, 25},
		{5, 3, 125},
	}
	for _, c := range cases {
		caseName := fmt.Sprintf("%d^%d", c.base, c.exponent)
		t.Run(caseName, func(t *testing.T) {
			got := Pow(c.base, c.exponent)
			if got != c.want {
				t.Errorf("want: %d, got: %d", c.want, got)
			}
		})
	}
}
