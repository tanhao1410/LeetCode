package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	codec := Constructor()
	root := codec.deserialize("[]")
	serialize := codec.serialize(root)
	fmt.Println(serialize)
}

//889. 根据前序和后序遍历构造二叉树
func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}

	//第一个节点为
	root := &TreeNode{
		Val: pre[0],
	}

	if len(pre) == 1 {
		return root
	}

	//将pre,post分成两个，要求新
	for i := 0; ; i++ {
		if post[i] == pre[1] {
			//找到了要切割的地方了
			root.Left = constructFromPrePost(pre[1:2+i], post[:i+1])
			root.Right = constructFromPrePost(pre[2+i:], post[i+1:len(post)-1])
			return root
		}
	}
}

//297. 二叉树的序列化与反序列化
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.[1,2,3,null,null,4,5]
func (this *Codec) serialize(root *TreeNode) string {
	res := "["
	//树不为空
	if root != nil {
		queue := []*TreeNode{root}
		//层次遍历
		for queueLen := len(queue); queueLen > 0; queueLen = len(queue) {

			for i := 0; i < queueLen; i++ {
				if queue[i] != nil {
					res += strconv.Itoa(queue[i].Val) + ","

					queue = append(queue, queue[i].Left, queue[i].Right)
				} else {
					res += "null,"
				}
			}
			queue = queue[queueLen:]
		}

	}

	//去除后面的,null
	res = strings.TrimRight(res, ",null")
	//去除最后一个逗号
	if res[len(res)-1] == ',' {
		res = res[:len(res)-1]
	}

	res += "]"
	return res
}

// Deserializes your encoded data to tree.[1,2,3,null,null,4,5]
func (this *Codec) deserialize(data string) *TreeNode {

	data = data[1 : len(data)-1]
	if len(data) == 0 {
		return nil
	}

	datas := strings.Split(data, ",")

	rootVal, _ := strconv.Atoi(datas[0])
	root := &TreeNode{
		Val: rootVal,
	}

	queue := []*TreeNode{root}

	index := 1
	//从queue中取出一个
	for queueLen := len(queue); queueLen > 0; queueLen = len(queue) {

		for i := 0; i < queueLen; i++ {

			curNode := queue[i]

			if index < len(datas) && datas[index] != "null" {

				indexVal, _ := strconv.Atoi(datas[index])

				curNode.Left = &TreeNode{
					Val: indexVal,
				}

				queue = append(queue, curNode.Left)
			}

			if index+1 < len(datas) && datas[index+1] != "null" {

				indexVal, _ := strconv.Atoi(datas[index+1])

				curNode.Right = &TreeNode{
					Val: indexVal,
				}

				queue = append(queue, curNode.Right)
			}

			index += 2
		}

		queue = queue[queueLen:]

	}

	return root
}
