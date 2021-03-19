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

		// Four points in a plane (z=0)
		{
			points: []Point{{1.0, 2.0, 0.0}, {3.0, -1.0, 0.0}, {-1.0, -50.0, 0.0}, {80.0, 1.0, 0.0}},
			want: true,
		},

		// Four points in a plane (y=0)
		{
			points: []Point{{1.0, 0.0, 6.0}, {3.0, 0.0, -40.0}, {-1.0, 0.0, 0.0}, {80.0, 0.0, 32.0}},
			want: true,
		},

		// Four points in a plane (x=0)
		{
			points: []Point{{0.0, 0.0, 6.0}, {0.0, 2.0, -40.0}, {0.0, -50.0, 0.0}, {0.0, -6.0, 32.0}},
			want: true,
		},

		// Four points in 3D not on a plane
		{
			points: []Point{{1.0, 2.0, 0.0}, {3.0, -1.0, 0.0}, {-1.0, -50.0, 0.0}, {80.0, 1.0, 1.0}},
			want: false,
		},

		// Four points in 3D not in the plane x + y = 0
		{
			points: []Point{{1.0, -1.0, 0.0}, {3.0, -3.0, 0.0}, {-1.0, 1.0, 0.0}, {-1.0, 1.0, 1.0}},
			want: true,
		},

		// Four points in 3D in the plane x - 2z = 1
		{
			points: []Point{{1.0, -1.0, 0.0}, {3.0, -3.1, 1.0}, {-1.0, 1.0, -1.0}, {-2.0, 1.0, -1.5}},
			want: true,
		},

		// Four points in 3D not in the plane x - 2z = 1 (add 0.01 to the z coordinate of the last point)
		{
			points: []Point{{1.0, -1.0, 0.0}, {3.0, -3.1, 1.0}, {-1.0, 1.0, -1.0}, {-2.0, 1.0, -1.49}},
			want: false,
		},
	}{
		got := shareHyperPlane(test.points, tol)
		if got != test.want {
			t.Errorf("Test #%d: Wanted %v got %v\n", i, test.want, got)
		}
	}
}

func TestInitialHull(t *testing.T) {
	points := []Point{{1.0, -1.0, 0.0}, {3.0, -3.1, 1.0}, {-1.0, 1.0, -1.0}, {-2.0, 1.0, -1.49}}
	facets := initialHull(points, 1e-6)

	want := []Facet{
		{1, 2, 3},
		{0, 2, 3},
		{0, 1, 3},
		{0, 1, 2},
	}

	if len(facets) != len(want) {
		t.Errorf("Wrong number of facets. Got %d want %d\n", len(facets), len(want))
	}
	for i := range facets {
		for j := range facets[i] {
			if facets[i][j] != want[i][j] {
				t.Errorf("Facet #%d: wanted %v got %v\n", i, want[i], facets[i])
			}
		}
	}
}