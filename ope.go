package validator

import (
	"errors"
	"reflect"
	"strconv"
)

//大于
//arg:v (u)int/(u)int8/(u)int16/(u)int32/(u)int64/float32/float64/string
func Gt(v interface{}, val float64) (b bool, err error) {
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
		reflect.Float64:
		if vv.Float() <= val {
			return false, nil
		}
	case reflect.String:
		f, err := strconv.ParseFloat(vv.String(), 64)
		if err != nil {
			return false, errors.New("无效的字符串数值")
		}
		if f <= val {
			return false, nil
		}
	default:
		return false, errors.New("不支持该类型的大于验证")
	}

	return true, nil
}

//大于等于
//arg:v (u)int/(u)int8/(u)int16/(u)int32/(u)int64/float32/float64/string
func Gte(v interface{}, val float64) (b bool, err error) {
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
		reflect.Float64:
		if vv.Float() < val {
			return false, nil
		}
	case reflect.String:
		f, err := strconv.ParseFloat(vv.String(), 64)
		if err != nil {
			return false, errors.New("无效的字符串数值")
		}
		if f < val {
			return false, nil
		}
	default:
		return false, errors.New("不支持该类型的大于等于验证")
	}

	return true, nil
}

//小于
//arg:v (u)int/(u)int8/(u)int16/(u)int32/(u)int64/float32/float64/string
func Lt(v interface{}, val float64) (b bool, err error) {
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
		reflect.Float64:
		if vv.Float() >= val {
			return false, nil
		}
	case reflect.String:
		f, err := strconv.ParseFloat(vv.String(), 64)
		if err != nil {
			return false, errors.New("无效的字符串数值")
		}
		if f >= val {
			return false, nil
		}
	default:
		return false, errors.New("不支持该类型的小于验证")
	}

	return true, nil
}

//小于等于
//arg:v (u)int/(u)int8/(u)int16/(u)int32/(u)int64/float32/float64/string
func Lte(v interface{}, val float64) (b bool, err error) {
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
		reflect.Float64:
		if vv.Float() > val {
			return false, nil
		}
	case reflect.String:
		f, err := strconv.ParseFloat(vv.String(), 64)
		if err != nil {
			return false, errors.New("无效的字符串数值")
		}
		if f > val {
			return false, nil
		}
	default:
		return false, errors.New("不支持该类型的小于等于验证")
	}

	return true, nil
}

//等于
//arg:v (u)int/(u)int8/(u)int16/(u)int32/(u)int64/float32/float64/string
/*
func Eq(v,val interface{}) (b bool, err error) {
	vt, vv, vk := getValAttr(v)

	//判断空指针
	if vt.Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil() {
		return false, nil
	}

	//获取val的值
	vt2, vv2, vk2 := getValAttr(val)


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
		reflect.Float64:
		if vv.Float() > val {
			return false, nil
		}
	case reflect.String:
		f, err := strconv.ParseFloat(vv.String(), 64)
		if err != nil {
			return false, errors.New("无效的字符串数值")
		}
		if f > val {
			return false, nil
		}
	default:
		return false, errors.New("不支持该类型的大于等于验证")
	}

	return b,err
}

*/
