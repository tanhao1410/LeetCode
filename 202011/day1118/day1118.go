package main

func main() {

}

//每日一题：134. 加油站
func canCompleteCircuit(gas []int, cost []int) int {
	//思路：可以从那一站出发，并走一圈，如果不行，下一个，都不行，返回-1
	for i := 0; i < len(gas); i++ {
		all := gas[i] //起始油量
		cur := i
		for j := (i + 1) % len(gas); cost[cur] <= all; j = (j + 1) % len(gas) {
			if j == i {
				return i
			}
			all += gas[j]    //补充油
			all -= cost[cur] //消耗油
			cur = j
		}
	}

	return -1
}
