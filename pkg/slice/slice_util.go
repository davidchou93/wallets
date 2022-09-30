package slice

import "errors"

func SliceContainElement[value comparable](slice []value, ele value) bool {
	for _, sliceEle := range slice {
		if sliceEle == ele {
			return true
		}
	}
	return false
}

func DiffTwoSlices[value comparable](newVals, oldVals []value) (added, removed []value) {
	for _, newVal := range newVals {
		found := false
		for _, oldVal := range oldVals {
			if newVal == oldVal {
				found = true
				break
			}
		}

		if !found {
			added = append(added, newVal)
		}
	}

	for _, oldVal := range oldVals {
		found := false
		for _, newVal := range newVals {
			if newVal == oldVal {
				found = true
				break
			}
		}

		if !found {
			removed = append(removed, oldVal)
		}
	}
	return
}

// UniqueSlice turns input slice into a 'set'(same element will only appear once)
func UniqueSlice[v comparable](oldSlice []v) []v {
	keys := make(map[v]bool)
	var list []v
	for _, entry := range oldSlice {
		if _, exist := keys[entry]; !exist {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func Filter[value any](slice []value, matchFunction func(value) bool) []value {
	var result []value
	for _, each := range slice {
		if matchFunction(each) {
			result = append(result, each)
		}
	}
	return result
}

// Find search for the first matched element which matchFunction(ele) = true
func Find[value any](slice []value, matchFunction func(value) bool) (v value, err error) {
	for _, ele := range slice {
		if matchFunction(ele) {
			v = ele
			return
		}
	}
	err = errors.New("not found")
	return
}

// Exist returns whether element matches matchFunction(ele) = true exist
func Exist[value any](slice []value, matchFunction func(value) bool) bool {
	for _, ele := range slice {
		if matchFunction(ele) {
			return true
		}
	}
	return false
}
