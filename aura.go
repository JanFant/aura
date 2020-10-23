package aura

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

const tagName = "aura"

func Marshal(v interface{}) []string {
	var sliceStr []string
	rv := reflect.ValueOf(v)
	t := rv.Type()
	for i := 0; i < t.NumField(); i++ {
		tag := rv.Type().Field(i).Tag.Get(tagName)
		if tag == "" || tag == "-" {
			continue
		}
		sliceStr = append(sliceStr, fmt.Sprintf("%v:%v", tag, rv.Field(i)))
	}
	return sliceStr
}

func UnMarshal(data []string, v interface{}) error {
	if reflect.ValueOf(v).Kind() != reflect.Ptr || reflect.ValueOf(v).IsNil() {
		return errors.New(fmt.Sprintf("struct not pointer or nil %v", reflect.TypeOf(v)))
	}
	rv := reflect.ValueOf(v).Elem()
	t := rv.Type()
	for _, str := range data {
		value := strings.Split(str, ":")
		for i := 0; i < t.NumField(); i++ {
			tag := rv.Type().Field(i).Tag.Get(tagName)
			if strings.Compare(value[0], tag) == 0 {
				if rv.Field(i).CanSet() {
					rv.Field(i).SetString(value[1])
				}
				break
			}
		}
	}

	return nil
}
