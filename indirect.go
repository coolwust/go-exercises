package exercises

import (
	"reflect"
	"encoding"
)

func Indirect(rv reflect.Value) (encoding.TextUnmarshaler, reflect.Value) {
	if rv.Kind() != reflect.Ptr && rv.Type().Name() != "" && rv.CanAddr() {
		rv = rv.Addr()
	}
	for {
		// rv is an non-nil interface
		if rv.Kind() == reflect.Interface && !rv.IsNil() {
			if e := rv.Elem(); e.Kind() == reflect.Ptr && !e.IsNil() {
				rv = e
			}
		}

		// rv is not a pointer
		if rv.Kind() != reflect.Ptr {
			break
		}

		// rv is a pointer
		if rv.IsNil() {
			rv.Set(reflect.New(rv.Type()))
		}
		if rv.NumMethod() > 0 {
			if um, ok := rv.Interface().(encoding.TextUnmarshaler); ok {
				return um, reflect.Value{}
			}
		}
		rv = rv.Elem()
	}
	return nil, reflect.Value{}
}
