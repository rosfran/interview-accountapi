package test_assertions

import (
	"fmt"
	"testing"
)

func AssertEqual(t *testing.T, param1 interface{}, param2 interface{}) string {

	var message string

	if param1 == param2 {
		return ""
	}

	message = fmt.Sprintf("AssertEqual - %v != %v", param1, param2)

	t.Fatal(message)

	return message
}

func AssertNotEmpty(t *testing.T, a interface{}) string {

	var message string

	if a != nil {
		return ""
	}

	message = fmt.Sprintf("AssertNotEmpty - %v != nil", a)

	t.Fatal(message)

	return message
}

func AssertNotNil(t *testing.T, a interface{}) string {

	var message string

	if a != nil {
		return ""
	}

	message = fmt.Sprintf("AssertNotNil - %v != nil", a)

	t.Fatal(message)

	return message
}

func AssertNil(t *testing.T, a interface{}) string {

	var message string

	if a == nil {
		return ""
	}

	message = fmt.Sprintf("AssertNil - %v != nil", a)

	t.Fatal(message)

	return message
}

func AssertNoError(t *testing.T, a interface{}) string {

	var message string

	if a == nil {
		return ""
	}

	message = fmt.Sprintf("AssertNoError - %v != nil", a)

	t.Fatal(message)

	return message
}

func AssertError(t *testing.T, a interface{}) string {

	var message string

	if a != nil {
		return ""
	}

	message = fmt.Sprintf("AssertError - %v != nil", a)

	t.Fatal(message)

	return message
}
