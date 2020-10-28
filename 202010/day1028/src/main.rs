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

impl Solution {
    //剑指 Offer 32 - I. 从上到下打印二叉树
    pub fn level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        //思路：采用队列的方式
        let mut queue = vec![];
        let mut res = vec![];
        if let Some(node) = root {
            queue.push(node);
        }
        while !queue.is_empty() {
            //从队列中弹出一个
            let node = queue.remove(0);
            res.push(node.borrow().val);
            if node.borrow().left.is_some(){
                queue.push(node.borrow_mut().left.clone().unwrap());
            }
            if node.borrow().right.is_some(){
                queue.push(node.borrow().right.clone().unwrap());
            }
        }
        res
    }
}

struct Solution {}

impl Solution {
    //每日一题：1207.独一无二的出现次数
    pub fn unique_occurrences(arr: Vec<i32>) -> bool {
        //思路；用map来记录数字出现的次数（也可以直接用数组来记录，因为数的大小是固定的）
        use std::collections::HashMap;
        use std::collections::HashSet;
        let mut m = HashMap::new();
        for i in arr.iter() {
            if m.contains_key(i) {
                m.insert(*i, m.get(i).unwrap() + 1);
            } else {
                m.insert(*i, 1);
            }
        }
        let mut s = HashSet::new();
        for i in m.iter() {
            if s.contains(i.1) {
                return false;
            } else {
                s.insert(*i.1);
            }
        }
        true
    }
}