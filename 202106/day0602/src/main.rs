fn main() {
    println!("Hello, world!");
}

use std::rc::Rc;
use std::cell::RefCell;

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

    //面试题 17.12. BiNode
    pub fn convert_bi_node(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        //把二叉搜索树转换为单向链表，要求依然符合二叉搜索树的性质
        //二叉搜索树的先序遍历是有序的。
        //思路：有左节点的，应该把左节点当成父，父当成左的右子树

        let mut arr: Vec<Rc<RefCell<TreeNode>>> = vec![];
        Self::middle_read(root, &mut arr);
        //构建新的
        if arr.is_empty() {
            return None;
        }
        let mut res = arr.remove(0);
        let mut p = res.clone();
        while !arr.is_empty() {
            let next = arr.remove(0);
            p.borrow_mut().right = Some(next.clone());
            p = next;
        }
        Some(res)
    }

    pub fn middle_read(root: Option<Rc<RefCell<TreeNode>>>, arr: &mut Vec<Rc<RefCell<TreeNode>>>) {
        if let Some(node) = root {
            if node.borrow().left.is_some() {
                Self::middle_read(node.borrow_mut().left.take(), arr);
            }
            let right = node.borrow_mut().right.take();
            arr.push(node);
            Self::middle_read(right, arr);
        }
    }
}