package reflections

import (
	"fmt"
	"reflect"
)

// GetBool uses reflections to get a bool typed field from
// a struct or a pointer to a struct and returns its value or
// and error if the given obj was invalid
func GetBool(obj interface{}, name string) (bool, error) {
	field, err := getField(obj, name)
	if err != nil {
		return false, err
	}
	if field.Type() != reflect.TypeOf(true) {
		return false, fmt.Errorf("%s is not a bool typed field", name)
	}
	return field.Interface().(bool), nil
}

// GetBool uses reflections to get a bool typed field from
// a struct or a pointer to a struct and returns its value or
// and error if the given obj was invalid
func GetString(obj interface{}, name string) (string, error) {
	field, err := getField(obj, name)
	if err != nil {
		return "", err
	}
	if field.Type() != reflect.TypeOf("") {
		return "", fmt.Errorf("%s is not a string typed field but %s", name, field.Type())
	}
	return field.Interface().(string), nil
}

func getField(obj interface{}, name string) (reflect.Value, error) {
	t := reflect.TypeOf(obj)
	kind := t.Kind()
	if kind != reflect.Struct && !(kind == reflect.Ptr && t.Elem().Kind() == reflect.Struct) {
		return reflect.Value{}, fmt.Errorf("only structs and pointer to structs are supported but got %s", kind)
	}
	v := reflect.ValueOf(obj)
	if kind == reflect.Ptr {
		v = v.Elem()
	}
	field := v.FieldByName(name)
	if !field.IsValid() {
		return reflect.Value{}, fmt.Errorf("%s is not a valid field name", name)
	}

	return field, nil
}
