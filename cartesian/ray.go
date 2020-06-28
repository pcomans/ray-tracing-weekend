package cartesian

type Ray struct {
	Origin    Point3
	Direction Vec3
}

func NewRay(origin Point3, direction Vec3) Ray {
	return Ray{
		Origin:    origin,
		Direction: direction,
	}
}

func (r Ray) at(t float64) Point3 {
	d := Mul(r.Direction, t)
	return Add(r.Origin, d)
}
