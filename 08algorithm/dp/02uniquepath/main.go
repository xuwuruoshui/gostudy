package main

// =====================
// https://leetcode-cn.com/problems/unique-paths/
//======================

func main(){
	println(uniquePaths(3,7))
}

// m代表向右,n代表向下
func uniquePaths(m int, n int) int {

	// 定义一个二维切片
	var result = make([][]int,m)
	for i := 0; i <m; i++ {
		result[i] = make([]int,n)
	}


	// 遍历得到result[m-1][n-1]
	for i := 0; i <m; i++ {
		for j := 0; j <n; j++ {		
			// 边界情况	
			if i==0 || j==0 {
				result[i][j]=1
			}else{
				// 最终结果的公式
				result[i][j] = result[i-1][j] + result[i][j-1]
			}
		}
	}

	return result[m-1][n-1]
}