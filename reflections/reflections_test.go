package reflections

import (
	"reflect"
	"testing"
)

type Test struct {
	A string
	B bool
	C int
}

type r struct {
	object      interface{}
	accessor    string
	expected    interface{}
	shouldError bool
	kind        reflect.Kind
}

func TestGetValue(t *testing.T) {
	o := Test{
		A: "hallo",
		B: true,
		C: 42,
	}
	for _, check := range []r{
		{o, "A", "hallo", false, reflect.String},
		{&o, "A", "hallo", false, reflect.String},
		{o, "B", "", true, reflect.String},
		{o, "C", "", true, reflect.String},
		{"asdf", "C", "", true, reflect.String},

		{o, "B", true, false, reflect.Bool},
		{&o, "B", true, false, reflect.Bool},
		{&o, "A", "hallo", true, reflect.Bool},
		{o, "A", "hallo", true, reflect.Bool},
		{o, "C", "", true, reflect.Bool},
		{"asdf", "C", "", true, reflect.Bool},

		{o, "C", 42, false, reflect.Int},
		{&o, "C", 42, false, reflect.Int},
		{o, "B", "", true, reflect.Int},
		{o, "A", "", true, reflect.Int},
		{"asdf", "C", "", true, reflect.Int},
	} {
		v, err := GetValue(check.object, check.accessor, check.kind)
		if err == nil {
			if !reflect.DeepEqual(v, check.expected) {
				t.Errorf("value mismatch: expected:%v got:%v", check.expected, v)
				continue
			}
			continue
		}
		if !check.shouldError {
			t.Errorf("error where no error was expected: %v", err)
		}
	}
}
