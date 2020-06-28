package cartesian

type Ray struct {
	Origin    *Point3
	Direction *Vec3
}

func (r Ray) at(t float64) Point3 {
	d := Mul(r.Direction, t)
	return Add(r.Origin, &d)
}
