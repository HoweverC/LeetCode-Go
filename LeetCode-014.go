package main

/*编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，返回空字符串 ""*/
/*解法一：逐位比较*/
/*先找出字符串数组中最短的字符串长度，在这个长度范围中比较每个字符串相同位置的字符是否相同*/

import (
	"fmt"
)
func main()  {
	strs := []string{"flower","flow","flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0{
		return ""
	}
	min :=len(strs[0])
	temp := 0

	for index, _ := range strs {
		temp =len(strs[index])
		if temp < min {
			min = len(strs[index])
		}
	}
	result :=""
	for i := 0 ; i < min ; i++{
		for j := 0 ; j < len(strs) - 1 ; j++{
			if strs[j][i] != strs[j+1][i] {
				return result
			}
		}
		result += string(strs[0][i])
	}
	return result
}