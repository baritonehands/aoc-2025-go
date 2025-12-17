package utils

import (
	"cmp"
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func (p Point) OrthogonalNeighbors(xMax, yMax int) []Point {
	ret := make([]Point, 0, 4)
	if p.X < xMax {
		ret = append(ret, Point{p.X + 1, p.Y})
	}
	if p.Y < yMax {
		ret = append(ret, Point{p.X, p.Y + 1})
	}
	if p.X > 0 {
		ret = append(ret, Point{p.X - 1, p.Y})
	}
	if p.Y > 0 {
		ret = append(ret, Point{p.X, p.Y - 1})
	}
	return ret
}

func (p Point) AllNeighbors(xMax, yMax int) []Point {
	ret := p.OrthogonalNeighbors(xMax, yMax)
	if p.X < xMax && p.Y < yMax {
		ret = append(ret, Point{p.X + 1, p.Y + 1})
	}
	if p.X > 0 && p.Y < yMax {
		ret = append(ret, Point{p.X - 1, p.Y + 1})
	}
	if p.X > 0 && p.Y > 0 {
		ret = append(ret, Point{p.X - 1, p.Y - 1})
	}
	if p.X < xMax && p.Y > 0 {
		ret = append(ret, Point{p.X + 1, p.Y - 1})
	}
	return ret
}

func PointCompareYX(p1 Point, p2 Point) int {
	y := cmp.Compare(p1.Y, p2.Y)
	if y != 0 {
		return y
	}
	x := cmp.Compare(p1.X, p2.X)
	return x
}

func PointCompareXY(p1 Point, p2 Point) int {
	x := cmp.Compare(p1.X, p2.X)
	if x != 0 {
		return x
	}
	y := cmp.Compare(p1.Y, p2.Y)
	return y
}

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) String() string {
	return fmt.Sprintf("(%d,%d,%d)", p.X, p.Y, p.Z)
}

func (p Point3D) Compare(other Point3D) int {
	return cmp.Or(cmp.Compare(p.X, other.X), cmp.Compare(p.Y, other.Y), cmp.Compare(p.Z, other.Z))
}

func (p Point3D) DistanceTo(other Point3D) float64 {
	return math.Sqrt(
		math.Pow(float64(other.X-p.X), 2) +
			math.Pow(float64(other.Y-p.Y), 2) +
			math.Pow(float64(other.Z-p.Z), 2))
}
