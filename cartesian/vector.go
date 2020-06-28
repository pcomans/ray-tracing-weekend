package cartesian

import "math"

// Vec3 represents a 3D vector
type Vec3 struct{ X, Y, Z float64 }

func NewVec3(x, y, z float64) Vec3 {
	return Vec3{X: x, Y: y, Z: z}
}

// Color represents an RGB color value
type Color = Vec3

func NewColor(x, y, z float64) Color {
	return Color{X: x, Y: y, Z: z}
}

// Point3 represents a 3D point
type Point3 = Vec3

func NewPoint3(x, y, z float64) Point3 {
	return Point3{X: x, Y: y, Z: z}
}

// Neg negates the vector
func Neg(v *Vec3) Vec3 {
	return Vec3{-v.X, -v.Y, -v.Z}
}

// Add adds two vectors
func Add(v *Vec3, w *Vec3) Vec3 {
	return Vec3{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

// Sub subtracts two vectors
func Sub(v *Vec3, w *Vec3) Vec3 {
	return Vec3{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

// Mul multiplies the vector with a scalar
func Mul(v *Vec3, s float64) Vec3 {
	return Vec3{v.X * s, v.Y * s, v.Z * s}
}

// Div multiplies the vector by a scalar
func Div(v *Vec3, s float64) Vec3 {
	return Mul(v, 1/s)
}

// LengthSquared returns the square of the length of the vector
func (v Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Length returns the length of the vector
func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

// Dot computes the dot product of two vectors
func Dot(v *Vec3, w *Vec3) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

// Cross computes the cross product of two vectors
func Cross(v *Vec3, w *Vec3) Vec3 {
	return Vec3{
		v.Y*w.Z - v.Z*w.Y,
		v.Z*w.X - v.X*w.Z,
		v.X*w.Y - v.Y*w.X,
	}
}
