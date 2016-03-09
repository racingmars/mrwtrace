package main

import "math"

type Vec3 struct {
	X, Y, Z float64
}

// Add two vectors.
func (v1 *Vec3) Add(v2 *Vec3) *Vec3 {
	return &Vec3{X: v1.X + v2.X, Y: v1.Y + v2.Y, Z: v1.Z + v2.Z}
}

func (v1 *Vec3) Sub(v2 *Vec3) *Vec3 {
	return &Vec3{X: v1.X - v2.X, Y: v1.Y - v2.Y, Z: v1.Z - v2.Z}
}

func (v1 *Vec3) Mult(v2 *Vec3) *Vec3 {
	return &Vec3{X: v1.X * v2.X, Y: v1.Y * v2.Y, Z: v1.Z * v2.Z}
}

func (v *Vec3) MultParts(m float64) *Vec3 {
	return &Vec3{X: v.X * m, Y: v.Y * m, Z: v.Z * m}
}

func (v1 *Vec3) Div(v2 *Vec3) *Vec3 {
	return &Vec3{X: v1.X / v2.X, Y: v1.Y / v2.Y, Z: v1.Z / v2.Z}
}

func (v *Vec3) DivParts(d float64) *Vec3 {
	return &Vec3{X: v.X / d, Y: v.Y / d, Z: v.Z / d}
}

func (v1 *Vec3) Dot(v2 *Vec3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func (v1 *Vec3) Cross(v2 *Vec3) *Vec3 {
	return &Vec3{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: -(v1.X*v2.Z - v1.Z*v2.X),
		Z: v1.X*v2.Y - v1.Y*v2.X}
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vec3) UnitVector() *Vec3 {
	return v.DivParts(v.Length())
}

// Rays

type Ray struct {
	A, B Vec3
}

func (r *Ray) Origin() *Vec3 {
	return &r.A
}

func (r *Ray) Direction() *Vec3 {
	return &r.B
}

func (r *Ray) PointAtParameter(t float64) *Vec3 {
	return r.A.Add(r.B.MultParts(t))
}
