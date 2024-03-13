package main

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

var (
	ErrRequired = errors.New("required fields are not filled")
	ErrMustPrt  = errors.New("object must be a pointer of struct")
)

func CheckRequiredFields(obj interface{}) error {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return ErrMustPrt
	}

	v := reflect.ValueOf(obj).Elem()

	for i := 0; i < v.NumField(); i++ {
		valueField := v.Field(i)
		typeField := v.Type().Field(i)

		requiredTag := typeField.Tag.Get("required")
		jsonTag := typeField.Tag.Get("json")
		jsonTag = strings.TrimRight(jsonTag, ",omitempty")

		if requiredTag == "" || requiredTag == "-" {
			continue
		}

		b, _ := strconv.ParseBool(requiredTag)

		if b {
			switch valueField.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if valueField.Int() == 0 {
					return ErrRequired
				}
			case reflect.Float32, reflect.Float64:
				if valueField.Float() == 0 {
					return ErrRequired
				}
			case reflect.String:
				if len(strings.TrimSpace(valueField.String())) == 0 {
					return ErrRequired
				}
			case reflect.Ptr:
				if valueField.Interface() == nil {
					return ErrRequired
				}
			case reflect.Slice:
				if valueField.Len() == 0 {
					return ErrRequired
				}
			}
		}
	}
	return nil
}
