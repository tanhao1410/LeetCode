package main

import (
	"container/list"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numberToWords(10000080))
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
