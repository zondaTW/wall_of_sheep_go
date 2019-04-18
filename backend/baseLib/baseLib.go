package baseLib

import (
	"fmt"
	"reflect"
	"strings"
)

func Dir(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)

	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name,
			strings.TrimPrefix(methType.String(), "func"))
	}
}

func FindIntArray(sourceArray []int, target int) bool {
	for _, source := range sourceArray {
		if source == target {
			return true
		}
	}
	return false
}

func FindStringArray(sourceArray []string, target string) bool {
	for _, source := range sourceArray {
		if source == target {
			return true
		}
	}
	return false
}
