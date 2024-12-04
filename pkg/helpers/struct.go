package helpers

import (
	"encoding/json"
	"log"
	"reflect"
)

type StructHelper struct{}

func (*StructHelper) StructToUnNilMap(s interface{}) map[string]interface{} {
	jsonData, err := json.Marshal(s)
	if err != nil {
		log.Printf("Error marshalling %v", err)
		return nil
	}

	nullableMap := make(map[string]interface{})
	if err := json.Unmarshal(jsonData, &nullableMap); err != nil {
		log.Printf("Error unmarshal %v", err)
		return nil
	}
	res := make(map[string]interface{})
	for key, val := range nullableMap {
		if isNilOrEmpty(val) {
			continue
		}
		res[key] = val
	}
	return res
}

func isNilOrEmpty(value interface{}) bool {
	if value == nil {
		return true
	}

	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.String:
		return val.String() == ""
	case reflect.Ptr, reflect.Interface:
		return val.IsNil()
	default:
		return false
	}
}
