package main

import "fmt"

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println(groupAnagrams(strs))
}

//54.螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	res := []int{}
	m:=len(matrix)
	if m == 0{
		return res
	}
	n := len(matrix[0])
	if n == 0{
		return res
	}

	x,y := 0,-1
	for {

		//向右走n步
		if n <= 0{
			break
		}

		for j:=0;j < n;j ++{
			y++
			res = append(res,matrix[x][y])
		}

		if m -1 <= 0{
			break
		}

		//向下走m -1步
		for j := 0;j < m -1;j ++{
			x++
			res = append(res,matrix[x][y])
		}

		if  n-1 <= 0{
			break
		}

		//向左走n-1步
		for j := 0;j < n-1;j ++{
			y--
			res = append(res,matrix[x][y])
		}

		if   m -2 <= 0{
			break
		}

		//向上走m -2步
		for j := 0;j < m -2;j++{
			x--
			res = append(res,matrix[x][y])
		}

		m-=2
		n-=2
	}

	return res
}

//字母异味词分组
func groupAnagrams(strs []string) [][]string {
	//打散，排序，hash
	res := [][]string{}
	//key:打散排序后的字符串，value：所有异分词集合
	m  := make(map[string][]string)
	for i:= 0 ;i < len(strs);i++{
		sortedString := CreateSortedString(strs[i])
		if item,ok:= m[sortedString];ok{
			item = append(item, strs[i])
			m[sortedString] = item
		}else{
			item := []string{strs[i]}
			m[sortedString] = item
		}
	}

	for _,v := range m{
		res = append(res, v)
	}

	return res
}

func CreateSortedString(s string) string{
	bytes := []byte(s);
	for i:=0;i < len(bytes) - 1;i ++{
		for j:=i;j< len(bytes);j++{
			if bytes[j] < bytes[i]{
				bytes[j],bytes[i] = bytes[i],bytes[j]
			}
		}
	}
	return string(bytes)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//环形链表2
func detectCycle(head *ListNode) *ListNode {

	if head == nil{
		return head
	}

	for fast,slow := head,head;fast != nil;{
		slow,fast = slow.Next,fast.Next
		if fast == nil{
			return nil
		}
		fast = fast.Next

		//说明有环了，在此相遇
		if fast == slow{
			for point := head;point != slow;{
				point,slow = point.Next,slow.Next
			}
			return slow
		}
	}

	return nil
}
