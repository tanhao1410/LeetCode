package main

import (
	"container/list"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numberToWords(10000080))
	fmt.Println(solveNQueens(4))
}

//面试题 03.06. 动物收容所
type AnimalShelf struct {
	Animals  list.List
	FirstCat *list.Element
	FirstDog *list.Element
}

func Constructor() AnimalShelf {
	return AnimalShelf{
		Animals:  list.List{},
		FirstCat: nil,
		FirstDog: nil,
	}
}

func (this *AnimalShelf) Enqueue(animal []int) {
	this.Animals.PushBack(animal)

	if this.FirstCat == nil && animal[1] == 0 {
		this.FirstCat = this.Animals.Back()
	} else if this.FirstDog == nil && animal[1] == 1 {
		this.FirstDog = this.Animals.Back()
	}

}

func (this *AnimalShelf) DequeueAny() []int {
	if this.Animals.Len() == 0 {
		return []int{-1, -1}
	}
	front := this.Animals.Front()
	if this.FirstDog == front {
		this.FirstDog = nil
		for p := front.Next(); p != nil; p = p.Next() {
			if p.Value.([]int)[1] == 1 {
				this.FirstDog = p
				break
			}
		}

	} else if this.FirstCat == front {
		this.FirstCat = nil
		for p := front.Next(); p != nil; p = p.Next() {
			if p.Value.([]int)[1] == 0 {
				this.FirstCat = p
				break
			}
		}
	}
	this.Animals.Remove(front)

	//element转成原有的格式
	return front.Value.([]int)
}

func (this *AnimalShelf) DequeueDog() []int {
	if this.FirstDog == nil {
		return []int{-1, -1}
	}

	res := this.FirstDog
	this.FirstDog = nil
	for p := res.Next(); p != nil; p = p.Next() {
		if p.Value.([]int)[1] == 1 {
			this.FirstDog = p
			break
		}
	}

	this.Animals.Remove(res)
	return res.Value.([]int)
}

func (this *AnimalShelf) DequeueCat() []int {
	if this.FirstCat == nil {
		return []int{-1, -1}
	}

	res := this.FirstCat
	this.FirstCat = nil
	for p := res.Next(); p != nil; p = p.Next() {
		if p.Value.([]int)[1] == 0 {
			this.FirstCat = p
			break
		}
	}

	this.Animals.Remove(res)
	return res.Value.([]int)
}

//面试题 16.08. 整数的英语表示
func numberToWords(num int) string {
	dict1 := map[int]string{
		1: "One", 2: "Two", 3: "Three", 4: "Four", 5: "Five", 6: "Six", 7: "Seven", 8: "Eight", 9: "Nine", 0: "Zero",
		10: "Ten", 11: "Eleven", 12: "Twelve", 13: "Thirteen", 14: "Fourteen", 15: "Fifteen", 16: "Sixteen", 17: "Seventeen",
		18: "Eighteen", 19: "Nineteen",
	}

	dic2 := map[int]string{
		2: "Twenty", 3: "Thirty", 4: "Forty", 5: "Fifty", 6: "Sixty", 7: "Seventy", 8: "Eighty", 9: "Ninety",
	}

	res := ""

	if num >= 1000000000 {
		pre := num / 1000000000
		res += numberToWords(pre)
		res += " Billion "
		num %= 1000000000
		if num == 0 {
			return strings.Trim(res, " ")
		}
	}

	if num >= 1000000 {
		pre := num / 1000000
		res += numberToWords(pre)
		res += " Million "
		num %= 1000000
		if num == 0 {
			return strings.Trim(res, " ")
		}
	}

	if num >= 1000 {
		pre := num / 1000
		res += numberToWords(pre)
		res += " Thousand "
		num %= 1000
		if num == 0 {
			return strings.Trim(res, " ")
		}
	}

	if num >= 100 {
		pre := num / 100
		res += numberToWords(pre)
		res += " Hundred "
		num %= 100
		if num == 0 {
			return strings.Trim(res, " ")
		}
	}

	if num < 20 {
		res += dict1[num]
		return strings.Trim(res, " ")
	}

	res += dic2[num/10]

	if num%10 > 0 {
		res += " "
		res += dict1[num%10]
	}

	return strings.Trim(res, " ")
}

//面试题 08.12. 八皇后
func solveNQueens(n int) [][]string {

	res := [][]string{}
	//用一个数组来记录每行放置的列号
	m := make([]int, n)

	nextLocation(m, 0, &res)

	return res
}

func nextLocation(m []int, l int, res *[][]string) {

	//说明放置成功了
	if l == len(m) {
		//把数字转成对应的棋盘形式
		resItem := []string{}
		for i := 0; i < l; i++ {
			row := ""
			for j := 0; j < l; j++ {
				if j == m[i] {
					row += "Q"
				} else {
					row += "."
				}
			}
			resItem = append(resItem, row)
		}
		*res = append(*res, resItem)
	}

	//判断是否可以放置，1，不能和前面相等，2，不能斜对角相等
out:
	for i := 0; i < len(m); i++ {

		for j := 0; j < l; j++ {
			if m[j] == i || l-j == m[j]-i || l-j == -m[j]+i {
				continue out
			}
		}

		m[l] = i

		//放置下一个
		nextLocation(m, l+1, res)

	}

}
