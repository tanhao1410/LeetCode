package main

func main() {

}

//每日一题：547. 省份数量
func findCircleNum(isConnected [][]int) int {
	res := 0
	m := make(map[int]map[int]bool)
	//第一次遍历形成直接相连的两城市
	for i := 0; i < len(isConnected); i++ {
		m[i] = make(map[int]bool)
		for j := 0; j < len(isConnected); j++ {
			if i != j && isConnected[i][j] == 1 {
				m[i][j] = true
			}
		}
	}
	//已经加入到某个省份的城市集合
	alreadyCitys := make(map[int]bool)

	//从第一个城市开始算起
	for i := 0; i < len(isConnected); i++ {
		//没有加入集合的才会计算
		if !alreadyCitys[i] {
			res++
			for noAdd := false; !noAdd; {
				//遍历它所有的能连接的城市
				noAdd = true
				for city, _ := range m[i] {
					//没有遍历过的城市
					if !alreadyCitys[city] {
						alreadyCitys[city] = true
						for city2, _ := range m[city] {
							if !m[i][city2] {
								//非直接相连的城市，将其加入进来，下次遍历时就会遍历到它
								m[i][city2] = true
								noAdd = false
							}
						}
					}
				}
			}
		}
	}

	return res
}
