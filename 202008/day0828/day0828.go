package main

//判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k
func containsNearbyDuplicate(nums []int, k int) bool {
	//思路：最简单的思路，遍历，然后往后看k个数，时间，o(kn)
	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j < k+1 && i+j < len(nums); j++ {
			if nums[i] == nums[i+j] {
				return true
			}
		}
	}
	return false
}

func containsNearbyDuplicate2(nums []int, k int) bool {
	//思路：用map[nums]index
	m := make(map[int]int, k)
	for index2, v := range nums {
		if index, ok := m[v]; ok && index2-index <= k {
			//说明之前有相同的数
			return true
		} else {
			m[v] = index2
		}
	}
	return false
}

type SubrectangleQueries struct {
	Rectangle [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{
		rectangle,
	}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			this.Rectangle[i][j] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.Rectangle[row][col]
}

//比较 arr[0] 与 arr[1] 的大小，较大的整数将会取得这一回合的胜利并保留在位置 0 ，
//较小的整数移至数组的末尾。当一个整数赢得 k 个连续回合时，游戏结束，该整数就是比赛的 赢家
func getWinner(arr []int, k int) int {
	//思路：

	//用一个标志记录上一局是谁赢了
	flag := 0 //代表谁赢了
	for count := 0; ; {
		if arr[0] > arr[1] {
			if flag == arr[0] {
				//上一次也是它赢
				count++
			} else {
				count = 1
				flag = arr[0]
			}
			if count == k || count >= len(arr) { //说明都比较一遍了，后面不用再比较了，肯定是第一个元素最大
				return arr[0]
			}
			//arr[1]移到最后
			arr[len(arr)-1] = arr[1]
			for j := 1; j < len(arr)-1; j++ {
				arr[j] = arr[j+1]
			}
		} else {
			if flag == arr[1] {
				//上一次也是它赢
				count++
			} else {
				count = 1
				flag = arr[1]
			}
			if count == k {
				return arr[1]
			}
			temp := arr[0]
			arr = arr[1:]
			arr = append(arr, temp)
		}
	}

}

//2,1,3,5,4,6,7
func getWinner2(arr []int, k int) int {
	//思路：数组可能很大，每次都真的移动的话，会导致超时。因此需要新的办法
	//两个指针
	first, second := 0, 1

	firstCount, secondCount := 0, 0
	for {

		if arr[first] > arr[second] {
			if first > second {
				second = first + 1
			} else {
				second = second + 1
			}
			firstCount++
			secondCount = 0
		} else {
			if first > second {
				first = first + 1
			} else {
				first = second + 1
			}
			firstCount = 0
			secondCount++
		}

		if firstCount >= k {
			return arr[first]
		} else if secondCount >= k {
			return arr[second]
		}

		//有走到最后了
		if first >= len(arr) {
			return arr[second]
		} else if second >= len(arr) {
			return arr[first]
		}

	}

}

//RLUD,ru
func judgeCircle(moves string) bool {
	r, u := 0, 0
	for i := 0; i < len(moves); i++ {
		switch i {
		case 'R':
			r++
		case 'L':
			r--
		case 'U':
			u++
		case 'D':
			u--
		}
	}
	return r == 0 && u == 0
}
