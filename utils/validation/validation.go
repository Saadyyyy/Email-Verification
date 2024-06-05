package validation

import (
	"errors"
	"reflect"
)

func CheckEmpty(data ...interface{}) error {
	for _, d := range data {
		v := reflect.ValueOf(d)
		switch v.Kind() {
		case reflect.Ptr:
			if v.IsNil() {
				return errors.New("error : data cannot be empty")
			}
		case reflect.String:
			if d == "" {
				return errors.New("error : data cannot be empty")
			}
		case reflect.Slice:
			if v.Len() == 0 {
				return errors.New("error: slice cannot be empty")
			}
		}
	}
	return nil
}
