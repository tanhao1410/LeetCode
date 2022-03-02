fn main() {
    println!("Hello, world!");
}


impl Solution {
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