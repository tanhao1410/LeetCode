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
    //606. 根据二叉树创建字符串
    pub fn tree2str(root: Option<Rc<RefCell<TreeNode>>>) -> String {
        //递归的形式
        if root.is_none() {
            return "".to_string();
        }
        //如果节点仅一个，那么
        if root.as_ref().unwrap().borrow().left.is_none()
            && root.as_ref().unwrap().borrow().right.is_none() {
            return root.as_ref().unwrap().borrow().val.to_string();
        }
        let mut res = root.as_ref().unwrap().borrow().val.to_string();

        //push左边
        res.push('(');
        res.push_str(&mut Self::tree2str(root.as_ref().unwrap().borrow_mut().left.take()));
        res.push(')');
        //push右边
        if root.as_ref().unwrap().borrow().right.is_some() {
            res.push('(');
            res.push_str(&mut Self::tree2str(root.as_ref().unwrap().borrow_mut().right.take()));
            res.push(')');
        }
        res
    }
}

struct Solution;