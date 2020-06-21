package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	v := Vec3{1, 2, 3}
	assert.Equal(t, v.x, 1.0)
	assert.Equal(t, v.y, 2.0)
	assert.Equal(t, v.z, 3.0)
}
func TestNeg(t *testing.T) {
	v := Vec3{-1, -2, -3}
	vn := v.Neg()

	if vn.x != 1.0 {
		t.Errorf("x = %f; want 1", vn.x)
	}
	if vn.y != 2.0 {
		t.Errorf("y = %f; want 2", vn.y)
	}
	if vn.z != 3.0 {
		t.Errorf("z = %f; want 3", vn.z)
	}
}
func TestAdd(t *testing.T) {
	v := Vec3{1, 2, 3}
	w := Vec3{4, 5, 6}

	s := v.Add(w)
	r := w.Add(v)

	if s != r {
		t.Errorf("s != t; s = %v; t = %v", s, r)
	}

	if s.x != 5.0 {
		t.Errorf("x = %f; want 5", s.x)
	}
	if s.y != 7.0 {
		t.Errorf("y = %f; want 7", s.y)
	}
	if s.z != 9.0 {
		t.Errorf("z = %f; want 9", s.z)
	}
}
