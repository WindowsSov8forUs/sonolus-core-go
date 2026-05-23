package errors

type InvalidEnumValueError struct {
	TypeName string
	Value    string
}

func (e InvalidEnumValueError) Error() string {
	return "invalid value for " + e.TypeName + ": " + e.Value
}
