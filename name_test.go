package faker

import (
	"reflect"
	"testing"
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
	language := RandomLanguage()
	name := Name(language)
	ty := reflect.TypeOf(name).Kind()
	if ty != reflect.String {
		t.Error("Error of return name")
	}
}

func TestFirstName(t *testing.T) {
	language := RandomLanguage()
	name := FirstName(language)
	ty := reflect.TypeOf(name).Kind()
	if ty != reflect.String {
		t.Error("Error of return name")
	}
}

func TestLastName(t *testing.T) {
	language := RandomLanguage()
	name := LastName(language)
	ty := reflect.TypeOf(name).Kind()
	if ty != reflect.String {
		t.Error("Error of return name")
	}
}

func TestSexFirstName(t *testing.T) {
	sex := RandSex()
	language := RandomLanguage()
	name := SexFirstName(language, sex)
	ty := reflect.TypeOf(name).Kind()
	if ty != reflect.String {
		t.Error("Error of return name")
	}
}

func TestSexLastName(t *testing.T) {
	language := RandomLanguage()
	sex := RandSex()
	name := SexLastName(language, sex)
	ty := reflect.TypeOf(name).Kind()
	if ty != reflect.String {
		t.Error("Error of return name")
	}
}
