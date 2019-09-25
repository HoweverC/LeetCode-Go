package main

/*给定一个整数数组 nums 和一个目标 target，在该数组中找出和为目标值的那两个整数，并返回他们的数组下标*/
/*解法一：遍历两个数组，求和*/
import "fmt"
func main()  {
	var nums =[]int {2,7,11,15}
	fmt.Println(twoSum(nums, 9))
}

func twoSum(nums []int, target int) []int {
	for firstIndex , firstNumber := range nums{
		for secondIndex , secondNumber := range nums {
			if firstIndex != secondIndex && firstNumber + secondNumber == target{
				return []int{firstIndex, secondIndex}
			}
		}
	}
	return nil
}