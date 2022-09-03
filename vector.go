package cart

import (
	"fmt"

	"github.com/eqgo/types"
)

// Vector is a 2D vector / point on the cartesian coordinate system with an X value and a Y value.
// The vector can be any numerical type.
type Vector[T types.Number] struct {
	X T
	Y T
}

// NewVector returns a new vector with the given type and x and y values.
func NewVector[T types.Number](x, y T) *Vector[T] {
	return &Vector[T]{x, y}
}

// NewVectorScalar returns a new vector with both x and y set to the given value.
func NewVectorScalar[T types.Number](s T) *Vector[T] {
	return &Vector[T]{s, s}
}

// ConvertTo converts the vector of type F to a vector of type T.
func ConvertTo[T, F types.Number](v *Vector[F]) *Vector[T] {
	return NewVector(T(v.X), T(v.Y))
}

// Clone returns the pointer to a copy of the vector
func (v *Vector[T]) Clone() *Vector[T] {
	return NewVector(v.X, v.Y)
}

// Copy copies the other vector to the vector, and returns the pointer to the updated vector
func (v *Vector[T]) Copy(other *Vector[T]) *Vector[T] {
	v.X, v.Y = other.X, other.Y
	return v
}

// Set sets the x and the y values of the vector to the given values, and returns the pointer to the updated vector
func (v *Vector[T]) Set(x, y T) *Vector[T] {
	v.X, v.Y = x, y
	return v
}

// SetScalar sets the x and the y values of the vector to the given scalar, and returns the pointer to the updated vector
func (v *Vector[T]) SetScalar(s T) *Vector[T] {
	v.X, v.Y = s, s
	return v
}

// SetDim sets the value of the given dimension, and returns the pointer to the updated vector
func (v *Vector[T]) SetDim(dim Dim, value T) *Vector[T] {
	if dim == X {
		v.X = value
	} else {
		v.Y = value
	}
	return v
}

// Dim gets the value of the given dimension
func (v *Vector[T]) Dim(dim Dim) T {
	if dim == X {
		return v.X
	}
	return v.Y
}

// SetByName sets the value of the given dimension given in string format, and returns the pointer to the updated vector
func (v *Vector[T]) SetByName(dim string, value T) *Vector[T] {
	return v.SetDim(DimByName(dim), value)
}

// GetByName gets the value of the given dimension given in string format
func (v *Vector[T]) GetByName(dim string) T {
	return v.Dim(DimByName(dim))
}

// Zero sets the x and the y of the vector to zero, and returns the pointer to the updated vector
func (v *Vector[T]) Zero() *Vector[T] {
	v.X, v.Y = 0, 0
	return v
}

// IsZero returns whether both the x and the y of the vector are 0
func (v *Vector[T]) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

// String formats the vector in (x, y) format.
func (v *Vector[T]) String() string {
	return fmt.Sprintf("(%v, %v)", v.X, v.Y)
}
