package validator

import (
	"reflect"
	"strings"
)

func getValAttr(v interface{}) (vt reflect.Type, vv reflect.Value, vk reflect.Kind) {
	//获取参数的类型和值
	vt = reflect.TypeOf(v)
	vv = reflect.ValueOf(v)

	//如果基本数据类型是指针
	if vv.Kind() == reflect.Ptr {
		vk = vv.Type().Elem().Kind()
	} else {
		vk = vv.Kind()
	}

	return
}

func bbb(b, b1 bool) bool {
	if b && b != b1 {
		b = false
	}
	return b
}

func tagToSlice(str string) (strs []string) {
	if str == "" {
		return strs
	}
	strs = strings.Split(str, ";")
	return strs
}
