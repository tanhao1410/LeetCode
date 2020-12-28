package main

func main() {

}

//面试题 04.01. 节点间通路
func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	if n == 0 {
		return false
	}
	//思路：用一个集合来记录start能够到达的所有点，结束条件为该集合加入了targe或加入不了新的节点了
	m := make(map[int]map[int]bool) //val -->(target,bool)
	for _, v := range graph {
		if mm, ok := m[v[0]]; ok {
			mm[v[1]] = true
		} else {
			m[v[0]] = make(map[int]bool)
			m[v[0]][v[1]] = true
		}
	}
	//一趟遍历过后，就知道了那些节点可以直接到哪些节点
	//从start开始
	mAlready := make(map[int]bool)
	if set, ok := m[start]; ok {
		for flag := true; flag; {
			flag = false
			//只要有新的加入进来，就把false改为true
			for k, _ := range set {
				//已经加入过的不用重复再加了
				//把它所能到达的都加入进来
				if kset, kok := m[k]; !mAlready[k] && kok {
					mAlready[k] = true
					for kk, _ := range kset {
						if !set[kk] {
							set[kk] = true
							flag = true
						}
					}
				}
			}
			if set[target] {
				return true
			}
		}
	}
	return false
}
