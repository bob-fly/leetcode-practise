package main

import "fmt"

/**
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

 

示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
示例 3：

输入：digits = [0]
输出：[1]
 

提示：

1 <= digits.length <= 100
0 <= digits[i] <= 9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/plus-one
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func main()  {
	nums:=[]int {1,1,2,0}
	nums2:=make([]int, len(nums))
	copy(nums2,nums)
	fmt.Printf("before: %v, after: %v",nums, plusOne(nums2))

}

func plusOne(digits []int) []int{
	if digits ==nil || len(digits)<0 {
		return nil
	}
	return doCycle(digits,len(digits)-1)
}

func doCycle(nums []int ,index int) []int {
	//退出条件
	if nums[index] !=9 {
		nums[index]+=1
		return nums
	}
	// num ==9
	nums[index]=0

	//首位为9
	if index ==0 {
		//数组后移
		copyNums := make([]int, len(nums)+1)
		copy(copyNums, nums)
		for i := 1; i < len(copyNums); i++ {
			copyNums[i] = nums[i-1]
		}
		copyNums[0] = 1
		return copyNums
	}
	return doCycle(nums,index-1)
}
