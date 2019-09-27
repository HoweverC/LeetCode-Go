package main
/*给出一个 32 位的有符号整数，将这个整数中每位上的数字进行反转。*/
/*解法一：循环模运算*/

import (
	"fmt"
	"math"
)
func main()  {

	fmt.Println(reverse(-123))
}

func reverse(x int) int {
	result := 0
	for x != 0 {
		temp := x % 10
		result = result *10 + temp
		if result > math.MaxInt32 || result < math.MinInt32 {
			return 0
		}
		x /= 10
	}
	return result
}