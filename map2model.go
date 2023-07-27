package base

import "reflect"

func Map2Model[T any](from any) *T {
	from = neverBePtr(from)
	to := reflect.ValueOf(new(T)).Elem()

	if from == nil {
		return nil
	}

	fields := parseField(from)
	for _, f := range fields {
		to = setField(to, f)
	}

	return to.Addr().Interface().(*T)
}

func Map2Models[T any](from any) []T {
	from = neverBePtr(from)
	fromVal := reflect.ValueOf(from)
	if fromVal.Kind() != reflect.Slice {
		panic("from must be a slice")
	}

	to := make([]T, fromVal.Len())
	for i := 0; i < fromVal.Len(); i++ {
		to[i] = *Map2Model[T](fromVal.Index(i).Interface())
	}
	return to
}

func MapModel2Model[T any](from any, to *T) *T {
	from = neverBePtr(from)
	toVal := reflect.ValueOf(to).Elem()

	if from == nil {
		return to
	}

	fields := parseField(from)
	for _, f := range fields {
		toVal = setField(toVal, f)
	}

	return toVal.Addr().Interface().(*T)
}
