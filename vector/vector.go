package vector

// Vec3 is a 3D vector
type Vec3 struct {
	x float64
	y float64
	z float64
}

// Neg negates the vector
func (v Vec3) Neg() Vec3 {
	return Vec3{-v.x, -v.y, -v.z}
}

// Add adds two vectors
func (v Vec3) Add(w Vec3) Vec3 {
	return Vec3{-v.x, -v.y, -v.z}
}
