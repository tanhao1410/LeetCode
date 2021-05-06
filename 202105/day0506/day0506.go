package main

func main() {

}

//1720. 解码异或后的数组
func decode(encoded []int, first int) []int {
	res := make([]int, len(encoded)+1)
	res[0] = first
	for i := 1; i <= len(encoded); i++ {
		res[i] = res[i-1] ^ encoded[i-1]
	}
	return res
}
