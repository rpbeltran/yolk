package types

type Primitive interface {
	Display() string
	// Operators
	Add(other Primitive) (Primitive, error)
	// Casting
	RequireNum() (*PrimitiveNum, error)
	RequireStr() (*PrimitiveStr, error)
	CastNum() (*PrimitiveNum, error)
	CastStr() (*PrimitiveStr, error)
}
