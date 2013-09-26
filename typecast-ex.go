/*
* 不同内置类型转换的例子
*/

package main

import (
	"fmt"
	"math"
	"os"
)


func IntToUint8(x int) (uint8, error) {
	if 0 <= x && math.MaxUint8 >= x {
		return uint8(x), nil
	}
	return 0, fmt.Errorf("%d 超出了uint8的数值区间.", x)
}


func Float64ToInt(f float64) (int, error) {
	if f >= math.MinInt32 && f <= math.MaxInt32 {
		w, f := math.Modf(f)
		if f >= 0.5 {
			w++
		}
		return int(w), nil
	}

	return 0, fmt.Errorf("%d 超出了int的数值区间.", f)
}

func main() {
	
	var i int = 255
	// var i int = -1
	// var i int = 256
	
	x, err := Int2Uint8(i)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("the type of x is %T\n", x)
	fmt.Printf("the value of x is %v\n", x)

}
