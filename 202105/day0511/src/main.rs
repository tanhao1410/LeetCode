use std::rc::Rc;
use std::cell::RefCell;

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

    //563. 二叉树的坡度
    pub fn find_tilt(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        Self::child_sum_tilt(&root).1
    }
    //子树包含自身之和，坡度
    pub fn child_sum_tilt(root:&Option<Rc<RefCell<TreeNode>>>)->(i32,i32){
        if let Some(node) = root.as_ref(){
            //先求它的左子树，和右子树
            let (left_sum,left_tilt) = Self::child_sum_tilt(&node.borrow().left);
            let (right_sum,right_tilt) = Self::child_sum_tilt(&node.borrow().right);
            return (left_sum + right_sum + node.borrow().val ,left_tilt+right_tilt + (left_sum-right_sum).abs());
        }
        (0,0)
    }

    //671. 二叉树中第二小的节点
    pub fn find_second_minimum_value(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        if let Some(node) = root {
            let num = node.borrow().val;
            return Self::find_second_minimum_value2(num, Some(node));
        }
        -1
    }

    //671. 二叉树中第二小的节点,返回最小的值，但大于某数
    pub fn find_second_minimum_value2(num: i32, root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        if let Some(node) = root {
            //没有子节点
            if node.borrow().left.is_none() && node.borrow().right.is_none() {
                if node.borrow().val > num {
                    return node.borrow().val;
                }
                return -1;
            }
            let left = Self::find_second_minimum_value2(num, node.borrow_mut().left.take());
            let right = Self::find_second_minimum_value2(num, node.borrow_mut().right.take());
            if left == -1 || (right != -1 && right < left) {
                return right;
            }
            if right == -1 || (left != -1 && left < right) {
                return left;
            }
            //两者相等时，且不等于-1
            return left;
        }
        -1
    }
}