package cartesian

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRayAt(t *testing.T) {
	o1 := Point3{0, 0, 0}
	d1 := Vec3{1, 1, 1}

	r1 := Ray{o1, d1}

	p1 := r1.at(1)

	assert.Equal(t, 1.0, p1.X)
	assert.Equal(t, 1.0, p1.Y)
	assert.Equal(t, 1.0, p1.Z)

	p2 := r1.at(2)

	assert.Equal(t, 2.0, p2.X)
	assert.Equal(t, 2.0, p2.Y)
	assert.Equal(t, 2.0, p2.Z)

	o2 := Point3{1, 2, 3}
	d2 := Vec3{1, 2, 3}

	r2 := Ray{o2, d2}

	p3 := r2.at(1)

	assert.Equal(t, 2.0, p3.X)
	assert.Equal(t, 4.0, p3.Y)
	assert.Equal(t, 6.0, p3.Z)

	p4 := r2.at(-1)

	assert.Equal(t, 0.0, p4.X)
	assert.Equal(t, 0.0, p4.Y)
	assert.Equal(t, 0.0, p4.Z)
}
