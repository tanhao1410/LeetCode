#![feature(destructuring_assignment)]

fn main() {
    println!("Hello, world!");
}


impl Solution {
    //696. 计数二进制子串
    pub fn count_binary_substrings(s: String) -> i32 {
        let mut res = 0;
        let (mut one, mut pre_one, mut zero, mut pre_zero) = (0, 0, 0, 0);
        for i in s.bytes() {
            if i == b'0' {
                (zero, pre_zero, one) = (zero + 1, zero + 1, 0);
                if zero <= pre_one {
                    res += 1;
                }
            } else {
                (one, pre_one, zero) = (one + 1, one + 1, 0);
                if one <= pre_zero {
                    res += 1;
                }
            }
        }
        res
    }
    //682. 棒球比赛
    pub fn cal_points(ops: Vec<String>) -> i32 {
        let mut scores: Vec<i32> = vec![];
        for i in ops {
            match i.as_str() {
                "+" => scores.push(scores[scores.len() - 1] + scores[scores.len() - 2]),
                "D" => scores.push(scores[scores.len() - 1] * 2),
                "C" => {
                    scores.remove(scores.len() - 1);
                }
                i => scores.push(i.parse::<i32>().unwrap())
            }
        }
        scores.into_iter().sum()
    }

    //655. 输出二叉树
    pub fn print_tree(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<String>> {
        Self::print_tree_height(&root, 0)
    }

    fn print_tree_height(root: &Option<Rc<RefCell<TreeNode>>>, height: usize) -> Vec<Vec<String>> {
        let m = height.max(Self::tree_height(root));
        if m == 0 {
            return vec![];
        }
        let n = 2usize.pow(m as u32) - 1;
        let mut res = vec![vec!["".to_string(); n]];
        if root.is_some() {
            res[0][n / 2] = root.as_ref().unwrap().borrow().val.to_string();
        }
        let default_tree = Rc::new(RefCell::new(TreeNode::new(0)));
        let left_res = Self::print_tree_height(&root.as_ref().unwrap_or(&default_tree).borrow().left, m - 1);
        let right_res = Self::print_tree_height(&root.as_ref().unwrap_or(&default_tree).borrow().right, m - 1);
        for i in 1..m {
            let mut row = left_res[i - 1].clone();
            row.push("".to_string());
            row.append(&mut right_res[i - 1].clone());
            res.push(row);
        }
        res
    }


    fn tree_height(root: &Option<Rc<RefCell<TreeNode>>>) -> usize {
        let mut res = 0;
        if let Some(node) = root {
            res += Self::tree_height(&node.borrow().left).max(
                Self::tree_height(&node.borrow().right)
            ) + 1
        }
        res
    }
    //1022. 从根到叶的二进制数之和
    pub fn sum_root_to_leaf(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        Self::sum_root_to_leaf_pre(root.as_ref().unwrap(), 0)
    }
    fn sum_root_to_leaf_pre(root: &Rc<RefCell<TreeNode>>, pre_value: i32) -> i32 {
        let root = root.borrow();
        let pre_value = pre_value * 2 + root.val;
        if root.left.is_none() && root.right.is_none() {
            return pre_value;
        }
        let mut res = 0;
        if root.left.is_some() {
            res += Self::sum_root_to_leaf_pre(root.left.as_ref().unwrap(), pre_value);
        }
        if root.right.is_some() {
            res += Self::sum_root_to_leaf_pre(root.right.as_ref().unwrap(), pre_value)
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

use std::rc::Rc;
use std::cell::RefCell;

struct Solution;