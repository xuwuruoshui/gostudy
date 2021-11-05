package main

import (
	"fmt"
	"math"
)

// =====================
// https://leetcode-cn.com/problems/coin-change/
//======================

func main() {
	coins := []int{1,2,5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {

	// 1.获取硬币的面值个数，创建一个数组获取结果
	length := len(coins)
	f := make([]int,amount + 1)

	var i int
	var j int
	f[0] = 0
	// 遍历结果切片
	for i = 1; i < len(f); i++ {
		f[i] = math.MaxInt
		// 遍历硬币切片
		for j = 0; j < length; j++ {
			// 1.必须剩余总值>=当前硬币面值
			// 2.当前硬币不能还是无穷大
			// 3.最重要的一步,取f(总值-当前面值)+1和上次遍历后的f(总值-当前面值)+1的最小值
			if i >= coins[j] && f[i-coins[j]] != math.MaxInt && f[i-coins[j]]+1 < f[i] {
				f[i] = f[i-coins[j]]+1
			}
		}
	}

	if f[amount] == math.MaxInt {
		return -1
	}else{
		return f[amount]
	}

	
}
