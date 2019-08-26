package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

//func (e ErrNegativeSqrt) Error() string {
////会不断调用e.Error()转换成string
//	return fmt.Sprintf("cannot Sqrt negative number:%v",e)
//}
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

var temps2 []float64

func isOldValue2(value float64) (isOld bool) {
	for _, v := range temps2 {
		if value == v {
			return true
		}
	}
	return false
}

func sqrtTest2(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := float64(1)
	for i := 0; i < 10; i++ {
		temp := (z*z - x) / (2 * z)
		if isOldValue2(temp) {
			return z, nil
		}
		temps2 = append(temps2, temp)
		z -= temp
		fmt.Println(z)
	}
	return z, nil
}

func main() {
	fmt.Println(sqrtTest2(895))
	fmt.Println(sqrtTest2(-895))
}
