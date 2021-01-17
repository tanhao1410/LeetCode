package main

func main() {

}

//每日一题：1232. 缀点成线
func checkStraightLine(coordinates [][]int) bool {
	disX := coordinates[0][0] - coordinates[1][0]
	disY := coordinates[0][1] - coordinates[1][1]
	for i := 2; i < len(coordinates); i++ {
		if disX*(coordinates[0][1]-coordinates[i][1]) != disY*(coordinates[0][0]-coordinates[i][0]) {
			return false
		}
	}

	return true
}
