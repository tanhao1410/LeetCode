package main

import "fmt"

func main() {
	fmt.Println(characterReplacement("AABABBA", 1))
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
