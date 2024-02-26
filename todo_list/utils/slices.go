package utils

import "reflect"

func FilterSliceByCondition[T any](slice []T, condition func(item T) bool) []T {
	var filteredSlice []T
	for _, item := range slice {
		isTrue := condition(item)
		if isTrue {
			filteredSlice = append(filteredSlice, item)
		}
	}
	return filteredSlice
}

func IndexOf[T any](slice []T, item T) int {
	for index := range slice {
		if reflect.DeepEqual(slice[index], item) {
			return index
		}

	}
	return -1
}

func MapSlice[T any](slice []T, cb func(item T) T) []T {
	var tmpSlice []T

	for _, item := range slice {
		tmpSlice = append(tmpSlice, cb(item))
	}

	return tmpSlice
}
