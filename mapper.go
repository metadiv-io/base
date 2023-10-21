package base

import (
	"reflect"

	"github.com/metadiv-io/mapper"
)

type Mapper[T any] struct {
	BeforeMap2Model func(from any) any
	AfterMap2Model  func(from any, to *T) *T
}

func (m *Mapper[T]) Map2Model(from any) *T {
	from = neverBePtr(from)
	if m.BeforeMap2Model != nil {
		from = m.BeforeMap2Model(from)
	}
	to := mapper.Map2Model[T](from)
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
		to[i] = *m.Map2Model(fromVal.Index(i).Interface())
	}
	return to
}

func neverBePtr(v any) any {
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		return reflect.ValueOf(v).Elem().Interface()
	}
	return v
}
