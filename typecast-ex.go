/*
* 不同内置类型转换的例子
*/

package main

import (
	"fmt"
	"math"
	"os"
)


func Int2Uint8(x int) (uint8, error) {
	if 0 <= x && math.MaxInt8 >= x {
		return uint8(x), nil
	}
	return 0, fmt.Errorf("%d 超出了uint8的数值区间.", x)
}

func main() {
	
	var i int = 127
	// var i int = -1
	// var i int = 128
	
	x, err := Int2Uint8(i)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("the type of x is %T\n", x)
	fmt.Printf("the value of x is %v\n", x)

}
