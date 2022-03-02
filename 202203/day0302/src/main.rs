fn main() {
    println!("Hello, world!");
}


impl Solution {
    //450. 删除二叉搜索树中的节点
    pub fn delete_node(mut root: Option<Rc<RefCell<TreeNode>>>, key: i32) -> Option<Rc<RefCell<TreeNode>>> {
        if root.is_none() {
            return None;
        }
        if root.as_ref().unwrap().borrow().val < key {
            let right = root.as_mut().unwrap().borrow_mut().right.take();
            root.as_mut().unwrap().borrow_mut().right = Self::delete_node(right, key);
            return root;
        } else if root.as_ref().unwrap().borrow().val > key {
            let left = root.as_mut().unwrap().borrow_mut().left.take();
            root.as_mut().unwrap().borrow_mut().left = Self::delete_node(left, key);
            return root;
        } else {
            //当前节点就是待删除的节点
            if root.as_ref().unwrap().borrow().left.is_none() {
                return root.as_mut().unwrap().borrow_mut().right.take();
            } else if root.as_ref().unwrap().borrow().right.is_none() {
                return root.as_mut().unwrap().borrow_mut().left.take();
            } else {
                let mut left = root.as_mut().unwrap().borrow_mut().left.take();
                let mut right = root.as_mut().unwrap().borrow_mut().right.take();
                let mut node = right.as_ref().unwrap().as_ptr();
                unsafe {
                    //得到它的最左边的节点
                    while (*node).left.is_some() {
                        node = (*node).left.as_ref().unwrap().as_ptr();
                    }
                    (*node).left = left;
                }
                return right;
            }
        }
    }

    //329. 矩阵中的最长递增路径
    pub fn longest_increasing_path(matrix: Vec<Vec<i32>>) -> i32 {
        //深度优先遍历，把所有的位置都放进去，以弹出的点开始向四周走，如果能走动，则把心加入的放入栈中
        //直到栈空。
        //能走动的意思是，对方比自己大，且走动后，的对方的距离 增加了。
        let mut stack = vec![];
        let mut path_len = vec![vec![1; matrix[0].len()]; matrix.len()];
        for i in 0..matrix.len() {
            for j in 0..matrix[0].len() {
                stack.push((i, j));
            }
        }
        let dircts = vec![vec![0, 1], vec![1, 0], vec![0, -1], vec![-1, 0]];
        while let Some((x, y)) = stack.pop() {
            //冲个方向开始动
            for dirct in &dircts {
                if x as i32 + dirct[0] >= 0 && x as i32 + dirct[0] < matrix.len() as i32
                    && y as i32 + dirct[1] >= 0 && y as i32 + dirct[1] < matrix[0].len() as i32
                    && matrix[(x as i32 + dirct[0]) as usize][(y as i32 + dirct[1]) as usize] > matrix[x][y]
                    && path_len[(x as i32 + dirct[0]) as usize][(y as i32 + dirct[1]) as usize] < path_len[x][y] + 1 {
                    path_len[(x as i32 + dirct[0]) as usize][(y as i32 + dirct[1]) as usize] = path_len[x][y] + 1;
                    stack.push(((x as i32 + dirct[0]) as usize, (y as i32 + dirct[1]) as usize));
                }
            }
        }
        path_len
            .into_iter()
            .flat_map(|v| v.into_iter())
            .max()
            .unwrap()
    }
    //113. 路径总和 II
    pub fn path_sum(root: Option<Rc<RefCell<TreeNode>>>, target_sum: i32) -> Vec<Vec<i32>> {
        if root.is_none() {
            return vec![];
        }
        Self::path_sum_re(&root, target_sum, vec![])
    }
    fn path_sum_re(root: &Option<Rc<RefCell<TreeNode>>>, target: i32, mut pre: Vec<i32>) -> Vec<Vec<i32>> {
        let mut res = vec![];
        let root = root.as_ref().unwrap().borrow();
        pre.push(root.val);

        if root.left.is_none() && root.right.is_none() {
            if target == root.val {
                res.push(pre);
            }
        } else {
            if root.left.is_some() {
                let pre2 = pre.clone();
                res.append(&mut Self::path_sum_re(&root.left, target - root.val, pre2));
            }
            if root.right.is_some() {
                res.append(&mut Self::path_sum_re(&root.right, target - root.val, pre));
            }
        }
        res
    }

    //199. 二叉树的右视图
    pub fn right_side_view(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        use std::collections::VecDeque;
        let mut res = vec![];
        let mut queue = VecDeque::new();
        if root.is_some() {
            queue.push_back(root.as_ref().unwrap().clone());
        }
        while !queue.is_empty() {
            let len = queue.len();
            for i in 0..len {
                let cur = queue.pop_front().unwrap();
                if i == len - 1 {
                    res.push(cur.borrow().val);
                }
                if cur.borrow().left.is_some() {
                    queue.push_back(cur.borrow().left.as_ref().unwrap().clone());
                }
                if cur.borrow().right.is_some() {
                    queue.push_back(cur.borrow().right.as_ref().unwrap().clone());
                }
            }
        }
        res
    }
}

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

struct Solution;

use std::rc::Rc;
use std::cell::RefCell;