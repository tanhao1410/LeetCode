package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(minimumOperations("yyyrryyy"))
}

//秋叶收藏集 r*y*r*
func minimumOperations(leaves string) int {

	//思路：开头和结尾的r都去掉，最后剩的就是y..y这样的字符串，若这样的字符串不存在，说明没有y ，结果为1
	//若有，则只需处理这个串即可。问题转化为了，把y..y==>所有的y是连续的。中间有多少r，有多少y呢。
	//1.把中间的r-->y ==>能直接求出来
	//2.以开头的y为基准，后面的全变r ==>也容易求出来
	//3.以结尾的y为基准，前面的全变r ==>也容易求出来
	//4.里面r比较多，且有y的情况适用，以中间的某个y为基准，两边全变r ，y...y...y ==>r...y...r


	//思路2：开头和结尾去掉，都变成r.中间的字母要么变成收藏集，要么y*在开头或结尾，其余的变成r.
	//1.先看开头或结尾，
	res := 0
	if leaves[0] != 'r'{
		res++
	}
	if leaves[len(leaves) -1] != 'r'{
		res ++
	}
	newLeaves := leaves[1:len(leaves)-1]

	count1,count2 := 0,0
	if len(newLeaves)>=3{
		count1 = minimumOperations(newLeaves)

		//y*在开头

	}else{
		if strings.Contains(newLeaves,"y"){
			count2=0
		}else{
			count2= 1
		}
	}

	if count1> count2{
		return res+count2
	}
	return res+count1
}