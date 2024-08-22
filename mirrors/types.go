package mirrors

import "reflect"

func T[T any]() reflect.Type {
    return TypeOf[T]()
}

func TypeOf[T any]() reflect.Type {
	var t T
	return reflect.TypeOf(t)
}

func Empty[T any]() T {
	var t T
	return t
}
