package tool

import "reflect"

func DiffSlices(s1, s2 []interface{}) []interface{} {
	var diff []interface{}

	for _, v1 := range s1 {
		found := false
		for _, v2 := range s2 {
			if reflect.DeepEqual(v1, v2) {
				found = true
				break
			}
		}
		if !found {
			diff = append(diff, v1)
		}
	}

	return diff
}
