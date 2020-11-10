fn main() {
    println!("Hello, world!");
    println!("{:?}",Solution::next_permutation(&mut vec![3, 2, 1]));
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

struct Solution {}

use std::rc::Rc;
use std::cell::RefCell;

impl Solution {

    //剑指 Offer 28. 对称的二叉树
    pub fn is_symmetric(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        //思路：按层次进行判别呢？中心距
        let mut queue = vec![];
        let mut queue_dis = vec![];
        if let Some(node) = root {
            queue.push(node);
            queue_dis.push(0);
        }
        let mut is_root = true;
        while !queue.is_empty() {
            let queue_size = queue.len();
            if queue_size % 2 != 0 && !is_root {//需 指定不是root的情况
                return false;
            }
            is_root = false;
            for i in 0..queue_size {
                let node = queue[i].clone();
                let node_val = node.borrow_mut().val;
                let other_val = queue[queue_size - 1 - i].borrow_mut().val;
                let node_dis = queue_dis[i];
                if node.borrow_mut().left.is_some() {
                    queue.push(node.borrow_mut().left.clone().unwrap());
                    queue_dis.push(node_dis - 1);
                }
                if node.borrow_mut().right.is_some() {
                    queue.push(node.borrow_mut().right.clone().unwrap());
                    queue_dis.push(node_dis + 1);
                }
                if node_val != other_val
                    || -node_dis != queue_dis[queue_size - 1 - i] {
                    return false;
                }
            }
            queue = queue[queue_size..].to_vec();
            queue_dis = queue_dis[queue_size..].to_vec();
        }
        true
    }


    //剑指 Offer 54. 二叉搜索树的第k大节点
    pub fn kth_largest(root: Option<Rc<RefCell<TreeNode>>>, k: i32) -> i32 {

        //思路1：根据左右个数来求
        // fn count(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        //     if let Some(node) = root {
        //         count(node.borrow_mut().left.clone()) + count(node.borrow_mut().right.clone()) + 1
        //     } else {
        //         0
        //     }
        // }
        // let (mut left_count, mut right_count) = (0, 0);
        // if let Some(node) = root{
        //     left_count = count(node.borrow_mut().left.clone());
        //     right_count = count(node.borrow_mut().right.clone());
        //     if left_count < k{
        //         Solution::kth_largest(node.borrow_mut().right.clone(),k - left_count)
        //     }else if left_count == k{
        //         return node.borrow_mut().val;
        //     }else{
        //         Solution::kth_largest(node.borrow_mut().left.clone(),k)
        //     }
        // }else{
        //     0
        // }
        //思路2：转换成vec![]
        fn dfs(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
            let mut res = vec![];
            if let Some(node) = root {
                res.append(&mut dfs(node.borrow_mut().right.clone()));
                res.push(node.borrow_mut().val);
                res.append(&mut dfs(node.borrow_mut().left.clone()));
            }
            res
        }
        dfs(root)[k as usize - 1]

        //思路3：反中序遍历，第K个结点即可，不用专门用个数组来存
    }

    //每日一题：31. 下一个排列
    pub fn next_permutation(nums: &mut Vec<i32>) {
        //思路：就是找对应的数值位，然后后面的按顺序摆放即可。
        let nums_size = nums.len();
        if nums_size < 2{
            return;
        }
        let mut i = nums_size as i32- 2;
        while i >= 0{
            //从后面找到一个比它大的最小的一个数
            let mut more_min = i  as usize;
            for j in i  as usize.. nums_size{
                if nums[j] > nums[i as usize]{
                    if more_min == i  as usize{
                        more_min = j;
                    }else if nums[j ] < nums[more_min ]{
                        more_min = j;
                    }
                }
            }
            if more_min != i  as usize{
                //把该数放置在i位置上，后面按顺序摆放即可
                let temp = nums[i as usize];
                nums[i as usize] = nums[more_min ];
                nums[more_min ] = temp;
                //后面的进行排序
                for j in i  as usize + 1 .. nums_size-1{
                    for k in j +1 ..nums_size{
                        if nums[k ] < nums[j ]{
                            let temp = nums[j ];
                            nums[j ] = nums[k];
                            nums[k ] = temp;
                        }
                    }
                }
                return;
            }else{
                //若不存在这样的数，则 i -= 1;如果直到结束都不存在,说明数组是递减的，反转即可。
                i -= 1;
            }
        }
        nums.reverse();
    }
}
