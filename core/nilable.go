package core

import "github.com/WindowsSov8forUs/sonolus-core-go/internal/nilable"

type Nilable[T any] = nilable.Nilable[T]

func Nil[T any]() Nilable[T] {
	return nilable.Nil[T]()
}

func Value[T any](value T) Nilable[T] {
	return nilable.Value(value)
}
