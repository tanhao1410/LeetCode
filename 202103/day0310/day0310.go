package main

import "container/list"

func main() {

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
