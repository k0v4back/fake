package faker

import (
	"testing"
	"reflect"
)

func TestRandSex(t *testing.T) {
	sex := RandSex()
	switch sex {
	case FEMALE:
	case MALE:
	default:
		t.Error("Error of generation sex")
	}
}

func TestName(t *testing.T) {
	language := []string{"en", "ru", "de"}
	for _, v := range language {
		name := Name(v)
		ty := reflect.TypeOf(name).Kind()
		if ty != reflect.String {
			t.Error("Error of return name")
		}
	}
}

func TestFirstName(t *testing.T) {
	language := []string{"en", "ru", "de"}
	for _, v := range language {
		name := FirstName(v)
		ty := reflect.TypeOf(name).Kind()
		if ty != reflect.String {
			t.Error("Error of return name")
		}
	}
}

func TestLastName(t *testing.T) {
	language := []string{"en", "ru", "de"}
	for _, v := range language {
		name := LastName(v)
		ty := reflect.TypeOf(name).Kind()
		if ty != reflect.String {
			t.Error("Error of return name")
		}
	}
}

func TestSexFirstName(t *testing.T) {
	language := []string{"en", "ru", "de"}
	sex := RandSex()
	for _, v := range language {
		name := SexFirstName(v, sex)
		ty := reflect.TypeOf(name).Kind()
		if ty != reflect.String {
			t.Error("Error of return name")
		}
	}
}

func TestSexLastName(t *testing.T) {
	language := []string{"en", "ru", "de"}
	sex := RandSex()
	for _, v := range language {
		name := SexLastName(v, sex)
		ty := reflect.TypeOf(name).Kind()
		if ty != reflect.String {
			t.Error("Error of return name")
		}
	}
}