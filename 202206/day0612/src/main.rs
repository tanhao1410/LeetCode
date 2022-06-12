fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //1145. 二叉树着色游戏
    pub fn btree_game_winning_move(root: Option<Rc<RefCell<TreeNode>>>, n: i32, x: i32) -> bool {
        //1.选择x的父节点，则除了x及其子节点外，其余的都是y，若大于，则可以赢
        let root_ref = root.as_ref().unwrap().borrow();
        if root_ref.val != x {
            //不是跟节点
            let x_num = Self::btee_node_num(&root, x);
            //可以选择x节点的父节点。
            if x_num * 2 < n as usize {
                return true;
            }
        }
        //2.选x的子节点，则该子节点总和 大于剩余的其它的所有 可以赢。
        let child = Self::btree_child(&root, x);
        if let Some(left) = child.0 {
            let left_num = Self::btee_node_num(&root, left);
            println!("{}", left_num);
            if left_num * 2 > n as usize {
                return true;
            }
        }

        if let Some(right) = child.1 {
            let right = Self::btee_node_num(&root, right);
            if right * 2 > n as usize {
                return true;
            }
        }

        //3.其余的地方都不应选择，或者说都不是最优选择。
        false
    }

    fn btree_child(root: &Option<Rc<RefCell<TreeNode>>>, n: i32) -> (Option<i32>, Option<i32>) {
        let mut res = (None, None);
        if let Some(root) = root.as_ref() {
            let root_ref = root.borrow();
            if root_ref.val == n {
                if root_ref.left.is_some() {
                    res.0 = Some(root_ref.left.as_ref().unwrap().borrow().val);
                }
                if root_ref.right.is_some() {
                    res.1 = Some(root_ref.right.as_ref().unwrap().borrow().val);
                }
                return res;
            }
            let left = Self::btree_child(&root_ref.left, n);
            let right = Self::btree_child(&root_ref.right, n);
            if left.0.is_some() || left.1.is_some() {
                return left;
            }
            return right;
        }
        res
    }

    fn btee_node_num(root: &Option<Rc<RefCell<TreeNode>>>, n: i32) -> usize {
        let mut res = 0;
        if let Some(root) = root {
            let root_ref = root.borrow();
            if root_ref.val == n {
                return Self::btree_num(root.clone());
            }
            let left = Self::btee_node_num(&root_ref.left, n);
            let right = Self::btee_node_num(&root_ref.right, n);
            res = left.max(right);
        }
        res
    }

    fn btree_num(root: Rc<RefCell<TreeNode>>) -> usize {
        let mut res = 1;
        if root.borrow().left.is_some() {
            res += Self::btree_num(root.borrow().left.as_ref().unwrap().clone());
        }
        if root.borrow().right.is_some() {
            res += Self::btree_num(root.borrow().right.as_ref().unwrap().clone());
        }
        res
    }

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