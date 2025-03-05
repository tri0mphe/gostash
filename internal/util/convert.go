package util

import "fmt"

func ConvertMapKeys(m map[interface{}]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range m {
		strKey := fmt.Sprintf("%v", k)
		res[strKey] = v
	}
	return res
}
