package types

type Primitive interface {
	Display() string
	// Binary Operators

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

	// - String Operators
	Concatenate(other Primitive) (Primitive, error)

	// - Logical Operators
	And(other Primitive) (Primitive, error)
	Or(other Primitive) (Primitive, error)

	// Casting
	RequireNum() (*PrimitiveNum, error)
	RequireStr() (*PrimitiveStr, error)
	RequireBool() (*PrimitiveBool, error)

	CastNum() (*PrimitiveNum, error)
	CastStr() (*PrimitiveStr, error)
	CastBool() (*PrimitiveBool, error)
}
