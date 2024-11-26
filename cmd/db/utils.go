package db

import (
	"fmt"
	"reflect"
	"strings"
)

func structToMap(input any) (map[string]any, error) {
	val := reflect.ValueOf(input)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input must be a struct, got %T", input)
	}

	typ := val.Type()
	result := make(map[string]any)

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)

		if field.PkgPath != "" {
			continue
		}

		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		fieldName := strings.Split(jsonTag, ",")[0]

		result[fieldName] = val.Field(i).Interface()
	}

	return result, nil
}

func sliceToStruct(data []any, output any) error {
	v := reflect.ValueOf(output)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("output must be a pointer to a struct")
	}

	structValue := v.Elem()
	structType := structValue.Type()

	if len(data) != structValue.NumField() {
		return fmt.Errorf("data length does not match the number of struct fields")
	}

	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		if !field.CanSet() {
			return fmt.Errorf("field %s cannot be set", structType.Field(i).Name)
		}

		fieldValue := reflect.ValueOf(data[i])
		if fieldValue.Type().ConvertibleTo(field.Type()) {
			field.Set(fieldValue.Convert(field.Type()))
		} else {
			return fmt.Errorf("cannot assign value of type %s to field %s of type %s",
				fieldValue.Type(), structType.Field(i).Name, field.Type())
		}
	}

	return nil
}
