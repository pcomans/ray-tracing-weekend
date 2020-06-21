package vector

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	v := Vec3{1, 2, 3}
	assert.Equal(t, 1.0, v.x)
	assert.Equal(t, 2.0, v.y)
	assert.Equal(t, 3.0, v.z)
}
func TestNeg(t *testing.T) {
	v := Vec3{-1, -2, -3}
	vn := Neg(v)

	assert.Equal(t, 1.0, vn.x)
	assert.Equal(t, 2.0, vn.y)
	assert.Equal(t, 3.0, vn.z)
}
func TestAdd(t *testing.T) {
	v := Vec3{1, 2, 3}
	w := Vec3{4, 5, 6}

	s := Add(v, w)
	r := Add(w, v)

	assert.Equal(t, s, r)
	assert.Equal(t, 5.0, s.x)
	assert.Equal(t, 7.0, s.y)
	assert.Equal(t, 9.0, s.z)
}

func TestMul(t *testing.T) {
	v := Vec3{1, 2, 3}

	s := Mul(v, 2)
	r := Mul(v, -1)

	assert.Equal(t, r, Neg(v))
	assert.Equal(t, 2.0, s.x)
	assert.Equal(t, 4.0, s.y)
	assert.Equal(t, 6.0, s.z)
}

func TestDiv(t *testing.T) {
	v := Vec3{1, 2, 3}

	s := Div(v, 2)
	r := Div(v, 0.5)

	assert.Equal(t, Mul(v, 2), r)
	assert.Equal(t, 0.5, s.x)
	assert.Equal(t, 1.0, s.y)
	assert.Equal(t, 1.5, s.z)
}

func TestLengthSquared(t *testing.T) {
	assert.Equal(t, 14.0, Vec3{1, 2, 3}.LengthSquared())
	assert.Equal(t, 14.0, Vec3{-1, 2, 3}.LengthSquared())
}

func TestLength(t *testing.T) {
	assert.Equal(t, math.Sqrt(3), Vec3{1, 1, 1}.Length())
}
