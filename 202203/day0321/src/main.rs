fn main() {
    println!("Hello, world!");
}

impl Solution {
    //653. 两数之和 IV - 输入 BST
    pub fn find_target(root: Option<Rc<RefCell<TreeNode>>>, k: i32) -> bool {
        let mut v = vec![];
        Self::read_tree(&root, &mut v);
        let mut l = 0;
        let mut r = v.len() - 1;
        while l < r {
            if v[l] + v[r] == k {
                return true;
            } else if v[l] + v[r] < k {
                l += 1;
            } else {
                r -= 1;
            }
        }
        false
    }

    fn read_tree(root: &Option<Rc<RefCell<TreeNode>>>, vec: &mut Vec<i32>) {
        if let Some(node) = root {
            //左中右
            Self::read_tree(&node.borrow().left, vec);
            vec.push(node.borrow().val);
            Self::read_tree(&node.borrow().right, vec);
        }
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
