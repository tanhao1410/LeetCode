package main

func main() {

}

//面试题 16.16. 部分排序
func subSort(array []int) []int {
	//先分别找从前从后哪个位置开始无序的
	mm, nn := -1, -1
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			mm = i
			break
		}
	}

	for i := len(array) - 2; i >= 0; i-- {
		if array[i] > array[i+1] {
			nn = i
			break
		}
	}

	//有序的
	if mm == -1 || nn == -1 {
		return []int{-1, -1}
	}

	//必有 m <= mm - 1 ,n >= nn + 1
	//寻找[mm,nn]中的最大值和最小值，
	min, max := array[mm-1], array[nn+1]
	for i := mm - 1; i <= nn+1; i++ {
		if array[i] < min {
			min = array[i]
		}

		if array[i] > max {
			max = array[i]
		}
	}

	//看min应在的位置,m应该尽可能的大。即如果碰到相等的，应该继续走。
	for i := 0; i <= len(array); i++ {
		if array[i] > min {
			//最小的数可以放在它的前面
			mm = i
			break
		}
	}

	if max > array[len(array)-1] {
		nn = len(array) - 1
	} else {
		//看max应在的位置,max 可以放到最后的
		for i := nn + 1; i < len(array); i++ {
			if array[i] >= max {
				nn = i - 1
				break
			}
		}
	}

	return []int{mm, nn}
}

//每日一题：456. 132模式
func find132pattern(nums []int) bool {

	//新思路：对于每一个数，可以很方便的找到，它左边的比它小的最小数，它右边比它小的最大数，如果存在最大数比最小数要大，说明存在。存下标
	if len(nums) < 3 {
		return false
	}
	min, max := make([]int, len(nums)), make([]int, len(nums))
	min[0] = -1
	max[len(nums)-1] = -1
	//更新min
	for i := 1; i < len(nums); i++ {
		//若它的前面不存在最小数，若前面比当前小，则为当前，否则没有
		//若它的前面存在最小数，
		if min[i-1] == -1 {
			if nums[i] > nums[i-1] {
				min[i] = i - 1
			} else {
				min[i] = -1
			}

		} else {
			//若前面存在最小数，如果我比前面的最小数还要小，那么为-1
			//如果比前面最小数，要大，那么 为前面最小数
			if nums[i] < nums[min[i-1]] {
				min[i] = -1
			} else {
				min[i] = min[i-1]
			}
		}
	}
	for i := len(nums) - 2; i >= 0; i-- {

		maxItem := -1
		for j := i + 1; j < len(nums); j++ {
			//比它小
			if nums[j] < nums[i] {
				if maxItem == -1 {
					maxItem = j
				} else {
					if nums[maxItem] < nums[j] {
						maxItem = j
					}
				}
			}
		}
		max[i] = maxItem
	}
	//新思路：对于每一个数，可以很方便的找到，它左边的比它小的最小数，它右边比它小的最大数，如果存在最大数比最小数要大，说明存在。存下标
	for i := 1; i < len(nums)-1; i++ {
		if max[i] != -1 && min[i] != -1 && nums[max[i]] > nums[min[i]] {
			return true
		}
	}
	return false
}
