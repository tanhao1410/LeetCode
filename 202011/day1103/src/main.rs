fn main() {
    println!("Hello, world!");
    println!("{}", Solution::valid_mountain_array(vec![0, 3, 2, 1]))
}

struct Solution {}

impl Solution {
    //每日一题：941. 有效的山脉数组
    pub fn valid_mountain_array(a: Vec<i32>) -> bool {
        if a.len() < 3 || a[1] <= a[0] {
            return false;
        }
        let mut pre = a[0];
        //先递增，后递减
        let mut is_increase = true;
        for i in 1..a.len() {
            if is_increase && a[i] > pre {} else if a[i] < pre {
                is_increase = false;
            } else {
                return false;
            }
            pre = a[i];
        }
        !is_increase
    }

    //剑指 Offer 32 - II. 从上到下打印二叉树 II
    pub fn level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
        let mut res = Vec::new();
        if root.is_none() {
            return res;
        }
        let mut queue = vec![root];
        let mut queue_len = queue.len();
        while queue_len > 0 {
            let mut res_item = vec![];
            for i in 0..queue_len {
                let v = queue.remove(0).unwrap();
                res_item.push(v.borrow_mut().val);
                if v.borrow_mut().left.is_some() {
                    queue.push(v.borrow_mut().left.clone());
                }
                if v.borrow_mut().right.is_some() {
                    queue.push(v.borrow_mut().right.clone());
                }
            }
            res.push(res_item);
            queue_len = queue.len();
        }
        res
    }

    //剑指 Offer 40. 最小的k个数
    pub fn get_least_numbers(arr: Vec<i32>, k: i32) -> Vec<i32> {
        if k == arr.len() as i32 {
            return arr.clone();
        }
        let mut res = vec![0; k as usize];
        for i in 0..k as usize {
            res[i] = arr[i];
        }
        res.sort();
        for i in k..arr.len() as i32 {
            let mut j = k - 1;
            if arr[i as usize] >= res[j as usize] {
                continue;
            }
            while j >= 0 && res[j as usize] > arr[i as usize] {
                if j != 0 && res[j as usize - 1] > arr[i as usize] {
                    res[j as usize] = res[j as usize - 1];
                } else {
                    res[j as usize] = arr[i as usize]
                }
                j -= 1;
            }
        }
        res
    }

    //剑指 Offer 38. 字符串的排列
    pub fn permutation(s: String) -> Vec<String> {
        let mut res = Vec::new();
        //思路：去重，然后排列组合即可
        fn create_str(mut pre: String, mut remain: Vec<char>, mut res: &mut Vec<String>) {
            if remain.len() == 0 {
                res.push(pre);
            } else {
                let mut pre_char = '&';
                for i in 0..remain.len() {
                    if pre_char == remain[i] {
                        continue;
                    }
                    let mut next = pre.clone();
                    let mut new_remain = remain.clone();
                    let v = new_remain.remove(i);
                    next.push(v);
                    create_str(next, new_remain, &mut res);
                    pre_char = v;
                }
            }
        }
        let mut remain = vec![];
        for i in s.chars() {
            remain.push(i);
        }
        remain.sort();
        create_str("".to_string(), remain, &mut res);
        res
    }

    //剑指 Offer 32 - III. 从上到下打印二叉树 III
    pub fn level_order2(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
        let mut res = Vec::new();
        if let None = root {
            return res;
        }
        let mut queue = vec![root.unwrap()];
        let mut queue_len = queue.len();
        let mut is_odd = true;
        while queue_len > 0 {
            let mut item = Vec::new();
            for i in 0..queue_len {
                let node = queue.remove(0);
                if is_odd {
                    item.push(node.borrow().val);
                } else {
                    item.insert(0, node.borrow().val);
                }
                if node.borrow().left.is_some() {
                    queue.push(node.borrow().left.clone().unwrap());
                }
                if node.borrow().right.is_some() {
                    queue.push(node.borrow().right.clone().unwrap());
                }
            }
            is_odd = !is_odd;
            res.push(item);
            queue_len = queue.len();
        }
        res
    }

    //剑指 Offer 33. 二叉搜索树的后序遍历序列
    pub fn verify_postorder(postorder: Vec<i32>) -> bool {
        //判断是否是某二叉搜索树的后序遍历，即左右中
        if postorder.len() < 3 {
            return true;
        }
        //根节点
        let last_node = *postorder.last().unwrap();
        //以根节点为界，分隔原来的序列
        let mut index = 0;
        while postorder[index] < last_node {
            index += 1;
        }

        //判断右边是否有小于该数的
        let mut index2 = index;
        while index2 < postorder.len() - 1 {
            index2 += 1;
            if postorder[index2] < last_node {
                return false;
            }
        }

        if index < postorder.len() - 1 {
            Solution::verify_postorder(postorder[..index].to_vec())
                && Solution::verify_postorder(postorder[index..postorder.len() - 1].to_vec())
        } else {
            Solution::verify_postorder(postorder[..index].to_vec())
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
use std::borrow::Borrow;