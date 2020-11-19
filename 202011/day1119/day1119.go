package main

func main() {

}

//每日一题：283. 移动零
func moveZeroes(nums []int) {
	//思路：遇到不是0的数就往前移动，移动到前面第一个0的位置处用一个变量来记录
	next := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 && next > -1 {
			nums[next], nums[i] = nums[i], 0
			next += 1
		} else if nums[i] == 0 && next == -1 {
			next = i
		}
	}
}
