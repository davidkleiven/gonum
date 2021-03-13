package qhull

import (
	"math"
	"gonum.org/v1/gonum/mat"
)

// Point represents a point in a Euclidean k-d space that satisfies the Comparable
// interface.
type Point []float64

// shareHyperPlane returns true if all points lies in a hyper plane.
func shareHyperPlane(points []Point, tol float64) bool {
	if len(points) == 0 {
		return true
	}

	dim := len(points[0])
	if len(points) <= dim {
		return true
	}

	// Transfer the points to a matrix and check and determine the unknown coefficients (a_{k})
	// describing the plane a_{1}*x_{1} + a_{2}*x_{2] + ... + a_{n}*x_{n} = d, d is some unknown
	// constant. To find all unknown a_{k} and d, we form a matrix A with dimensions k x (n+1)
	// where the first n columns corresponds to the coordinates x_{1}...x_{n} and last is set
	// to -1. k is the number of points. Furthermore, we have a coefficient vector q of length n+1,
	// where the first n elements are the a_{k} coefficients and the last corresponds to the d
	// coefficent. We then solve the system Aq = 0. All points share a hyperplane if
	// inf_norm(q) < tol.
	A := mat.NewDense(len(points), dim, nil)
	for i := range points {
		for j, v := range points[i][:dim-1] {
			A.Set(i, j, v)
		}
		A.Set(i, dim-1, 1.0)
	}

	b := mat.NewVecDense(len(points), nil)
	for i, p := range points {
		b.SetVec(i, p[dim-1])
	}
	coeff := mat.NewVecDense(dim, nil)
	coeff.SolveVec(A, b)

	closestInPlane := mat.NewVecDense(b.Len(), nil)
	closestInPlane.MulVec(A, coeff)
	b.SubVec(b, closestInPlane)
	return mat.Norm(b, math.Inf(1)) < tol
}