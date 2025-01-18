package helpers

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// ExpectEq compares two values and returns an error if they are not equal.
func ExpectEq(actual, expected interface{}, fieldName string) error {
	if !reflect.DeepEqual(actual, expected) {
		return fmt.Errorf("%s mismatch: got '%v', want '%v'", fieldName, actual, expected)
	}
	return nil
}

// CheckStruct compares two structs and optionally includes the "ID" field.
func CheckStruct(obj, expected interface{}, checkID bool) error {
	var errorMessages []string

	// Get the reflection types and values
	objVal := reflect.ValueOf(obj)
	expectedVal := reflect.ValueOf(expected)

	// Ensure both are pointers to structs
	if objVal.Kind() != reflect.Ptr || expectedVal.Kind() != reflect.Ptr {
		return errors.New("both obj and expected must be pointers to structs")
	}
	objVal = objVal.Elem()
	expectedVal = expectedVal.Elem()
	if objVal.Kind() != reflect.Struct || expectedVal.Kind() != reflect.Struct {
		return errors.New("both obj and expected must be structs")
	}

	// Iterate over fields in the struct
	for i := 0; i < objVal.NumField(); i++ {
		field := objVal.Type().Field(i)
		fieldName := field.Name

		// Skip the "ID" field if checkID is false
		if !checkID && fieldName == "ID" {
			continue
		}

		objField := objVal.Field(i)
		expectedField := expectedVal.Field(i)

		// Skip unexported fields
		if !objField.CanInterface() {
			continue
		}

		// Compare fields using ExpectEq
		if err := ExpectEq(objField.Interface(), expectedField.Interface(), fieldName); err != nil {
			errorMessages = append(errorMessages, err.Error())
		}
	}

	if len(errorMessages) > 0 {
		return errors.New(strings.Join(errorMessages, ". "))
	}
	return nil
}
