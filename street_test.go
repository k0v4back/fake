package faker

import (
	"testing"
	"reflect"
)

func TestStreet(t *testing.T) {
	street := RandomStreet()
	typeOfStreet := reflect.TypeOf(street).Kind()
	if typeOfStreet != reflect.String {
		t.Error("Error of return random street")
	}
}

func TestRandomStreet(t *testing.T) {
	language := RandomLanguage()
	street := Street(language)
	typeOfStreet := reflect.TypeOf(street).Kind()
	if typeOfStreet != reflect.String {
		t.Error("Error of return random street")
	}
}