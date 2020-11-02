fn main() {
    println!("Hello, world!");
    println!("{}", Solution::valid_mountain_array(vec![0, 3, 2, 1]))
}

struct Solution {}

impl Solution {
    //每日一题：941. 有效的山脉数组
    pub fn valid_mountain_array(a: Vec<i32>) -> bool {
        if a.len() < 3 || a[1] <= a[0] {
            return false;
        }
        let mut pre = a[0];
        //先递增，后递减
        let mut is_increase = true;
        for i in 1..a.len() {
            if is_increase && a[i] > pre {} else if a[i] < pre {
                is_increase = false;
            } else {
                return false;
            }
            pre = a[i];
        }
        !is_increase
    }

    //剑指 Offer 32 - II. 从上到下打印二叉树 II
    pub fn level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
        let mut res = Vec::new();
        if root.is_none() {
            return res;
        }
        let mut queue = vec![root];
        let mut queue_len = queue.len();
        while queue_len > 0 {
            let mut res_item = vec![];
            for i in 0..queue_len {
                let v = queue.remove(0).unwrap();
                res_item.push(v.borrow_mut().val);
                if v.borrow_mut().left.is_some() {
                    queue.push(v.borrow_mut().left.clone());
                }
                if v.borrow_mut().right.is_some() {
                    queue.push(v.borrow_mut().right.clone());
                }
            }
            res.push(res_item);
            queue_len = queue.len();
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
use std::borrow::Borrow;