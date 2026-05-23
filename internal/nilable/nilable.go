package nilable

import (
	"bytes"
	"encoding/json"
)

type Nilable[T any] struct {
	isNil bool
	value T
}

func (n Nilable[T]) MarshalJSON() ([]byte, error) {
	if n.isNil {
		return []byte("null"), nil
	}
	return json.Marshal(n.value)
}

func (n *Nilable[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(bytes.TrimSpace(data), []byte("null")) {
		var zero T
		n.value = zero
		n.isNil = true
		return nil
	}
	n.isNil = false
	return json.Unmarshal(data, &n.value)
}

func (n Nilable[T]) IsNil() bool {
	return n.isNil
}

func (n Nilable[T]) Value() (T, bool) {
	return n.value, !n.isNil
}

func Nil[T any]() Nilable[T] {
	return Nilable[T]{isNil: true}
}

func Value[T any](value T) Nilable[T] {
	return Nilable[T]{isNil: false, value: value}
}
