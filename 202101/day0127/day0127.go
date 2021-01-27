package main

func main() {

}

//676. 实现一个魔法字典
type MagicDictionary struct {
	m map[int][]string
}

func Constructor() MagicDictionary {
	return MagicDictionary{m: map[int][]string{}}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, s := range dictionary {
		sLen := len(s)
		if _, ok := this.m[sLen]; ok {
			this.m[sLen] = append(this.m[sLen], s)
		} else {
			this.m[sLen] = []string{s}
		}
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	//从同长度的里面找
	sLen := len(searchWord)
	if words, ok := this.m[sLen]; ok {
		for _, word := range words {
			diff := 0
			//判断是否仅一个字母不相同
			for i := 0; i < sLen; i++ {
				if word[i] != searchWord[i] {
					diff++
				}
			}
			if diff == 1 {
				return true
			}
		}
	}
	return false
}
