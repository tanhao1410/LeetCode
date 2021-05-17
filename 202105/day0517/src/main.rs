fn main() {
    println!("Hello, world!");
}
use std::rc::Rc;
use std::cell::RefCell;


#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}
pub struct Solution {}

impl Solution {

    //993. 二叉树的堂兄弟节点
    pub fn is_cousins(root: Option<Rc<RefCell<TreeNode>>>, x: i32, y: i32) -> bool {
        //层次遍历，
        let mut queue = vec![root];
        while !queue.is_empty() {
            let old_len = queue.len();
            for i in 0..old_len {
                if let Some(node) = queue.remove(0) {
                    let mut flag = false;
                    if node.borrow().left.is_some() {
                        if node.borrow().left.as_ref().unwrap().borrow().val == x
                            || node.borrow().left.as_ref().unwrap().borrow().val == y {
                            flag = true;
                        }
                        queue.push(node.borrow_mut().left.take());
                    }
                    if node.borrow().right.is_some() {
                        //处在同一个父节点之下
                        if flag && (node.borrow().right.as_ref().unwrap().borrow().val == x
                            || node.borrow().right.as_ref().unwrap().borrow().val == y) {
                            return false;
                        }
                        queue.push(node.borrow_mut().right.take());
                    }
                }
            }

            let count = queue.iter().fold(0, |mut p, q| {
                if let Some(node) = q {
                    if node.borrow().val == x || node.borrow().val == y {
                        p += 1;
                    }
                }
                p
            });
            //判断x,y 是否是在同一层中
            if count == 2 {
                return true;
            } else if count == 1 {
                return false;
            }
        }

        false
    }


    //405. 数字转换为十六进制数
    pub fn to_hex(num: i32) -> String {
        if num == 0 {
            return "0".to_string();
        }

        let mut bi_vec = vec![false; 32];
        let mut maske = 1;
        for i in 0..32 {
            bi_vec[31 - i] = num & maske != 0;
            maske <<= 1;
        }

        let mut ox_vecc = vec![];
        for i in 0..8 {
            let mut n = 0;
            for j in 0..4 {
                if bi_vec[4 * i + j] {
                    n += 1 << (3 - j)
                }
            }
            if n < 10 {
                ox_vecc.push('0' as u8 + n as u8);
            } else {
                ox_vecc.push('a' as u8 + n as u8 - 10);
            }
        }
        String::from_utf8(ox_vecc).unwrap().trim_start_matches('0').to_string()
    }

    //507. 完美数
    pub fn check_perfect_number(num: i32) -> bool {
        (1..10001).filter(|&i| i * i <= num && num % i == 0 && num != i)
            .fold(0, |p, q| p + q + num / q) == 2 * num
    }
}