package errors

type InvalidEnumValueError struct {
	TypeName string
	Value    string
}

func (e InvalidEnumValueError) Error() string {
	return "invalid value for " + e.TypeName + ": " + e.Value
}

type UnknownUnionTypeError struct {
	Union string
	Type  string
}

func (e UnknownUnionTypeError) Error() string {
	return "unknown " + e.Union + " type: " + e.Type
}

type UnknownUnionShapeError struct {
	Union string
}

func (e UnknownUnionShapeError) Error() string {
	return "unknown " + e.Union + " shape"
}
