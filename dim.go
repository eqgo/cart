package cart

// Dim represents a dimension, either X or Y
type Dim bool

const (
	// X is the dimension X
	X = false
	// Y is the dimension Y
	Y = true
)

// DimByName returns the dimension gotten by its case insensitive name.
// It panics if the given name does not match a dimension.
func DimByName(dim string) Dim {
	if dim == "x" || dim == "X" {
		return X
	} else if dim == "y" || dim == "Y" {
		return Y
	}
	panic("DimByName: the given dimension is not a dimension: " + dim)
}

func (d Dim) String() string {
	if d == X {
		return "x"
	}
	return "y"
}
