package envloader

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func Parse(config interface{}) error {
	val := reflect.ValueOf(config).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		structField := typ.Field(i)

		tag := structField.Tag.Get("env")
		if tag == "" {
			continue
		}

		envValue := os.Getenv(tag)
		if envValue == "" {
			if structField.Tag.Get("required") == "true" {
				return errors.New("missing required environment variable: " + tag)
			}
			continue
		}

		err := setField(field, envValue, structField.Tag.Get("default"))
		if err != nil {
			return err
		}
	}

	return nil
}

func setField(field reflect.Value, value string, defaultValue string) error {
	if value == "" && defaultValue != "" {
		value = defaultValue
	}

	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(intVal)
	case reflect.Bool:
		boolVal, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		field.SetBool(boolVal)
	case reflect.Slice:
		elements := strings.Split(value, ",")
		slice := reflect.MakeSlice(field.Type(), len(elements), len(elements))
		for i, el := range elements {
			err := setField(slice.Index(i), el, "")
			if err != nil {
				return err
			}
		}
		field.Set(slice)
	default:
		return errors.New("unsupported type: " + field.Type().String())
	}

	return nil
}
