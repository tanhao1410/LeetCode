fn main() {
    println!("Hello, world!");
}

//523. 连续的子数组和
pub fn check_subarray_sum(nums: Vec<i32>, k: i32) -> bool {
    //思路：前缀和 + %k 余数 相等 余数相等，说明 相减 即为 整除
    //思路：前缀和 + %k 余数 相等 余数相等，说明 相减 即为 整除
    let mut m = std::collections::HashMap::new();
    //先求前缀和
    let mut sums = nums;
    for i in 0..sums.len() {
        if i > 0{
            sums[i] += sums[i - 1];
        }
        let mo = sums[i] % k;
        if mo == 0{
            return true;
        }
        if m.contains_key(&mo){
            if i - m.get(&mo).unwrap() > 1{
                return true;
            }
        }else{
            m.insert(mo,i);
        }
    }
    false
}

use std::rc::Rc;
use std::cell::RefCell;

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

    //606. 根据二叉树创建字符串
    pub fn tree2str(root: Option<Rc<RefCell<TreeNode>>>) -> String {
        let res = Self::tree2str2(root.as_ref());
        res[1..res.len() - 1].to_string()
    }

    pub fn tree2str2(root: Option<&Rc<RefCell<TreeNode>>>) -> String {
        if let Some(node) = root {
            match (node.borrow().left.as_ref(), node.borrow().right.as_ref()) {
                (None, None) => String::from("(") + &node.borrow().val.to_string() + ")",
                (_, None) => "(".to_string() + &node.borrow().val.to_string() + &Self::tree2str2(node.borrow().left.as_ref()) + ")",
                _ => "(".to_string() + &node.borrow().val.to_string()
                    + &Self::tree2str2(node.borrow().left.as_ref())
                    + &Self::tree2str2(node.borrow().right.as_ref())
                    + ")"
            }
        } else {
            "()".to_string()
        }
    }

    //面试题 17.12. BiNode
    pub fn convert_bi_node(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        //把二叉搜索树转换为单向链表，要求依然符合二叉搜索树的性质
        //二叉搜索树的先序遍历是有序的。
        //思路：有左节点的，应该把左节点当成父，父当成左的右子树

        let mut arr: Vec<Rc<RefCell<TreeNode>>> = vec![];
        Self::middle_read(root, &mut arr);
        //构建新的
        if arr.is_empty() {
            return None;
        }
        let mut res = arr.remove(0);
        let mut p = res.clone();
        while !arr.is_empty() {
            let next = arr.remove(0);
            p.borrow_mut().right = Some(next.clone());
            p = next;
        }
        Some(res)
    }

    pub fn middle_read(root: Option<Rc<RefCell<TreeNode>>>, arr: &mut Vec<Rc<RefCell<TreeNode>>>) {
        if let Some(node) = root {
            if node.borrow().left.is_some() {
                Self::middle_read(node.borrow_mut().left.take(), arr);
            }
            let right = node.borrow_mut().right.take();
            arr.push(node);
            Self::middle_read(right, arr);
        }
    }
}