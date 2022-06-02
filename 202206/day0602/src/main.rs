fn main() {
    println!("Hello, world!");
}

impl Solution {
    //450. 删除二叉搜索树中的节点
    pub fn delete_node(root: Option<Rc<RefCell<TreeNode>>>, key: i32) -> Option<Rc<RefCell<TreeNode>>> {
        let default = Rc::new(RefCell::new(TreeNode::new(0)));
        default.borrow_mut().left = root;
        Self::delete_node_parent(&default.borrow().left, key, default.clone(), true);
        let res = default.borrow_mut().left.take();
        res
    }

    fn delete_node_parent(root: &Option<Rc<RefCell<TreeNode>>>, key: i32, parent: Rc<RefCell<TreeNode>>, is_left: bool) {
        if let Some(node) = root {
            if node.borrow().val > key {
                Self::delete_node_parent(&node.borrow().left, key, node.clone(), true);
            } else if node.borrow().val < key {
                Self::delete_node_parent(&node.borrow().right, key, node.clone(), false);
            } else {
                //待删除的节点是否有右节点
                let mut ref_mut = node.borrow_mut();
                let left = ref_mut.left.take();
                let right = ref_mut.right.take();
                let parent = parent.as_ptr() as *mut TreeNode;
                if let Some(right) = right {
                    let min_node = Self::min_node(right.clone());
                    min_node.borrow_mut().left = left;
                    unsafe {
                        //let x = parent.borrow_mut() as *mut TreeNode;
                        if is_left {
                            (*parent).left = Some(right);
                        } else {
                            (*parent).right = Some(right);
                        }
                    }
                } else {
                    //待删除的节点没有右节点
                    unsafe {
                        if is_left {
                            (*parent).left = left;
                        } else {
                            (*parent).right = left;
                        }
                    }
                }
            }
        }
    }

    fn min_node(root: Rc<RefCell<TreeNode>>) -> Rc<RefCell<TreeNode>> {
        if let Some(next) = &root.borrow().left {
            return Self::min_node(next.clone());
        }
        root
    }
}

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

struct Solution;

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