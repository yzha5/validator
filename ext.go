package validator

import (
	"reflect"
	"regexp"
)

func Value(v interface{}, rules ...string) (b bool, rst *Result) {
	b = true
	rst = &Result{
		Err: nil,
		Msg: make(map[string][]Rule),
	}

	//获取参数类型的字符串
	vts := reflect.TypeOf(v).String()

	for _, rule := range rules {

		//为空
		if rule == string(NOTNULL) {
			b1, e1 := NotNull(v)
			if e1 != nil {
				return b1, &Result{Err: e1}
			}

			//这个，太难解析了，还是看代码吧。
			b = bbb(b, b1)

			rst.Msg[vts] = append(rst.Msg[vts], NOTNULL)
		}

		//邮箱
		if rule == string(EMAIL) {

		}

		//手机号码
		if rule == string(PHONE) {

		}

		//大于
		if bGt, _ := regexp.MatchString(`GT\((.*)\)`, rule); bGt {

		}
	}

	return b, rst
}
