package main

/*判断一个整数是否是回文数*/
/*解法一:如果整数大于零，循环取模，反转后进行判断，如果小于零，不是回文数*/
import (
	"fmt"
)
func main()  {

	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	temp := x
	var result int
	for temp != 0 {
		result = result *10 + temp % 10
		temp /= 10
	}
	if x < 0 || result != x {
		return false
	}
	return true
}