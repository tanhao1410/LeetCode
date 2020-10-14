package main

import "fmt"

func main() {
	fmt.Println(commonChars([]string{"cool","lock","cook"}))
}

//1002.查找常用字符
func commonChars(A []string) []string {
	res := []string{}

	countByte := make([]int8, 26)
	for i := 0; i < len(A[0]); i++ {
		countByte[A[0][i]-'a'] += 1
	}

	for i := 1; i < len(A); i++ {
		tempBytes := make([]int8, 26)
		for j := 0; j < len(A[i]); j++ {
			if countByte[A[i][j]-'a'] >= 1 {
				//说明该字母之前有
				tempBytes[A[i][j]-'a'] += 1
			}
		}

		for j := 0; j < 26; j++ {
			if tempBytes[j] < countByte[j] {
				countByte[j] = tempBytes[j]
			}
		}
	}

	for i := 0; i < 26; i++ {
		for j := countByte[i]; j > 0; j-- {
			res = append(res, string(i+'a'))
		}
	}

	return res
}
