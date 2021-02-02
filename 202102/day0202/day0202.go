package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(characterReplacement("AABABBA", 1))
}

type Num struct {
	val   int
	index int
}

//1331. 数组序号转换
func arrayRankTransform(arr []int) []int {
	//思路：排序。用怎么安排数字，原来的位置信息还能记录吗
	res := make([]int, len(arr))
	arrNums := make([]Num, len(arr))
	for i := 0; i < len(arr); i++ {
		arrNums[i] = Num{
			val:   arr[i],
			index: i,
		}
	}
	sort.Slice(arrNums, func(i, j int) bool {
		return arrNums[i].val < arrNums[j].val
	})

	seq := 1

	for i := 0; i < len(arrNums); i++ {
		res[arrNums[i].index] = seq
		//如果为最后一个了，或者，下一个和当前不相等
		if i == len(arrNums)-1 || arrNums[i].val != arrNums[i+1].val {
			seq++
		}
	}

	return res
}

//208. 实现 Trie (前缀树)
type Trie struct {
	tries   []*Trie
	endFlag bool
}

func Constructor3() Trie {
	return Trie{
		tries: make([]*Trie, 26),
	}
}

func (this *Trie) Insert(word string) {
	curTrie := this
	for i := 0; i < len(word); i++ {
		if curTrie.tries[word[i]-'a'] == nil {
			curTrie.tries[word[i]-'a'] = &Trie{tries: make([]*Trie, 26)}
		}
		curTrie = curTrie.tries[word[i]-'a']
	}
	curTrie.endFlag = true
}

func (this *Trie) Search(word string) bool {
	curTrie := this
	for i := 0; ; i++ {
		if curTrie.tries[word[i]-'a'] == nil {
			return false
		}
		curTrie = curTrie.tries[word[i]-'a']
		if i == len(word)-1 {
			return curTrie.endFlag
		}
	}
}

func (this *Trie) StartsWith(prefix string) bool {
	curTrie := this
	for i := 0; ; i++ {
		if curTrie.tries[prefix[i]-'a'] == nil {
			return false
		}
		curTrie = curTrie.tries[prefix[i]-'a']
		if i == len(prefix)-1 {
			return true
		}
	}
}

//每日一题：424. 替换后的最长重复字符
func characterReplacement(s string, k int) int {
	res := 0
	return res
}
