fn main() {
    println!("Hello, world!");
    println!("{}", Solution::min_sub_array_len(7, vec![2, 3, 1, 2, 4, 3]));
}

impl Solution {
    //209. 长度最小的子数组
    pub fn min_sub_array_len(target: i32, nums: Vec<i32>) -> i32 {
        let mut start = 0;
        let mut sum = nums[0];
        let mut res = if sum < target { i32::MAX } else { 1 };
        for end in 1..=nums.len() {
            if sum >= target {
                while sum >= target {
                    sum -= nums[start];
                    start += 1;
                }
                //因为此时窗口里面的值小于target了，因此是+2
                res = res.min(end as i32 - start as i32 + 1);
            }
            if end < nums.len() {
                //不够
                sum += nums[end];
            }
        }
        if res == i32::MAX { 0 } else { res }
    }
    //201. 数字范围按位与
    pub fn range_bitwise_and(m: i32, n: i32) -> i32 {
        (0..32)
            .map(|e| 1 << 31 - e)
            .take_while(|&e| m & e == n & e)
            .map(|e| m & e)
            .sum()
    }
    //24. 两两交换链表中的节点
    pub fn swap_pairs(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() || head.as_ref().unwrap().next.is_none() {
            return head;
        }
        let mut res = Some(Box::new(ListNode::new(0)));
        let mut p = res.as_mut().unwrap();
        //每一次弹出两个
        while let Some(mut node) = head {
            let swap_node = node.next.take();
            if let Some(mut swap_node) = swap_node {
                head = swap_node.next.take();
                //本来是node -> swap_node =>断开了
                swap_node.next = Some(node);
                p.next = Some(swap_node);
                p = p.next.as_mut().unwrap().next.as_mut().unwrap();
            } else {
                p.next = Some(node);
                break;
            }
        }
        res.as_mut().unwrap().next.take()
    }

    //230. 二叉搜索树中第K小的元素
    pub fn kth_smallest(mut root: Option<Rc<RefCell<TreeNode>>>, k: i32) -> i32 {
        if let Some(node) = root {
            let left_length = Self::tree_length(&node.borrow().left);
            if left_length == k - 1 {
                node.borrow().val
            } else if left_length > k - 1 {
                Self::kth_smallest(node.borrow_mut().left.take(), k)
            } else {
                Self::kth_smallest(node.borrow_mut().right.take(), k - left_length - 1)
            }
        } else {
            unreachable!()
        }
    }

    fn tree_length(root: &Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut res = 0;
        if let Some(node) = root {
            res += 1;
            res += Self::tree_length(&node.borrow().left);
            res += Self::tree_length(&node.borrow().right);
        }
        res
    }
    //112. 路径总和
    pub fn has_path_sum(mut root: Option<Rc<RefCell<TreeNode>>>, target_sum: i32) -> bool {
        if let Some(node) = root {
            let cur_val = node.borrow().val;
            if target_sum == cur_val && node.borrow().left.is_none() && node.borrow().right.is_none() {
                return true;
            }
            let mut ref_mut = node.borrow_mut();
            return Self::has_path_sum(ref_mut.right.take(), target_sum - cur_val)
                || Self::has_path_sum(ref_mut.left.take(), target_sum - cur_val);
        }
        false
    }
    //456. 132 模式
    pub fn find132pattern(nums: Vec<i32>) -> bool {
        if nums.len() < 3 {
            return false;
        }
        //思路：单调栈-找后面比自己小的最大值，前面的最小值,不包括自己
        let mut pre_min = vec![1000000001; nums.len()];
        for i in 1..nums.len() {
            pre_min[i] = pre_min[i - 1].min(nums[i - 1]);
        }
        //单调栈
        let mut stack = vec![*nums.last().unwrap()];
        for i in (1..nums.len() - 1).rev() {
            let cur = nums[i];
            if stack.is_empty() {
                stack.push(cur);
            } else {
                //判断它后面最大数是多少
                let min = pre_min[i];
                let mut top = *stack.last().unwrap();
                while !stack.is_empty() && *stack.last().unwrap() < cur {
                    top = stack.pop().unwrap();
                }
                if top > min && cur > top {
                    return true;
                }
                stack.push(cur);
            }
        }
        false
    }
    //66. 加一
    pub fn plus_one(mut digits: Vec<i32>) -> Vec<i32> {
        //主要是判断是有无进位
        let mut flag = 1;
        for digit in digits.iter_mut().rev() {
            if flag == 0 {
                break;
            }
            *digit += flag;
            flag = *digit / 10;
            *digit %= 10;
        }
        if flag == 1 {
            let mut res = vec![1];
            res.append(&mut digits);
            return res;
        }
        digits
    }
    //2016. 增量元素之间的最大差值
    pub fn maximum_difference(nums: Vec<i32>) -> i32 {
        let mut max = 0;
        let mut res = -1;
        for i in nums.into_iter().rev() {
            if max > i {
                res = res.max(max - i);
            } else {
                max = i;
            }
        }
        res
    }
}

struct Solution;

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

#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode {
            next: None,
            val,
        }
    }
}