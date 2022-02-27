fn main() {
    println!("Hello, world!");
}

impl Solution {
    //3. 无重复字符的最长子串
    pub fn length_of_longest_substring(s: String) -> i32 {
        // 用一个map 记录每一个字母再前面的位置，-1代表无。
        //子串最大为26，用一个滑动窗口，开始时是1，然后依次加入新的字母，看该字母在前面是否存在，如果存在，看是否在窗口中
        //如果不在窗口里，则继续。否则，窗口开始位置定为它的下一个位置
        let mut pre_index = vec![-1; 128];
        let bytes = s.as_bytes();
        let mut start = 0;
        let mut res = 0;
        for i in 0..s.len() {
            //在窗口中
            if pre_index[bytes[i] as usize] != -1 && pre_index[bytes[i] as usize] >= start as i32 {
                start = pre_index[bytes[i] as usize] as usize + 1;
            }
            pre_index[bytes[i] as usize] = i as i32;
            res = res.max(i - start + 1);
        }
        res as i32
    }
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
