package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/**螺旋矩阵2,rust解决
pub fn generate_matrix(n: i32) -> Vec<Vec<i32>> {
        //let mut res: Vec<Vec<i32>> = Vec::new();
        ////初始化
        //for i in 0..n{
        //    let mut item = Vec::new();
        //    for i in 0 .. n{
        //        item.insert(i as usize,0);
        //    }
        //    res.insert(i as usize,item);
        //}
		//定义二维向量
		let mut res= vec![vec![0;n as usize];n as usize];

        let mut count = 1;
        let mut x = 0;
        let mut y = 0;

        let mut nn = n;
        while count <= n * n{

            //先 向右走n步
            for i in 0..nn {

                res[x][y] = count;
                count += 1;
                y += 1;
            }

            y -= 1;

            //向下走n -1步
            for i in 0..nn-1{
                x += 1;
                res[x][y] = count;
                count +=1;
            }

            //向左走n-1步
            for i in 0..nn-1{
                y -=1;
                res[x][y]= count;
                count +=1;
            }

            //向上走n-2步
            for i in 0 .. nn -2{
                x-=1;
                res[x][y] = count;
                count +=1;
            }

            nn -= 2;
            y += 1;
        }

        return res;
    }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	connect2(root.Left, root.Right, nil)
	return root
}

func connect2(left, right, brother *Node) {
	if left != nil { //左孩子不为空
		if right != nil { //右也不为空

			left.Next = right
			for nextBrother := brother; nextBrother != nil; nextBrother = nextBrother.Next {

				if nextBrother.Left != nil {
					right.Next = nextBrother.Left
					break
				}
				if nextBrother.Right != nil {
					right.Next = nextBrother.Right
					break
				}
			}

			connect2(right.Left, right.Right, right.Next)
		} else { //左不空，右空

			//brother虽然没有子节点，但是，它的下一个可能有子节点也算！
			for nextBrother := brother; nextBrother != nil; nextBrother = nextBrother.Next {

				if nextBrother.Left != nil {
					left.Next = nextBrother.Left
					break
				}
				if nextBrother.Right != nil {
					left.Next = nextBrother.Right
					break
				}
			}
		}
		//递归处理
		connect2(left.Left, left.Right, left.Next)

	} else {              //左孩子为空
		if right != nil { //右也不空
			for nextBrother := brother; nextBrother != nil; nextBrother = nextBrother.Next {

				if nextBrother.Left != nil {
					right.Next = nextBrother.Left
					break
				}
				if nextBrother.Right != nil {
					right.Next = nextBrother.Right
					break
				}
			}
			connect2(right.Left, right.Right, right.Next)
		}
	}
}
