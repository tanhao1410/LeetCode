package main

func main() {

}

//208. 实现 Trie (前缀树)
type Trie struct {
	End  bool
	Next []*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		End:  false,
		Next: make([]*Trie, 26),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	bytes := []byte(word)
	for i := 0; i < len(bytes); i++ {

		if cur.Next[bytes[i]-'a'] == nil {
			cur.Next[bytes[i]-'a'] = &Trie{
				End:  false,
				Next: make([]*Trie, 26),
			}
		}

		cur = cur.Next[bytes[i]-'a']
	}

	cur.End = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	bytes := []byte(word)
	for i := 0; i < len(bytes); i++ {

		if cur.Next[bytes[i]-'a'] == nil {
			return false
		}

		cur = cur.Next[bytes[i]-'a']
	}

	return cur.End
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	bytes := []byte(prefix)
	for i := 0; i < len(bytes); i++ {

		if cur.Next[bytes[i]-'a'] == nil {
			return false
		}

		cur = cur.Next[bytes[i]-'a']
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
