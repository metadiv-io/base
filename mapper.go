package base

import "reflect"

type Mapper[T any] struct {
	BeforeMap2Model func(from any) any
	AfterMap2Model  func(from any, to *T) *T
}

func (m *Mapper[T]) Map2Model(from any) *T {
	from = neverBePtr(from)
	if m.BeforeMap2Model != nil {
		from = m.BeforeMap2Model(from)
	}
	to := Map2Model[T](from)
	if m.AfterMap2Model != nil {
		to = m.AfterMap2Model(from, to)
	}
	return to
}

func (m *Mapper[T]) Map2Models(from any) []T {
	from = neverBePtr(from)
	fromVal := reflect.ValueOf(from)
	if fromVal.Kind() != reflect.Slice {
		panic("from must be a slice")
	}

	to := make([]T, fromVal.Len())
	for i := 0; i < fromVal.Len(); i++ {
		to = append(to, *m.Map2Model(fromVal.Index(i).Interface()))
	}
	return to
}

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
		to = append(to, *Map2Model[T](fromVal.Index(i).Interface()))
	}
	return to
}
