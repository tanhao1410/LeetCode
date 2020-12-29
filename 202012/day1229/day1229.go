package main

import "fmt"

func main() {
	fmt.Println(pathWithObstacles([][]int{{0, 1, 0}, {0, 1, 1}, {0, 0, 0}}))
	fmt.Println(minPatches([]int{1, 1, 2, 3, 5, 10}, 200))
	fmt.Println(smallestK([]int{4, 5}, 4))
}

//面试题 17.14. 最小K个数
func smallestK(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}
	//思路：通过快速排序的思想进行
	//1.先根据第一个数进行划分，将数组划分成两段，前面都是小于它的，后面都是大于等于它的
	i, j := 1, len(arr)-1
	for j >= i {

		for i < len(arr) && arr[i] < arr[0] {
			i++
		}

		for j > 0 && arr[j] >= arr[0] {
			j--
		}

		//交换
		if j > i {
			arr[i], arr[j] = arr[j], arr[i]
		}

	}
	if j >= 0 {
		arr[0], arr[j] = arr[j], arr[0]
	}

	if j == k-1 {
		return arr[:k]
	} else if j > k-1 {
		return smallestK(arr[:j], k)
	} else {
		return append(arr[:j+1], smallestK(arr[j+1:], k-j-1)...)
	}
}

//每日一题：330. 按要求补齐数组
func minPatches(nums []int, n int) int {
	//思路：肯定需要1，没有的话要加上。除非有多个1，否则2也是必需的。
	//先看前面能组成多少一下的任意数 ，注意1,2,4,8这些，如果形成有形成不了的数了，就应该添上
	res := 0
	//先根据n生成一般的数组合，后面数之和

	used := make([]int, len(nums))
	//没有被使用的数字的和
	//getNoUsedSum := func() int {
	//	res := 0
	//	for k, v := range used {
	//		if v == 0 {
	//			res += nums[k]
	//		}
	//	}
	//	return res
	//}

	//能组成的任意数
	r := 0
	//先看前面能形成多少的任意数
	if len(nums) == 0 || nums[0] != 1 {
		res++
		r = 1
	} else {
		r = 1
		//该1已被使用了
		used[0] = 1
	}

	//得到比N大的最近的2的幂，1，2,4,8
	get2N := func(n int) int {
		for res := 2; ; res <<= 1 {
			if res > n {
				return res
			}
		}
	}

	for all := r; all < n; all = r {

		flag := false
		for i := 0; i < len(nums); i++ {
			if used[i] == 0 && r+1 == nums[i] {
				r = r + 1
				used[i] = 1
				flag = true
				break
			}
		}

		if !flag {
			//不一定每一次都要从里面补，如果后面还有这样的数呢
			a := get2N(r)
			res++
			r = a<<1 - 1
		}
	}

	return res
}

//面试题 08.02. 迷路的机器人
func pathWithObstacles(obstacleGrid [][]int) [][]int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return nil
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	res := [][]int{}
	//返回一条可行的即可，dp代表某个结点是否能走到终点。
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	copy(dp, obstacleGrid)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[i][j] = obstacleGrid[i][j]
			} else if obstacleGrid[i][j] == 0 && ((i < m-1 && dp[i+1][j] == 0) || (j < n-1 && dp[i][j+1] == 0)) {
				dp[i][j] = 0
			} else {
				dp[i][j] = 1
			}
		}
	}
	if dp[0][0] == 1 {
		return nil
	}
	for i, j := 0, 0; i < m-1 && j < n-1; {
		res = append(res, []int{i, j})
		//向下走
		if i+1 < m && dp[i+1][j] == 0 {
			i = i + 1
		} else if j+1 < n && dp[i][j+1] == 0 {
			//向右走
			j = j + 1
		}
	}
	res = append(res, []int{m - 1, n - 1})
	return res
}
