package main

import "fmt"

func main() {

}

//最短的回文串，在字符串前面加字符aacecaaa->aaacecaaa
func shortestPalindrome(s string) string {
	//思路：判断是否是回文串，双指针即可。
	//同样采用双指针 的思路，后面指针不停的往前走，前面指针如果和后面指针指向的相同，则向后走，不相同，则插入一个，再往后走。
	//问题是，当插入一个字符后，尾指针的指向会改动，需要处理。
	for pre, tail := 0, len(s)-1; pre < tail; tail-- {
		if s[pre] != s[tail] {
			s = s[:pre] + string(s[tail]) + s[pre:]
			fmt.Println(s)
			tail++
		}
		pre++
	}

	return s
}

//只能在前面添加字符串，不能再中间加aabba-->"abbaabba"
func shortestPalindrome2(s string) string {
	fmt.Println(len(s))

	appendCount := 0 //代表前面已经插入了几个字符了
	for pre, tail := 0, len(s)-1; pre < tail; {
		//fmt.Println(pre)
		if s[pre] != s[tail] {
			s, appendCount = string(s[len(s)-1])+s, appendCount+1
			fmt.Println(appendCount)
			//然后可以将插入的字符都改成和后面相同的了
			for i := 1; i < appendCount; i++ {
				//s[1]=s[len(s)-2] s[2]=s[len(s)-3]
				s = s[0:i] + string(s[len(s)-i-1]) + s[i+1:]
			}

			pre, tail = appendCount, len(s)-appendCount-1

		} else {
			pre, tail = pre+1, tail-1
		}
	}
	return s
}

//虽然完成了，但长字符串超时
func shortestPalindrome3(s string) string {

	appendCount := 0 //代表前面已经插入了几个字符了
	for pre, tail := 0, len(s)-1; pre < tail; {

		if s[pre] != s[tail] {

			//s,appendCount = string(s[len(s)-1]) + s,appendCount+1
			//然后可以将插入的字符都改成和后面相同的了
			appendCount++
			//s前面的字母应该是确定了的了。
			temp := []byte{}
			fmt.Println(appendCount)
			for i := len(s) - 1; i > len(s)-1-appendCount; i-- {
				temp = append(temp, s[i])
			}

			s = string(temp) + s

			pre, tail = appendCount, len(s)-appendCount-1

		} else {
			pre, tail = pre+1, tail-1
		}
	}
	return s
}

//虽然完成了，但长字符串超时
func shortestPalindrome4(s string) string {
	//如果能求出来确定要加多少个字符就可以抛弃一层循环了
	//思路：先找一个从头开始的最长的回文串，然后就知道了要加多少了。。
	i := len(s) - 1
	for ; i > 0; i-- {
		if isPalindrome(s[:i+1]) {
			//fmt.Println(i)
			break
		}
	}
	//i代表需要补几个数，len(s)-1-i

	//s前面的字母应该是确定了的了。
	temp := []byte{}
	for j := len(s) - 1; j > i; j-- {
		temp = append(temp, s[j])
	}

	//fmt.Println(temp)
	s = string(temp) + s

	return s
}

func isPalindrome(s string) bool {
	for pre, tail := 0, len(s)-1; pre < tail; {
		if s[pre] != s[tail] {
			return false
		}
		pre, tail = pre+1, tail-1
	}
	return true
}

//允许更改一个数的情况判断非递减数列 5 7 1 8 true
func checkPossibility(nums []int) bool {
	//更改数的时候注意前后
	flag := true //可以更改的数目量

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			//看能否更改 nums[i-1] <= nums[i+1]若存在的话
			if flag {
				if i-1 == -1 || nums[i-1] <= nums[i+1] {
					flag = false
				} else if i+2 == len(nums) || nums[i] <= nums[i+2] {
					//还可以更改i后面的数的
					flag = false
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}

	return true
}

//在数组中找出由三个数组成的最大乘积,不会越界，有负数,数的范围[-1000, 1000]
func maximumProduct(nums []int) int {
	//思路：找最大的三个数，如果有负数的话，需要至少拿两个负数过来。无论怎样，最大的正数都要被使用的
	//其实就是找5个数，最的正数，次大的正数，第三的正数。最小的负数，次小的负数

	//如果都是负数的话，又需要最大的几个负数了。所以上面的想法不完全对的。
	return 0
}

//从后往前动 23909
func moveZeroes(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			j := i
			for ; j < len(nums)-1 && nums[j] != 0; j++ {
				nums[j] = nums[j+1]
			}
			nums[j] = 0
		}
	}
}
