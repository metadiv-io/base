package base

import "reflect"

type BaseMapper[T any] struct {
	BeforeMap2Model func(from any) any
	AfterMap2Model  func(from any, to *T) *T
}

func (m *BaseMapper[T]) Map2Model(from any) *T {
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

func (m *BaseMapper[T]) Map2Models(from []any) []T {
	var to []T
	for _, f := range from {
		to = append(to, *m.Map2Model(f))
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

func Map2Models[T any](from []any) []T {
	var to []T
	for _, f := range from {
		to = append(to, *Map2Model[T](f))
	}
	return to
}
