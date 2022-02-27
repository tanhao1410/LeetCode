fn main() {
    println!("Hello, world!");
}

impl Solution {
    //1367. 二叉树中的列表
    pub fn is_sub_path(mut head: Option<Box<ListNode>>, root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        if let Some(node) = head {
            if let Some(tree_node) = root {
                let mut ref_mut = tree_node.borrow_mut();
                if node.val == tree_node.borrow().val {
                    if Self::is_equal(&node.next, &ref_mut.left)
                        || Self::is_equal(&node.next, &ref_mut.right) {
                        return true;
                    }
                }
                Self::is_sub_path((Some(node.clone())), ref_mut.left.take())
                    || Self::is_sub_path(Some(node), ref_mut.right.take())
            } else {
                false
            }
        } else {
            true
        }
    }
    //必须是接连相等
    fn is_equal(mut head: &Option<Box<ListNode>>, root: &Option<Rc<RefCell<TreeNode>>>) -> bool {
        if let Some(node) = head {
            if let Some(tree_node) = root {
                if node.val != tree_node.borrow().val {
                    return false;
                }
                Self::is_equal(&node.next, &tree_node.borrow().left)
                    || Self::is_equal(&node.next, &tree_node.borrow().right)
            } else {
                false
            }
        } else {
            true
        }
    }
}

#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode {
            next: None,
            val,
        }
    }
}
Definition for a binary tree node.
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
