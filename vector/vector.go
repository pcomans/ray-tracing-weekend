package vector

import "math"

// Vec3 represents a 3D vector
type Vec3 struct {
	x float64
	y float64
	z float64
}

// Color represents a 3D point
type Color Vec3

// Point3 represents an RGB color value
type Point3 Vec3

// Neg negates the vector
func Neg(v Vec3) Vec3 {
	return Vec3{-v.x, -v.y, -v.z}
}

// Add adds two vectors
func Add(v Vec3, w Vec3) Vec3 {
	return Vec3{v.x + w.x, v.y + w.y, v.z + w.z}
}

// Mul multiplies the vector with a scalar
func Mul(v Vec3, s float64) Vec3 {
	return Vec3{v.x * s, v.y * s, v.z * s}
}

// Div multiplies the vector by a scalar
func Div(v Vec3, s float64) Vec3 {
	return Mul(v, 1/s)
}

// LengthSquared returns the square of the length of the vector
func (v Vec3) LengthSquared() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

// Length returns the length of the vector
func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}
