package utils

import "reflect"

// InArray checks is value exists in an Array (Slice)
func InArray(value interface{}, array interface{}) bool {
	return AtArrayPosition(value, array) != -1
}

// AtArrayPosition finds the position(int) of value in an Array(Slice)
func AtArrayPosition(value interface{}, array interface{}) (index int) {
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) == true {
				index = i

				return
			}
		}
	}

	return
}

// InSlice checks if an element in a slice
func InSlice[T comparable](values []T, value T) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}

	return false
}

// IndexOf returns an elements position in a slice
// Return -1 if element not found in slice
func IndexOf[T comparable](values []T, value T) int {
	for i, v := range values {
		if v == value {
			return i
		}
	}

	return -1
}
