package qhull

import (
	"testing"
)

func TestSharePlane(t *testing.T) {
	tol := 1e-6
	for i, test := range []struct{
		points []Point
		want bool
	}{
		{
			points: []Point{},
			want: true,
		},
		{
			points: []Point{{1.0, 2.0}},
			want: true,
		},
		{
			points: []Point{{1.0, 2.0}, {3.0, 4.0}},
			want: true,
		},

		// Three points not on a line
		{
			points: []Point{{1.0, 2.0}, {3.0, 4.0}, {-1.0, 2.0}},
			want: false,
		},
		// Three points on a line
		{
			points: []Point{{1.0, 2.0}, {3.0, 4.0}, {-1.0, 0.0}},
			want: true,
		},

		// Four points in 3D on a line
		{
			points: []Point{{1.0, 2.0, 0.0}, {3.0, 4.0, 0.0}, {-1.0, 0.0, 0.0}, {0.0, 1.0, 0.0}},
			want: true,
		},

		// Four points in 3D on a plane (z=0)
		{
			points: []Point{{1.0, 2.0, 0.0}, {3.0, -1.0, 0.0}, {-1.0, -50.0, 0.0}, {80.0, 1.0, 0.0}},
			want: true,
		},

		// Four points in 3D not on a plane
		{
			points: []Point{{1.0, 2.0, 0.0}, {3.0, -1.0, 0.0}, {-1.0, -50.0, 0.0}, {80.0, 1.0, 1.0}},
			want: false,
		},
	}{
		got := shareHyperPlane(test.points, tol)
		if got != test.want {
			t.Errorf("Test #%d: Wanted %v got %v\n", i, test.want, got)
		}
	}
}