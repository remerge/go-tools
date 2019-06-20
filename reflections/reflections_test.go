package reflections

import "testing"

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
}

func TestGetStringField(t *testing.T) {
	o := Test{
		A: "hallo",
	}
	for _, check := range []r{
		r{o, "A", "hallo", false},
		r{&o, "A", "hallo", false},
		r{o, "B", "", true},
		r{o, "C", "", true},
		r{"asdf", "C", "", true},
	} {
		v, err := GetString(check.object, check.accessor)
		if err == nil {
			if v != check.expected.(string) {
				t.Errorf("value mismatch: expected:%s got:%s", check.expected, v)
				continue
			}
			continue
		}
		if !check.shouldError {
			t.Errorf("error where no error was expected: %v", err)
		}
	}
}
func TestGetStringBool(t *testing.T) {
	o := Test{
		B: true,
	}
	for _, check := range []r{
		r{o, "B", true, false},
		r{&o, "B", true, false},
		r{&o, "A", "hallo", true},
		r{o, "A", "hallo", true},
		r{o, "C", "", true},
		r{"asdf", "C", "", true},
	} {
		v, err := GetBool(check.object, check.accessor)
		if err == nil {
			if v != check.expected.(bool) {
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
