package cart

import (
	"fmt"

	"github.com/eqgo/types"
)

// Vec is a 2D vector / point on the cartesian coordinate system with an X value and a Y value.
// The vector can be any numerical type.
type Vec[T types.Number] struct {
	X T
	Y T
}

// NewVec returns a new vector with the given type and x and y values.
func NewVec[T types.Number](x, y T) *Vec[T] {
	return &Vec[T]{x, y}
}

// NewVecScalar returns a new vector with both x and y set to the given value.
func NewVecScalar[T types.Number](s T) *Vec[T] {
	return &Vec[T]{s, s}
}

// ConvertVecTo converts the vector of type F to a vector of type T.
func ConvertVecTo[T, F types.Number](v *Vec[F]) *Vec[T] {
	return NewVec(T(v.X), T(v.Y))
}

// Clone returns the pointer to a copy of the vector
func (v *Vec[T]) Clone() *Vec[T] {
	return NewVec(v.X, v.Y)
}

// Copy copies the other vector to the vector, and returns the pointer to the updated vector
func (v *Vec[T]) Copy(other *Vec[T]) *Vec[T] {
	v.X, v.Y = other.X, other.Y
	return v
}

// Set sets the x and the y values of the vector to the given values, and returns the pointer to the updated vector
func (v *Vec[T]) Set(x, y T) *Vec[T] {
	v.X, v.Y = x, y
	return v
}

// SetScalar sets the x and the y values of the vector to the given scalar, and returns the pointer to the updated vector
func (v *Vec[T]) SetScalar(s T) *Vec[T] {
	v.X, v.Y = s, s
	return v
}

// SetDim sets the value of the given dimension, and returns the pointer to the updated vector
func (v *Vec[T]) SetDim(dim Dim, value T) *Vec[T] {
	if dim == X {
		v.X = value
	} else {
		v.Y = value
	}
	return v
}

// Dim gets the value of the given dimension
func (v *Vec[T]) Dim(dim Dim) T {
	if dim == X {
		return v.X
	}
	return v.Y
}

// SetByName sets the value of the given dimension given in string format, and returns the pointer to the updated vector
func (v *Vec[T]) SetByName(dim string, value T) *Vec[T] {
	return v.SetDim(DimByName(dim), value)
}

// GetByName gets the value of the given dimension given in string format
func (v *Vec[T]) GetByName(dim string) T {
	return v.Dim(DimByName(dim))
}

// Zero sets the x and the y of the vector to zero, and returns the pointer to the updated vector
func (v *Vec[T]) Zero() *Vec[T] {
	v.X, v.Y = 0, 0
	return v
}

// IsZero returns whether both the x and the y of the vector are 0
func (v *Vec[T]) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

// String formats the vector in (x, y) format.
func (v *Vec[T]) String() string {
	return fmt.Sprintf("(%v, %v)", v.X, v.Y)
}
