package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
	"time"
)

type StructTmp struct {
	TestInt      int    `validate:"NOTNULL;GT=32"`
	TestIntPtr   *int   `validate:"EQ=324"`
	TestIntSlice []*int `validate:"LEN=GT:2"`
	TestNotSet   int
	TestString   *string `validate:"LIKE=abc"`
}

func TestStruct(t *testing.T) {
	var structTest = new(StructTmp)
	//structTest.TestInt=123
	b, rst := Struct(structTest)
	fmt.Println(b, rst)
}

func TestValue(t *testing.T) {
	a := 4 + 3i
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(real(a))
	fmt.Println(imag(a))

	var b complex128 = 3 / 2i
	fmt.Println(b)
	fmt.Println(real(b))
	fmt.Println(imag(b))

	var c = 0i
	fmt.Println(reflect.ValueOf(c).IsZero())

	//var cplx complex64=1231231.23123
	//b,rst:=Value(cplx,"GT(13800138001)")
	//fmt.Println(b,rst)
}

func TestNotNull(t *testing.T) {
	var ss = new(StructTmp)
	//var ss = StructTmp{}
	//ss=nil
	fmt.Println(NotNull(ss))
	fmt.Println(NotNull("ss"))
}

func TestEmail(t *testing.T) {
	b, err := Email("sldfj@lfje")
	fmt.Println(b, err)
	fmt.Println("========================")

	//^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$
	//\w+ 匹配多个 任意字母、数字、下划线、汉字 的字符
	//[-+.] 匹配字符 + - .
	//* 匹配前一个子表达式0交或多次

	bb, err := regexp.MatchString(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`,
		"abc_def_23%4@1a_.wfj")
	fmt.Println(bb, err)

	fmt.Println("========================")

	bb, err = Phone("17512345678", true)
	fmt.Println(bb, err)
}

func TestIdn(t *testing.T) {
	b18Year := 2019
	b18Mon := 12
	b18Day := 12

	nowYear := time.Now().Year()
	nowYear, mon, nowDay := time.Now().Date()

	yearDiff := nowYear - int(b18Year)
	monDiff := int(mon) - int(b18Mon)
	dayDiff := nowDay - int(b18Day)

	if dayDiff < 0 {
		monDiff--
	}

	if monDiff < 0 {
		yearDiff--
	}

	if yearDiff < 0 {
		fmt.Println("还没出生，哪来的身份证号！")
	}
	fmt.Println(yearDiff)
	fmt.Println("===================================123===================")
	b, err := Idn("110000199911223333")
	fmt.Println(b, err)
}

func TestIdn2(t *testing.T) {
	b, err := regexp.MatchString(
		`^(1[1-5]|2[1-3]|3[1-7]|4[1-6]|5[0-4]|6[1-5])\d{4}[1|2]\d{3}(0[1-9]|1[0-2])([0-2][1-9]|3[0|1])\d{3}(\d|x|X)$`,
		"46092129891231263T")
	fmt.Println(b, err)
}
