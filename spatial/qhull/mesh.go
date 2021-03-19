package qhull

type Facet []int

// SimplexMesh represents a mesh formed by simplices (triangles in 2D, tetrahedrons in 3D etc.)
type SimplexMesh struct {
	Facets []Facet
	Points []Point
}

// FacetsWithPoint returns the index of all facets containing the passed point
func (sm SimplexMesh) FacetsWithPoint(point int) []int {
	facets := []int{}
	for i, f := range sm.Facets {
		if isInSliceInt(f, point) {
			facets = append(facets, i)
		}
	}
	return facets
}

// NeighbouringFacets returns all neighbours of the current facets. Two facets are neighbours if they have at least
// one point in common
func (sm SimplexMesh) NeighbouringFacets(facet int) []int {
	neighbours := []int{}
	for i, f := range sm.Facets {
		if (i != facet) && numSharedInt(f, sm.Facets[facet]) > 0 {
			neighbours = append(neighbours, i)
		}
	}
	return neighbours
}

// FacetCoordinates returns the coordinates of the points that define the facet
func (sm SimplexMesh) FacetCoordinates(facet int) []Point {
	pts := make([]Point, len(sm.Facets[facet]))
	for i, v := range sm.Facets[facet] {
		pts[i] = v
	}
	return pts
}

// isInSliceInt returns true if the value is in slice
func isInSliceInt(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// numSharedInt counts the number of integers that are present in both arrays
func numSharedInt(slice1 []int, slice2 []int) int {
	noShared := 0
	for _, v := range slice1 {
		if isInSliceInt(slice2, v) {
			noShared++
		}
	}
	return noShared
}
