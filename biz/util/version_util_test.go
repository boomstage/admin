package util

import (
	"fmt"
	"testing"
)

func TestVersion(t *testing.T) {
	fmt.Println("LessThanOrEqualV:", LessThanOrEqualV("1.1.3xx", "1.1.3wwww"))
	fmt.Println("LessThanOrEqualV:", LessThanOrEqualV("1.1.2", "1.1.3wwww"))
	fmt.Println("LessThanOrEqualV:", LessThanOrEqualV("1.1.3", "1.1.3"))

	fmt.Println("GreaterThanOrEqualV:", GreaterThanOrEqualV("1.1.3xx", "1.1.3wwww"))
	fmt.Println("GreaterThanOrEqualV:", GreaterThanOrEqualV("1.1.2", "1.1.3wwww"))
	fmt.Println("GreaterThanOrEqualV:", GreaterThanOrEqualV("1.1.3", "1.1.3"))

	fmt.Println(ConstraintsVV("1.1.3", "<1.2, >=1.1.3"))
	vv := SortVV([]string{"1.1", "0.7.1", "1.4-beta", "1.4", "2"})
	for _, v := range vv {
		fmt.Println(v)
	}
	fmt.Println(vv)

	fmt.Println("相等：", CompareV("1.1.3", "1.1.3"))
	fmt.Println("大于：", CompareV("1.2.3", "1.1.3"))
	fmt.Println("小于：", CompareV("1.1.3", "1.1.4"))

}
