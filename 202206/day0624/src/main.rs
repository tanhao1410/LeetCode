fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //剑指 Offer II 046. 二叉树的右侧视图
    pub fn right_side_view(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        use std::collections::VecDeque;
        let mut res = vec![];
        if let Some(node) = root {
            let mut queue = VecDeque::new();
            queue.push_back(node);
            while !queue.is_empty() {
                let queue_len = queue.len();
                let mut item = 0;
                for _ in 0..queue_len {
                    let head = queue.pop_front().unwrap();
                    let head_ref = head.borrow();
                    head_ref.left.as_ref().map_or((), |e| queue.push_back(e.clone()));
                    head_ref.right.as_ref().map_or((), |e| queue.push_back(e.clone()));
                    item = head_ref.val;
                }
                res.push(item);
            }
        }
        res
    }

    //2265. 统计值等于子树平均值的节点数
    pub fn average_of_subtree(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        Self::sum_count_res(root.unwrap()).2
    }

    //返回总和，节点数，答案
    fn sum_count_res(root: Rc<RefCell<TreeNode>>) -> (i32, i32, i32) {
        let root = root.borrow();
        let left_res = root.left.as_ref().map_or((0, 0, 0), |e| Self::sum_count_res(e.clone()));
        let right_res = root.right.as_ref().map_or((0, 0, 0), |e| Self::sum_count_res(e.clone()));

        let mut res = (root.val + left_res.0 + right_res.0,
                       1 + left_res.1 + right_res.1,
                       left_res.2 + right_res.2);
        if res.0 / res.1 == root.val {
            res.2 += 1;
        }
        res
    }

    //1026. 节点与其祖先之间的最大差值
    pub fn max_ancestor_diff(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        Self::max_min_maxdiff(root.as_ref().unwrap().clone()).2
    }

    fn max_min_maxdiff(root: Rc<RefCell<TreeNode>>) -> (i32, i32, i32) {
        let root = root.borrow();
        let mut res = (root.val, root.val, 0);
        let process_inner = |e: Rc<RefCell<TreeNode>>, res: &mut (i32, i32, i32)| {
            let inner_res = Self::max_min_maxdiff(e);
            res.0 = res.0.max(inner_res.0);
            res.1 = res.1.min(inner_res.1);
            res.2 = res.2.max(inner_res.2).max((root.val - res.0).abs()).max((root.val - res.1).abs());
        };
        root.left.as_ref().map_or((), |e| process_inner(e.clone(), &mut res));
        root.right.as_ref().map_or((), |e| process_inner(e.clone(), &mut res));
        res
    }
    //515. 在每个树行中找最大值
    pub fn largest_values(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        use std::collections::VecDeque;
        let mut queue = VecDeque::new();
        let mut res = vec![];
        if let Some(node) = root {
            queue.push_back(node);
            while !queue.is_empty() {
                let queue_len = queue.len();
                let mut max = i32::MIN;
                for _ in 0..queue_len {
                    let head = queue.pop_front().unwrap();
                    let head_ref = head.borrow();
                    max = max.max(head_ref.val);
                    head_ref.left.as_ref().map_or((), |e| queue.push_back(e.clone()));
                    head_ref.right.as_ref().map_or((), |e| queue.push_back(e.clone()));
                }
                res.push(max);
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