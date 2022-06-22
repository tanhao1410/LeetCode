fn main() {
    println!("Hello, world!");
}

impl Solution {
    //1382. 将二叉搜索树变平衡
    pub fn balance_bst(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        let mut list = vec![];
        Self::mid_read_tree(root, &mut list);
        Self::sorted_list_bst(&list)
    }

    fn mid_read_tree(root: Option<Rc<RefCell<TreeNode>>>, list: &mut Vec<Rc<RefCell<TreeNode>>>) {
        if let Some(node) = root {
            Self::mid_read_tree(node.borrow_mut().left.take(), list);
            let right = node.borrow_mut().right.take();
            list.push(node.clone());
            Self::mid_read_tree(right, list);
        }
    }

    fn sorted_list_bst(list: &[Rc<RefCell<TreeNode>>]) -> Option<Rc<RefCell<TreeNode>>> {
        let mut res = None;
        if list.len() > 0 {
            let mut root = list[list.len() / 2].clone();
            root.borrow_mut().left = Self::sorted_list_bst(&list[..list.len() / 2]);
            root.borrow_mut().right = Self::sorted_list_bst(&list[list.len() / 2 + 1..]);
            res = Some(root);
        }
        res
    }

    //1315. 祖父节点值为偶数的节点和
    pub fn sum_even_grandparent(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let sum_son = |root: Rc<RefCell<TreeNode>>| {
            let mut res = 0;
            let node_ref = root.borrow();
            res += node_ref.left.as_ref().map_or(0, |e| e.borrow().val);
            res += node_ref.right.as_ref().map_or(0, |e| e.borrow().val);
            res
        };

        let sum_grand_son = |root: Rc<RefCell<TreeNode>>| {
            let mut res = 0;
            let node_ref = root.borrow();
            if node_ref.left.is_some() {
                res += sum_son(node_ref.left.as_ref().unwrap().clone());
            }
            if node_ref.right.is_some() {
                res += sum_son(node_ref.right.as_ref().unwrap().clone());
            }
            res
        };

        let mut res = 0;
        if let Some(node) = root {
            if node.borrow().val % 2 == 0 {
                res += sum_grand_son(node.clone());
            }

            res += Self::sum_even_grandparent(node.borrow_mut().left.take());
            res += Self::sum_even_grandparent(node.borrow_mut().right.take());
        }
        res
    }

    //513. 找树左下角的值
    pub fn find_bottom_left_value(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut queue = std::collections::VecDeque::new();
        queue.push_back(root.unwrap());
        let mut res = 0;
        while !queue.is_empty() {
            let queue_len = queue.len();
            res = queue.get(0).unwrap().borrow().val;
            for _ in 0..queue_len {
                let pop = queue.pop_front().unwrap();
                if pop.borrow().left.is_some() {
                    queue.push_back(pop.borrow().left.as_ref().unwrap().clone());
                }
                if pop.borrow().right.is_some() {
                    queue.push_back(pop.borrow().right.as_ref().unwrap().clone());
                }
            }
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