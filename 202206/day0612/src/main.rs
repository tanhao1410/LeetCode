fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //1008. 前序遍历构造二叉搜索树
    pub fn bst_from_preorder<T: AsRef<[i32]>>(preorder: T) -> Option<Rc<RefCell<TreeNode>>> {
        let preorder = preorder.as_ref();
        if let Some(v) = preorder.get(0) {
            let mut split_index = preorder.len();
            for i in 0..split_index {
                if preorder[i] > *v {
                    split_index = i;
                    break;
                }
            }
            let mut root = TreeNode::new(*v);
            root.left = Self::bst_from_preorder(&preorder[1..split_index]);
            root.right = Self::bst_from_preorder(&preorder[split_index..]);
            return Some(Rc::new(RefCell::new(root)));
        }
        None
    }

    // 890.查找和替换模式
    pub fn find_and_replace_pattern(words: Vec<String>, pattern: String) -> Vec<String> {
        let pattern_bytes = pattern.as_bytes();
        //长度一致，对应关系一致
        let match_pattern = |word: &String| {
            let mut match_vec = vec![0u8; 128];
            let mut match_rev = vec![0u8; 128];
            let word_bytes = word.as_bytes();
            word_bytes.iter().zip(pattern_bytes.iter()).all(|(&i, &j)| {
                let res = match_vec[i as usize] == 0 || match_vec[i as usize] == j;
                let res2 = match_rev[j as usize] == 0 || match_rev[j as usize] == i;
                match_vec[i as usize] = j;
                match_rev[j as usize] = i;
                res && res2
            })
        };

        words.into_iter()
            .filter(match_pattern)
            .collect()
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
use std::ops::Deref;