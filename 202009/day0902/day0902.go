package main

import (
	"fmt"
	"strconv"
)

func main() {
	//nums := []int{0,1,2,4,5,7}
	//fmt.Println(summaryRanges(nums))
	grid := [][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}
	fmt.Println(countServers(grid))
}

//三角形的最小路径和
func minimumTotal(triangle [][]int) int {
	//动态规划算法：dp[i] 第i步走的哪dp[0]=0,dp[i]  val 代表数组的下标
	//方法不行，每一步的最小，不代表最后一步最小，解决方式，看这一行中最小的
	dp := make([]int, len(triangle))
	dp[0] = 0

	for i := 1; i < len(triangle); i++ {
		//dp[i]依赖于dp[i-1],若

	}
	return 0
}

//统计参与通信的服务器
func countServers(grid [][]int) int {
	//思路：统计每一行有两个及以上的数量，统计每一列有两个以上的数量。问题在于可能会重复计算了
	//1.用一个map做记录key:251*i + j 2.思路，将已统计的矩阵中的数改为2
	//先统计行
	count := 0
	for i := 0; i < len(grid); i++ {
		rowNum, overOne, tempPosition := 0, false, 0
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				if overOne {
					grid[i][j] = 2 //不能直接给2的，超过一个才能给
					rowNum++
				} else {
					//说明是每一行的第一个，处理一下
					overOne = true
					tempPosition = i*251 + j
				}
			}
		}
		if rowNum > 0 {
			count += rowNum + 1
			grid[tempPosition/251][tempPosition%251] = 2
		}
	}

	//再统计列
	for i := 0; i < len(grid[0]); i++ {
		colNum, haveJoin := 0, 0
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == 1 {
				colNum++
			} else if grid[j][i] == 2 {
				haveJoin++
			}
		}
		if haveJoin+colNum > 1 {
			count += colNum
		}
	}

	return count
}

//[0,1,2,4,5,7,8] 无重复，有序 ["0->2","4->5","7"]
func summaryRanges(nums []int) []string {
	res := []string{}
	//i每次走一步，j只有i走到了新连续时，j跟随i
	for i, j, s := 0, 0, ""; i < len(nums); i++ {

		if i == j {
			s = strconv.Itoa(nums[i])
		}

		if i-j != nums[i]-nums[j] {
			//不连续了
			if i-j > 1 {
				s = s + "->" + strconv.Itoa(nums[i-1])
			}
			res = append(res, s)
			j, i, s = i, i-1, ""
		}

		//看j是否是最后一个
		if j == len(nums)-1 {
			s = strconv.Itoa(nums[j])
			res = append(res, s)
			break
		}
		if i == len(nums)-1 {
			s = s + "->" + strconv.Itoa(nums[i])
			res = append(res, s)
			break
		}
	}
	return res
}
