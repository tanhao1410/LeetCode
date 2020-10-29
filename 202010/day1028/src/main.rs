fn main() {
    //println!("Hello, world!");
    //Solution::count_digit_one(16);
    let start = std::time::SystemTime::now();
    for _ in 0..100000000 {
        let mut arr = [1, 2, 3, 4, 5, 6, 7, 8, 9];//速度比vec快4倍以上
        BubbleSort(&mut arr);
    }

    let end = SystemTime::now().duration_since(start).unwrap().as_nanos();
    println!("{}", end);
}

fn BubbleSort(arr:&mut [i32;9]) {
    for i in 0..arr.len() - 1 {
        for j in 0..arr.len() - 1 - i {
            if arr[j] < arr[j + 1] {
                let temp = arr[j];
                arr[j] = arr[j +1];
                arr[j +1] = temp;
            }
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

use std::rc::Rc;
use std::cell::RefCell;
use std::collections::HashMap;
use std::time::SystemTime;
use std::mem::swap;

impl Solution {
    //剑指 Offer 32 - I. 从上到下打印二叉树
    pub fn level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        //思路：采用队列的方式
        let mut queue = vec![];
        let mut res = vec![];
        if let Some(node) = root {
            queue.push(node);
        }
        while !queue.is_empty() {
            //从队列中弹出一个
            let node = queue.remove(0);
            res.push(node.borrow().val);
            if node.borrow().left.is_some() {
                queue.push(node.borrow_mut().left.clone().unwrap());
            }
            if node.borrow().right.is_some() {
                queue.push(node.borrow().right.clone().unwrap());
            }
        }
        res
    }

    pub fn level_order2(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        // Passed 0ms 2.1mb
        let (mut root, mut res, mut arr) = (root, Vec::new(), Vec::new());
        if root.is_some() { arr.push(root); }
        while !arr.is_empty() {
            let mut children = Vec::new();
            for mut node in arr {
                let mut node = node.as_mut().unwrap().borrow_mut();
                res.push(node.val);
                if node.left.is_some() { children.push(node.left.take()); }
                if node.right.is_some() { children.push(node.right.take()); }
            }
            arr = children;
        }
        res
    }

    //剑指 Offer 07. 重建二叉树
    pub fn build_tree(preorder: Vec<i32>, inorder: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
        if preorder.is_empty() {
            return None;
        }
        let mut root = TreeNode { val: preorder[0], left: None, right: None };

        //切割inorder和preorder
        //先切中序
        let mut index = 0;
        while inorder[index] != preorder[0] {
            index += 1;
        }
        let inorder_left = inorder[..index].to_vec();
        let inorder_right = inorder[index + 1..].to_vec();

        let preorder_left = preorder[1..1 + inorder_left.len()].to_vec();
        let preorder_right = preorder[1 + inorder_left.len()..].to_vec();

        root.left = Solution::build_tree(preorder_left, inorder_left);
        root.right = Solution::build_tree(preorder_right, inorder_right);
        return Some(Rc::new(RefCell::new(root)));
    }
}

struct Solution {}

impl Solution {
    //每日一题：1207.独一无二的出现次数
    pub fn unique_occurrences(arr: Vec<i32>) -> bool {
        //思路；用map来记录数字出现的次数（也可以直接用数组来记录，因为数的大小是固定的）
        use std::collections::HashMap;
        use std::collections::HashSet;
        let mut m = HashMap::new();
        for i in arr.iter() {
            if m.contains_key(i) {
                m.insert(*i, m.get(i).unwrap() + 1);
            } else {
                m.insert(*i, 1);
            }
        }
        let mut s = HashSet::new();
        for i in m.iter() {
            if s.contains(i.1) {
                return false;
            } else {
                s.insert(*i.1);
            }
        }
        true
    }
}