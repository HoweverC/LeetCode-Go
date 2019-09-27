package main

/*罗马数字转整数*/
/*解法：使用map保存七个罗马数字及其对应的整数，遍历字符串，如果当前字符不是字符串的第一个字符或者对应的整数比下一个字符对应的整数小，减法，如果大，加法*/
/*做减法时，因为不可能是第一个字符，所以前一个字符已经做过加法，所以要减去二倍的前一个字符*/

import (
	"fmt"
)
func main()  {

	fmt.Println(romanToInt("V"))
}


func romanToInt(s string) int {
	RomanMap := map[rune] int {'I':1, 'V':5, 'X':10, 'L':50, 'C':100, 'D':500, 'M':1000}
	var result int
	for index,value :=range s{
		if index == 0 || RomanMap[value] <= RomanMap[rune(s[index-1])]{
			result += RomanMap[value]
		}else{
			result =result - 2 * RomanMap[rune(s[index-1])] + RomanMap[value]
		}
	}
	return  result
}