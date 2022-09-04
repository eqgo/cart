package cart

import (
	"fmt"

	"github.com/eqgo/types"
)

// Func is a mathematical function on the 2D cartestian coordinate system
// of the form y = f(x) that takes one x input and returns one y output.
type Func[T types.Number] func(x T) T

// NewFunc makes a new function n(x) = f(x)
func NewFunc[T types.Number](f func(x T) T) Func[T] {
	return f
}

// NewFuncConst makes a new function n(x) = c
func NewFuncConst[T types.Number](c T) Func[T] {
	return func(x T) T {
		return c
	}
}

// ConvertFuncTo converts the function of type F to a function of type T.
func ConvertFuncTo[T, F types.Number](f Func[F]) Func[T] {
	return func(x T) T {
		return T(f(F(x)))
	}
}

// Add returns the function n(x) = f(x) + g(x)
func (f Func[T]) Add(g Func[T]) Func[T] {
	return func(x T) T {
		return f(x) + g(x)
	}
}

// AddConst returns the function n(x) = f(x) + c
func (f Func[T]) AddConst(c T) Func[T] {
	return func(x T) T {
		return f(x) + c
	}
}

// Sub returns the function n(x) = f(x) - g(x)
func (f Func[T]) Sub(g Func[T]) Func[T] {
	return func(x T) T {
		return f(x) - g(x)
	}
}

// SubConst returns the function n(x) = f(x) - c
func (f Func[T]) SubConst(c T) Func[T] {
	return func(x T) T {
		return f(x) - c
	}
}

// Mul returns the function n(x) = f(x) * g(x)
func (f Func[T]) Mul(g Func[T]) Func[T] {
	return func(x T) T {
		return f(x) * g(x)
	}
}

// MulConst returns the function n(x) = f(x) * c
func (f Func[T]) MulConst(c T) Func[T] {
	return func(x T) T {
		return f(x) * c
	}
}

// Div returns the function n(x) = f(x) / g(x)
func (f Func[T]) Div(g Func[T]) Func[T] {
	return func(x T) T {
		return f(x) / g(x)
	}
}

// DivConst returns the function n(x) = f(x) / c
func (f Func[T]) DivConst(c T) Func[T] {
	return func(x T) T {
		return f(x) / c
	}
}

// Of returns the function n(x) = f(g(x))
func (f Func[T]) Of(g Func[T]) Func[T] {
	return func(x T) T {
		return f(g(x))
	}
}

// Reflect returns f reflected over the given axis
func (f Func[T]) Reflect(axis Dim) Func[T] {
	if axis == X {
		return f.ReflectX()
	}
	return f.ReflectY()
}

// ReflectX returns the function n(x) = -f(x)
func (f Func[T]) ReflectX() Func[T] {
	return func(x T) T {
		return -f(x)
	}
}

// ReflectY returns the function n(x) = f(-x)
func (f Func[T]) ReflectY() Func[T] {
	return func(x T) T {
		return f(-x)
	}
}

// String formats the function as a string
func (f Func[T]) String() string {
	return fmt.Sprintf("%T", f)
}
