fn main() {
    println!("Hello, world!");
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

use std::rc::Rc;
use std::cell::RefCell;

struct Solution;

impl Solution {
    //96. 不同的二叉搜索树
    pub fn num_trees(n: i32) -> i32 {
        //思路：以 某一个节点为根节点。则有 比它大的 一部分组成树，比它小的组成树，两者相乘，即总数
        // n = 3 时，1 为根， dp[2] * dp[0],2为根，dp[1] * dp[1] 3为根，dp[2] * dp[1]
        //dp[i]即，n为几时 数量
        let mut dp = vec![0; n as usize + 1];
        dp[0] = 1;
        for i in 1..n as usize + 1 {
            for j in 0..i {
                dp[i] += dp[j] * dp[i - j - 1];
            }
        }
        dp[n as usize]
    }

    //617. 合并二叉树
    pub fn merge_trees(root1: Option<Rc<RefCell<TreeNode>>>, root2: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        if root1.is_none() {
            return root2;
        }
        if root2.is_none() {
            return root1;
        }
        {
            let mut root1_mut = root1.as_ref().unwrap().borrow_mut();
            let mut root2_mut = root2.as_ref().unwrap().borrow_mut();
            root1_mut.val = root1_mut.val + root2_mut.val;
            let root1_left = root1_mut.left.take();
            let root1_right = root1_mut.right.take();

            let root2_left = root2_mut.left.take();
            let root2_right = root2_mut.right.take();
            root1_mut.left = Solution::merge_trees(root1_left, root2_left);
            root1_mut.right = Solution::merge_trees(root1_right, root2_right);
        }

        root1
    }
}