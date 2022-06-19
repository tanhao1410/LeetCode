fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //508. 出现次数最多的子树元素和
    pub fn find_frequent_tree_sum(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut map = HashMap::new();
        Self::tree_sum(&root, &mut map);
        let max = map.iter().map(|e| *e.1).max().unwrap();
        map.into_iter()
            .filter(|e| e.1 == max)
            .map(|e| e.0)
            .collect()
    }

    fn tree_sum(root: &Option<Rc<RefCell<TreeNode>>>, map: &mut HashMap<i32, i32>) -> i32 {
        let mut sum = 0;
        if let Some(node) = root {
            let left_sum = Self::tree_sum(&node.borrow().left, map);
            let right_sum = Self::tree_sum(&node.borrow().right, map);
            sum = node.borrow().val + left_sum + right_sum;
            *map.entry(sum).or_insert(0) += 1;
        }
        sum
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
use std::collections::HashMap;