package validator

import (
	"fmt"
	"reflect"
	"testing"
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
