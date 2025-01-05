package types

type Primitive interface {
	// - Math
	Add(other Primitive) (Primitive, error)
	AddInplace(other Primitive) error
	Subtract(other Primitive) (Primitive, error)
	SubtractInplace(other Primitive) error
	Multiply(other Primitive) (Primitive, error)
	MultiplyInplace(other Primitive) error
	Divide(other Primitive) (Primitive, error)
	DivideInplace(other Primitive) error
	IntDivide(other Primitive) (Primitive, error)
	IntDivideInplace(other Primitive) error
	Modulo(other Primitive) (Primitive, error)
	ModuloInplace(other Primitive) error
	RaisePower(other Primitive) (Primitive, error)
	RaisePowerInplace(other Primitive) error

	Negate() (Primitive, error)

	// - String Operators
	Concatenate(other Primitive) (Primitive, error)
	ConcatenateInPlace(other Primitive) error

	// - Logical Operators
	And(other Primitive) (Primitive, error)
	Or(other Primitive) (Primitive, error)
	AndInplace(other Primitive) error
	OrInplace(other Primitive) error

	Not() (Primitive, error)

	// - Comparisons
	Equal(other Primitive) bool
	LessThan(other Primitive) (bool, error)

	// Casting
	RequireNum() (*PrimitiveNum, error)
	RequireStr() (*PrimitiveStr, error)
	RequireBool() (*PrimitiveBool, error)

	CastImplicitNum() (*PrimitiveNum, error)
	CastExplicitNum() (*PrimitiveNum, error)

	Display() string
	Truthy() bool
}
