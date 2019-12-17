package validator

import (
	"fmt"
	"testing"
)

func TestGt(t *testing.T) {
	fmt.Println(Gt("3n21.11", 66.33))
}

func TestLt(t *testing.T) {
	fmt.Println("abc" > "def")
	fmt.Println(Lte("321.11", 321.10))
}
