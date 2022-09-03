package cart

// Add adds the other vector to the vector, and returns the pointer to the updated vector
func (v *Vector[T]) Add(other *Vector[T]) *Vector[T] {
	v.X += other.X
	v.Y += other.Y
	return v
}

// AddScalar adds the scalar to the x and the y of the vector, and returns the pointer to the updated vector
func (v *Vector[T]) AddScalar(s T) *Vector[T] {
	v.X += s
	v.Y += s
	return v
}

// Sub subtracts the other vector from the vector, and returns the pointer to the updated vector
func (v *Vector[T]) Sub(other *Vector[T]) *Vector[T] {
	v.X -= other.X
	v.Y -= other.Y
	return v
}

// SubScalar subtracts the scalar from the x and the y of the vector, and returns the pointer to the updated vector
func (v *Vector[T]) SubScalar(s T) *Vector[T] {
	v.X -= s
	v.Y -= s
	return v
}

// Mul multiplies the x and the y of the vector with the x and the y of the other vector,
// and returns the pointer to the updated vector
func (v *Vector[T]) Mul(other *Vector[T]) *Vector[T] {
	v.X *= other.X
	v.Y *= other.Y
	return v
}

// MulScalar multiplies the x and the y of the vector with the scalar, and returns the pointer to the updated vector
func (v *Vector[T]) MulScalar(s T) *Vector[T] {
	v.X *= s
	v.Y *= s
	return v
}

// Div divides the x and the y of the vector by the x and the y of the other vector,
// and returns the pointer to the updated vector
func (v *Vector[T]) Div(other *Vector[T]) *Vector[T] {
	v.X /= other.X
	v.Y /= other.Y
	return v
}

// DivScalar divides the x and the y of the vector by the scalar, and returns the pointer to the updated vector
func (v *Vector[T]) DivScalar(s T) *Vector[T] {
	v.X /= s
	v.Y /= s
	return v
}

// Min sets the x and the y of the vector to be the minimum value of themselves and the x and the y of the other vector,
// and returns the pointer to the updated vector
func (v *Vector[T]) Min(other *Vector[T]) *Vector[T] {
	if v.X > other.X {
		v.X = other.X
	}
	if v.Y > other.Y {
		v.Y = other.Y
	}
	return v
}

// MinScalar sets the x and the y of the vector to be the minimum value of themselves and the scalar,
// and returns the pointer to the updated vector
func (v *Vector[T]) MinScalar(s T) *Vector[T] {
	if v.X > s {
		v.X = s
	}
	if v.Y > s {
		v.Y = s
	}
	return v
}

// Max sets the x and the y of the vector to be the maximum value of themselves and the x and the y of the other vector,
// and returns the pointer to the updated vector
func (v *Vector[T]) Max(other *Vector[T]) *Vector[T] {
	if v.X < other.X {
		v.X = other.X
	}
	if v.Y < other.Y {
		v.Y = other.Y
	}
	return v
}

// MaxScalar sets the x and the y of the vector to be the maximum value of themselves and the scalar,
// and returns the pointer to the updated vector
func (v *Vector[T]) MaxScalar(s T) *Vector[T] {
	if v.X < s {
		v.X = s
	}
	if v.Y < s {
		v.Y = s
	}
	return v
}

// Clamp sets the x of the vector to be in the range [min.X, max.X] and the y of the vector
// to be in the range [min.Y, max.Y], and returns the pointer to the updated vector
func (v *Vector[T]) Clamp(min, max *Vector[T]) *Vector[T] {
	return v.Max(min).Min(max)
}

// ClampScalar sets the x and the y of the vector to be in the range [min, max],
// and returns the pointer to the updated vector
func (v *Vector[T]) ClampScalar(min, max T) *Vector[T] {
	return v.MaxScalar(min).MinScalar(max)
}

// Negate negates the x and the y of the vector, and returns the pointer to the updated vector
func (v *Vector[T]) Negate() *Vector[T] {
	v.X, v.Y = -v.X, -v.Y
	return v
}

// Invert swaps the x and the y of the vector, and returns the pointer to the updated vector
func (v *Vector[T]) Invert() *Vector[T] {
	v.X, v.Y = v.Y, v.X
	return v
}

// Func applies the function f to the x and the y of the vector, and returns the pointer to the updated vector
func (v *Vector[T]) Func(f func(T) T) *Vector[T] {
	v.X, v.Y = f(v.X), f(v.Y)
	return v
}

// Equals returns whether the vector is equal to the other vector
func (v *Vector[T]) Equals(other *Vector[T]) bool {
	return (v.X == other.X) && (v.Y == other.Y)
}

// AlmostEquals returns whether the vector is almost equal to the other vector using the given tolerance
func (v *Vector[T]) AlmostEquals(other *Vector[T], tol T) bool {
	return (other.X > v.X-tol) && (other.X < v.X+tol) && (other.Y > v.Y-tol) && (other.Y < v.Y+tol)
}
