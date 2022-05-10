package math_test

import (
	"sandbox/math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	m := math.MathImpl{}
	assert.Equal(t, 5, m.Add(2, 3))
}
