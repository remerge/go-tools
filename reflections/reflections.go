package reflections

import (
	"fmt"
	"reflect"
)

// GetBool uses reflections to get a bool typed field from
// a struct or a pointer to a struct and returns its value or
// and error if the given obj was invalid
func GetBool(obj interface{}, name string) (bool, error) {
	field, err := getFieldKindChecked(obj, name, reflect.Bool)
	if err != nil {
		return false, err
	}
	return field.Interface().(bool), nil
}

// GetString uses reflections to get a string typed field from
// a struct or a pointer to a struct and returns its value or
// and error if the given obj was invalid
func GetString(obj interface{}, name string) (string, error) {
	field, err := getFieldKindChecked(obj, name, reflect.String)
	if err != nil {
		return "", err
	}
	return field.Interface().(string), nil
}

// GetInt uses reflections to get a int typed field from
// a struct or a pointer to a struct and returns its value or
// and error if the given obj was invalid
func GetInt(obj interface{}, name string) (int, error) {
	field, err := getFieldKindChecked(obj, name, reflect.Int)
	if err != nil {
		return 0, err
	}
	return field.Interface().(int), nil
}

// GetInt32 uses reflections to get a int32 typed field from
// a struct or a pointer to a struct and returns its value or
// and error if the given obj was invalid
func GetInt32(obj interface{}, name string) (int32, error) {
	field, err := getFieldKindChecked(obj, name, reflect.Int32)
	if err != nil {
		return 0, err
	}
	return field.Interface().(int32), nil
}

// GetInt64 uses reflections to get a int64 typed field from
// a struct or a pointer to a struct and returns its value or
// and error if the given obj was invalid
func GetInt64(obj interface{}, name string) (int64, error) {
	field, err := getFieldKindChecked(obj, name, reflect.Int)
	if err != nil {
		return 0, err
	}
	return field.Interface().(int64), nil
}

// GetValue uses reflections to get a typed field from
// a struct or a pointer to a struct and returns its value or
// and error if the given obj was invalid
func GetValue(obj interface{}, name string, kind reflect.Kind) (interface{}, error) {
	field, err := getFieldKindChecked(obj, name, kind)
	if err != nil {
		return 0, err
	}
	return field.Interface(), nil
}

func getFieldKindChecked(obj interface{}, name string, kind reflect.Kind) (reflect.Value, error) {
	field, err := getField(obj, name)
	if err != nil {
		return reflect.Value{}, err
	}
	if field.Type().Kind() != kind {
		return reflect.Value{}, fmt.Errorf("%s is not of kind %v typed field but %v", name, kind, field.Type().Kind())
	}
	return field, nil
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
