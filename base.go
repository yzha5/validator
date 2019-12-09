package validator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Result struct {
	Err error
	Msg map[string][]string
}

type Rule string

const (
	NOTNULL Rule = "NOTNULL"
	EMAIL   Rule = "EMAIL"
	PHONE   Rule = "PHONE"
	GT      Rule = "GT"
	GTE     Rule = "GTE"
	LT      Rule = "LT"
	LTE     Rule = "LTE"
	EQ      Rule = "EQ"
	NEQ     Rule = "NEQ"
	LIKE    Rule = "LIKE"
	BETW    Rule = "BETW"
	NBETW   Rule = "NBETW"
	REG     Rule = "REG"
	LOW     Rule = "LOW"
	CAP     Rule = "CAP"
	LETTER  Rule = "LETTER"
	NUMERIC Rule = "NUMERIC"
	AN      Rule = "AN"
	ANP     Rule = "ANP"
	URL     Rule = "URL"
	LEN     Rule = "LEN"
	MINLEN  Rule = "MINLEN"
	MAXLEN  Rule = "MAXLEN"
	MIN     Rule = "MIN"
	MAX     Rule = "MAX"
	ONEOF   Rule = "ONEOF"
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
		Msg: make(map[string][]string),
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
							rst.Msg[t.Field(i).Name] = append(rst.Msg[t.Field(i).Name], str)
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

func Value(v interface{}, rules ...string) (b bool, rst *Result) {
	//获取参数的类型和值
	to := reflect.TypeOf(v)
	vo := reflect.ValueOf(v)
	fmt.Println("TypeOf", to)
	fmt.Println("ValueOf", vo)

	//获取基本数据类型
	vok := vo.Kind()
	//如果基本数据类型是指针
	if vok == reflect.Ptr {
		//将vok转为非指针的基本数据类型
		vok = vo.Type().Elem().Kind()
	}
	rst = &Result{
		Err: nil,
		Msg: make(map[string][]string),
	}
	b = true

	//开始匹配类型
	switch vok {
	//数字类型可以使用
	//NOTNULL
	//PHONE
	//GT
	//GTE
	//LT
	//LTE
	//EQ
	//NEQ
	//LIKE
	//BETW
	//NBETW
	//REG
	//NUMERIC
	//LEN
	//MINLEN
	//MAXLEN
	//MIN
	//MAX
	//ONEOF
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
		//reflect.Complex64,
		//reflect.Complex128:
		for _, rule := range rules {
			//不是空 0，""，nil都会返回false
			if "NOTNULL" == rule {
				if vo.IsZero() {
					b = false
					rst.Msg[to.Name()] = append(rst.Msg[to.Name()], string(rule))
				}
			}

			//11位数电话号码
			//支持号段：
			//130-139、145、147、149-153、155-159、162、165-167、171-173、175-178、180-189、
			//191、193、195、198-199、1700-1709
			if "PHONE" == rule {
				if vok == reflect.Float32 ||
					vok == reflect.Float64 ||
					vok == reflect.Complex64 ||
					vok == reflect.Complex128 {
					return false, &Result{Err: errors.New("PHONE 不支持float/complex类型")}
				}
				switch {
				case vo.Int() >= 13000000000 && vo.Int() <= 13999999999,
					vo.Int() >= 14500000000 && vo.Int() <= 14599999999,
					vo.Int() >= 14900000000 && vo.Int() <= 15399999999,
					vo.Int() >= 15500000000 && vo.Int() <= 15999999999,
					vo.Int() >= 16200000000 && vo.Int() <= 16299999999,
					vo.Int() >= 16500000000 && vo.Int() <= 16799999999,
					vo.Int() >= 17100000000 && vo.Int() <= 17399999999,
					vo.Int() >= 17500000000 && vo.Int() <= 17899999999,
					vo.Int() >= 18000000000 && vo.Int() <= 18999999999,
					vo.Int() >= 19100000000 && vo.Int() <= 19199999999,
					vo.Int() >= 19300000000 && vo.Int() <= 19399999999,
					vo.Int() >= 19500000000 && vo.Int() <= 19599999999,
					vo.Int() >= 19800000000 && vo.Int() <= 19899999999,
					vo.Int() >= 17000000000 && vo.Int() <= 17099999999:
				default:
					b = false
					rst.Msg[to.Name()] = append(rst.Msg[to.Name()], string(rule))
				}
			}
			//大于某值
			if bGt, _ := regexp.MatchString(`GT\((.*)\)`, rule); bGt {
				//提取规则里的值
				ruleVal := strings.Trim(rule, "GT()")
				//将规则值转为float64
				ruleF64, e := strconv.ParseFloat(ruleVal, 10)
				if e != nil {
					return false, &Result{Err: errors.New("GT() 规则需要数字")}
				}
				//如果参数值小于等于规则设置的值
				if float64(vo.Int()) <= ruleF64 {
					b = false
					rst.Msg[to.Name()] = append(rst.Msg[to.Name()], string(rule))
				}
			}

		}
	}

	return b, rst
}

func numericNotNull() {

}

func tagToSlice(str string) (strs []string) {
	if str == "" {
		return strs
	}
	strs = strings.Split(str, ";")
	return strs
}
