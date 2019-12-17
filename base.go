package validator

import (
	"errors"
	"reflect"
)

//不为空
func NotNull(v interface{}) (b bool, err error) {
	vt, vv, vk := getValAttr(v)

	//判断空指针
	if vt.Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil() {
		return false, nil
	}

	switch vk {
	//数字类型
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
		reflect.Complex64,
		reflect.Complex128, reflect.String:
		if vv.IsZero() {
			return false, nil
		}
		//切片、map类型
	case reflect.Slice, reflect.Array, reflect.Map:
		if vv.Len() == 0 {
			return false, nil
		}
	default:
		return false, errors.New("不支持该类型的空值验证")
	}
	return true, nil
}
