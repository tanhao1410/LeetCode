use std::rc::Rc;
use std::cell::RefCell;

fn main() {
    println!("Hello, world!");
}

#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    pub(crate) fn new(val: i32) -> Self {
        ListNode {
            next: None,
            val,
        }
    }

    //1290. 二进制链表转整数
    pub fn get_decimal_value(head: Option<Box<ListNode>>) -> i32 {
        let mut head = head;
        let mut res = 0;
        while let Some(mut node) = head.take() {
            res <<= 1;
            res += node.val;
            head = node.next.take();
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

    //404. 左叶子之和
    pub fn sum_of_left_leaves(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        //只有根节点才算
        let mut res = 0;
        if let Some(mut node) = root {
            if let Some(left) = node.borrow_mut().left.take() {
                if left.borrow().left.is_none() && left.borrow().right.is_none() {
                    res += left.borrow().val;
                } else {
                    res += TreeNode::sum_of_left_leaves(Some(left));
                }
            }
            res += TreeNode::sum_of_left_leaves(node.borrow_mut().right.take());
        }
        res
    }
}

struct Solution {}

impl Solution {

    //剑指 Offer 56 - I. 数组中数字出现的次数 - 借鉴 写法
    pub fn single_numbers3(nums: Vec<i32>) -> Vec<i32> {
        let mask = nums.iter().fold(0, |pre, &cur| pre ^ cur);
        let diff = mask & (-mask);
        let found = nums.iter().fold(0, |pre, &cur| match cur & diff {
            0 => pre,
            _ => pre ^ cur
        });

        vec![found, mask ^ found]
    }

    //剑指 Offer 56 - I. 数组中数字出现的次数-位运算方式
    pub fn single_numbers2(nums: Vec<i32>) -> Vec<i32> {
        //先求出所有的数进行异或的结果
        let mut two = 0;
        for &i in nums.iter() {
            two ^= i;
        }
        //找到为1的某位
        let mut i = -1;
        while two.count_ones() != 1 {
            two &= i;
            i <<= 1;
        }

        let mut res1 = 0;
        let mut res2 = 0;
        for &i in nums.iter(){
            if two & i != 0{
                res1 ^= i;
            }else{
                res2 ^= i;
            }
        }
        vec![res1,res2]
    }

    //剑指 Offer 56 - I. 数组中数字出现的次数
    pub fn single_numbers(nums: Vec<i32>) -> Vec<i32> {
        let mut res = vec![];
        let mut map = std::collections::HashMap::<i32, i32>::new();

        nums.iter().for_each(|&num| {
            let mut count = map.entry(num).or_insert(0);
            *count += 1;
        });

        map.iter().for_each(|(&k, &v)| {
            if v == 1 {
                res.push(k);
            }
        });
        res
    }
}