package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValidSerialization2("9,3,4,#,#,1,#,#,2,#,6,#,#"))
}

//331. 验证二叉树的前序序列化
func isValidSerialization(preorder string) bool {
	if preorder == "#" {
		return true
	}
	//思路：一个数字后面紧接两个#,那么他们三个可以看做一个#，如果最后整个字符串变为一个#，则为true
	nodes := strings.Split(preorder, ",")
	for i := 0; i < len(nodes); i++ {
		//是数字的话，看下它紧接着的后面两个是否是#
		if nodes[i] != "#" && i+2 < len(nodes) && nodes[i+1] == "#" && nodes[i+2] == "#" {
			newPreorder := strings.Join(nodes[:i], ",")
			newPreorder += ","
			newPreorder += strings.Join(nodes[i+2:], ",")
			return isValidSerialization(strings.Trim(newPreorder, ","))
		}
	}

	return false
}

//331. 验证二叉树的前序序列化
func isValidSerialization2(preorder string) bool {
	if preorder == "#" {
		return true
	}
	//思路：一个数字后面紧接两个#,那么他们三个可以看做一个#，如果最后整个字符串变为一个#，则为true
	nodes := strings.Split(preorder, ",")
	newPreorder := ""
	for i := 0; i < len(nodes); i++ {
		//是数字的话，看下它紧接着的后面两个是否是#
		if nodes[i] != "#" && i+2 < len(nodes) && nodes[i+1] == "#" && nodes[i+2] == "#" {
			newPreorder += "#"
			//跳过两个
			i += 2
		} else {
			newPreorder += nodes[i]
		}
		if i != len(nodes)-1 {
			newPreorder += ","
		}
	}
	fmt.Println(newPreorder)
	//没有一个可以消除的，直接返回false
	if newPreorder == preorder {
		return false
	}

	return isValidSerialization2(newPreorder)
}
