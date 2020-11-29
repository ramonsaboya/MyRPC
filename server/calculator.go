package server

import (
	"math"

	"github.com/ramonsaboya/myrpc/commons"
)

type Calculator struct {
	id        string
	available bool
}

func (Calculator) EquationRoots(a, b, c int) commons.EquationRoots {
	delta := delta(a, b, c)

	roots := make([]int, 0)
	if delta == 0 {
		roots = append(roots, (b*-1)/(2*a))
	} else if delta > 0 {
		roots = append(roots, (int(math.Sqrt(delta))-b)/(2*a))
		roots = append(roots, ((-1*int(math.Sqrt(delta)))-b)/(2*a))
	}

	return commons.EquationRoots{Roots: roots}
}

func delta(a, b, c int) float64 {
	return float64((b * b) - (4 * a * c))
}
