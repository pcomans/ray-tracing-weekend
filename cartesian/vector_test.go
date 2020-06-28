package cartesian

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVec3Constructor(t *testing.T) {
	v := Vec3{1, 2, 3}
	assert.Equal(t, 1.0, v.X)
	assert.Equal(t, 2.0, v.Y)
	assert.Equal(t, 3.0, v.Z)
}
func TestNeg(t *testing.T) {
	v := Vec3{-1, -2, -3}
	vn := Neg(&v)

	assert.Equal(t, 1.0, vn.X)
	assert.Equal(t, 2.0, vn.Y)
	assert.Equal(t, 3.0, vn.Z)
}
func TestAdd(t *testing.T) {
	v := Vec3{1, 2, 3}
	w := Vec3{4, 5, 6}

	s := Add(&v, &w)
	r := Add(&w, &v)

	assert.Equal(t, s, r)
	assert.Equal(t, 5.0, s.X)
	assert.Equal(t, 7.0, s.Y)
	assert.Equal(t, 9.0, s.Z)
}

func TestMul(t *testing.T) {
	v := Vec3{1, 2, 3}

	s := Mul(&v, 2)
	r := Mul(&v, -1)

	assert.Equal(t, r, Neg(&v))
	assert.Equal(t, 2.0, s.X)
	assert.Equal(t, 4.0, s.Y)
	assert.Equal(t, 6.0, s.Z)
}

func TestDiv(t *testing.T) {
	v := Vec3{1, 2, 3}

	s := Div(&v, 2)
	r := Div(&v, 0.5)

	assert.Equal(t, Mul(&v, 2), r)
	assert.Equal(t, 0.5, s.X)
	assert.Equal(t, 1.0, s.Y)
	assert.Equal(t, 1.5, s.Z)
}

func TestLengthSquared(t *testing.T) {
	assert.Equal(t, 14.0, Vec3{1, 2, 3}.LengthSquared())
	assert.Equal(t, 14.0, Vec3{-1, 2, 3}.LengthSquared())
}

func TestLength(t *testing.T) {
	assert.Equal(t, math.Sqrt(3), Vec3{1, 1, 1}.Length())
}

func TestDot(t *testing.T) {
	v := Vec3{1, 2, 3}
	w := Vec3{4, 5, 6}

	s := Dot(&v, &w)
	r := Dot(&w, &v)

	assert.Equal(t, s, r)
	assert.Equal(t, 32.0, s)
}

func TestCross(t *testing.T) {
	v := Vec3{1, 2, 3}
	w := Vec3{4, 5, 6}

	s := Cross(&v, &w)
	r := Cross(&w, &v)

	assert.Equal(t, Vec3{-3.0, 6.0, -3.0}, s)
	assert.Equal(t, Vec3{3.0, -6.0, 3.0}, r)
}
