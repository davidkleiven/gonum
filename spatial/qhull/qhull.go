package qhull

import (
	"math"
	"gonum.org/v1/gonum/mat"
)

// Point represents a point in a Euclidean space. Satisfies the mat.Vector interface.
type Point []float64

// At returns the element at row i. It panics if i is out of bounds or if j is not zero.
func (p Point) At(i, j int) float64 {
	if j != 0 {
		panic("j must be zero for Point type")
	}
	return p[i]
}

// AtVec returns the element at row i
func (p Point) AtVec(i int) float64 {
	return p[i]
}

// Len returns the length of the vector
func (p Point) Len() int {
	return len(p)
}

// Dims returns the number of rows and columns in the matrix. Columns is always 1.
func (p Point) Dims() (int, int) {
	return p.Len(), 1
}

// T performs an implicit transpose by returning the receiver inside a Transpose.
func (p Point) T() mat.Matrix {
	return mat.Transpose{p}
}

// centroid calculates the centroid of the set of points
func centroid(points []Point) Point {
	c := make(Point, len(points[0]))
	for _, p := range points {
		for i, v := range p {
			c[i] += v
		}
	}

	for i := range c {
		c[i] /= float64(len(points))
	}
	return c
}

// fitPlane fits a plane to the set of points and return the normal vector.
func fitPlane(points []Point) mat.Vector {
	if len(points) == 0 {
		panic("No points given, can't fit a plane")
	}

	dim := len(points[0])
	
	// By shifting the points by the centroid, we ensure that the origin lies in the plane
	c := centroid(points)

	A := mat.NewDense(dim, len(points), nil)
	for i := range points {
		for j, v := range points[i] {
			A.Set(j, i, v - c[j])
		}
	}

	// Calculate the normal vector to the plane 
	var svd mat.SVD
	svd.Factorize(A, mat.SVDFullU)
	var U mat.Dense
	svd.UTo(&U)

	// The normal vector is given by the 
	return U.ColView(dim-1)
}

// dot calculates calculates the dot product between a point and a vector
func dotPointVector(p Point, v mat.Vector) float64 {
	s := 0.0
	for j, x := range p {
		s += x*v.AtVec(j)
	}
	return s
}

// shareHyperPlane returns true if all points lies in a hyper plane.
func shareHyperPlane(points []Point, tol float64) bool {
	if len(points) == 0 {
		return true
	}

	dim := len(points[0])
	if len(points) <= dim {
		return true
	}

	c := centroid(points)
	normal := fitPlane(points)
	
	centroidDotNormal := mat.Dot(c, normal)
	// If all points are perpendicular to the normal, they are in the same plane
	for _, p := range points {
		// Want to check if (p - c) * normal = 0, where c is the centroid
		if math.Abs(mat.Dot(p, normal) - centroidDotNormal) > tol {
			return false
		}
	}
	return true
}