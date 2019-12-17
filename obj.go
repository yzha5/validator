package validator

import (
	"errors"
	"fmt"
	"reflect"
)

func Struct(o interface{}) (b bool, rst *Result) {
	var t = reflect.TypeOf(o)
	var v = reflect.ValueOf(o)
	//判断是否是struct类型
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
		fmt.Println()
		if t.Kind() != reflect.Struct {
			return false, &Result{Err: errors.New(t.Name() + "不是struct类型")}
		}
	} else if t.Kind() != reflect.Struct {
		return false, &Result{Err: errors.New(t.Name() + "不是struct类型")}
	}

	rst = &Result{
		Err: nil,
		Msg: make(map[string][]Rule),
	}
	b = true

	//遍历结构体成员
	//1.先获取tag
	//2.如果设置了validate，对结构体成员类型进行判断
	for i := 0; i < v.NumField(); i++ {

		//判断结构体成员是否设置了validate
		strs := tagToSlice(t.Field(i).Tag.Get("validate"))
		if len(strs) != 0 {

			//获取结构体成员的类型
			fk := v.Field(i).Kind()
			//如果成员类型是指针
			if fk == reflect.Ptr {
				//将指针类型fk重新赋值为基本数据类型
				fk = v.Field(i).Type().Elem().Kind()

				//如果是结构体
				//if fk == reflect.Struct {
				//	Struct(v.Field(i))
				//}
			}
			//如果是结构体
			//if fk == reflect.Struct {
			//	Struct(v.Field(i))
			//}
			switch fk {
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
				reflect.Complex128:

				for _, str := range strs {
					switch Rule(str) {
					case NOTNULL:
						if v.Field(i).IsZero() {
							b = false
							rst.Msg[t.Field(i).Name] = append(rst.Msg[t.Field(i).Name], NOTNULL)
						}
					}
				}
			case reflect.String:

			case reflect.Ptr, reflect.Struct:
			case reflect.Slice:

			default:
				return false, &Result{Err: errors.New("[" + t.Field(i).Name + ":" + t.Field(i).Type.String() + "] 不支持该类型的数据验证")}
			}
		}
	}
	return b, rst
}
