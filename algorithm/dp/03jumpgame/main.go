package main

import "fmt"

// =====================
// https://leetcode-cn.com/problems/jump-game/
//======================

func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}

	// 创建切片
	f := make([]bool, len(nums))

	f[0] = true
	// 遍历赋值
	for i := 1; i < len(f); i++ {
		f[i] = false
		// 每一个结果能否跳过去
		for j := 0; j <= i; j++ {
			// 1.上一步是否能跳过来
			// 2.下一步能否跳过去
			if f[j] && i <= nums[j]+j {
				f[i] = true
			}
		}
	}

	return f[len(f)-1]
}
