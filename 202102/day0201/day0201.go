package main

import "fmt"

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	fmt.Println(cache.Get(1))

	dic := Constructor2()
	dic.AddWord("bad")
	dic.AddWord("dad")
	dic.AddWord("mad")
	fmt.Println(dic.Search("pad"))
	fmt.Println(dic.Search("bad"))
}

//211. 添加与搜索单词 - 数据结构设计
type WordDictionary struct {
	tries []*WordDictionary
	flag  bool
}

func Constructor2() WordDictionary {
	return WordDictionary{
		tries: make([]*WordDictionary, 26),
	}
}

func (this *WordDictionary) AddWord(word string) {
	//
	curTrie := this
	for i := 0; ; i++ {
		if curTrie.tries[word[i]-'a'] == nil {
			curTrie.tries[word[i]-'a'] = &WordDictionary{
				tries: make([]*WordDictionary, 26),
			}
		}

		curTrie = curTrie.tries[word[i]-'a']

		//最后一个字母的话，添加字符串
		if i == len(word)-1 {
			curTrie.flag = true
			break
		}

	}
}

func (this *WordDictionary) Search(word string) bool {

	curTrie := this
	for i := 0; i < len(word); i++ {
		//可以代替任何字母
		if word[i] == '.' {
			for _, v := range curTrie.tries {
				if v != nil && v.Search(word[i+1:]) {
					return true
				}
			}
			return false
		}

		//没有该字母
		if curTrie.tries[word[i]-'a'] == nil {
			return false
		}

		curTrie = curTrie.tries[word[i]-'a']
		//最后一个字母了
		if i == len(word)-1 {
			return curTrie.flag
		}
	}
	return curTrie.flag
}

//面试题 16.25. LRU 缓存
type LRUCache struct {
	//存储数
	m map[int]int
	//访问序列
	seq []int
	//容量
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:        make(map[int]int),
		seq:      []int{},
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if res, ok := this.m[key]; ok {
		//会更新索引
		next := this.seq[0]
		for i := 0; i < len(this.seq) && next != key; i++ {
			//往前走一位
			this.seq[i+1], next = next, this.seq[i+1]
		}
		this.seq[0] = key
		return res
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {

	//看key是否存在
	if this.Get(key) != -1 {
		//更新值
		this.m[key] = value
		//更新seq 即可。
		next := this.seq[0]
		for i := 0; i < len(this.seq) && next != key; i++ {
			//往前走一位
			this.seq[i+1], next = next, this.seq[i+1]
		}
		this.seq[0] = key
		return
	}

	//思路：如果大于容量，要删除一个
	if len(this.m) < this.capacity {
		this.seq = append([]int{key}, this.seq...)
	} else {
		//需要删除一个，删除哪一个？删除seq中的最后一项即可
		delete(this.m, this.seq[len(this.seq)-1])
		this.seq = this.seq[:len(this.seq)-1]
		this.seq = append([]int{key}, this.seq...)
	}
	this.m[key] = value
}

//每日一题：888. 公平的糖果棒交换
func fairCandySwap(A []int, B []int) []int {
	res := []int{}
	//只是交换一根，思路：先求和，
	sumA, sumB := 0, 0
	for _, v := range A {
		sumA += v
	}
	m := make(map[int]bool)
	for _, v := range B {
		sumB += v
		m[v] = true
	}
	bigThen := (sumA + sumB) / 2
	//转换成了求两数之和
	for _, v := range A {
		if m[v-bigThen] {
			res = append(res, v, v-bigThen)
			return res
		}
	}

	return res
}
