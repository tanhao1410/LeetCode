fn main() {
    println!("Hello, world!");
    //println!("{:?}",Solution::next_permutation(&mut vec![3, 2, 1]));
    println!("{}", Solution::compare_version("1.00002".to_string(), "001.12.0.01.0.0.1".to_string()))
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
use std::collections::hash_map::VacantEntry;

impl Solution {


    //222. 完全二叉树的节点个数
    pub fn count_nodes(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        //思路：最简单的思路，遍历即可。
        let mut res = 0;
        let mut queue = vec![];
        if let Some(node) = root{
            queue.push(node);
        }
        while !queue.is_empty(){
            let node = queue.pop().unwrap();
            res += 1;
            if node.borrow_mut().left.is_some(){
                queue.push(node.borrow_mut().left.clone().unwrap());
            }
            if node.borrow_mut().right.is_some(){
                queue.push(node.borrow_mut().right.clone().unwrap());
            }
        }
        res
    }

    //199. 二叉树的右视图
    pub fn right_side_view(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        //思路：层次遍历可以解决，
        let mut res = vec![];
        let mut queue = vec![];
        if let Some(node) = root {
            queue.push(node);
        }
        while !queue.is_empty() {
            let queue_size = queue.len();
            res.push(queue[queue_size - 1].borrow_mut().val);
            for i in 0..queue_size {
                let left = queue[i].borrow_mut().left.clone();
                let right = queue[i].borrow_mut().right.clone();
                if left.is_some() {
                    queue.push(left.unwrap());
                }
                if right.is_some() {
                    queue.push(right.unwrap());
                }
            }
            queue = queue[queue_size..].to_vec();
        }
        res
    }

    //面试题 04.03. 特定深度节点链表
    pub fn list_of_depth(tree: Option<Rc<RefCell<TreeNode>>>) -> Vec<Option<Box<ListNode>>> {

        //层次遍历
        let mut res = vec![];
        let mut queue = vec![];
        if tree.is_some() {
            queue.push(tree.unwrap());
        }
        while !queue.is_empty() {
            let queue_size = queue.len();
            let mut head = None;
            for i in 0..queue_size {
                let mut tree_node = queue[i].clone();
                let mut list_node = ListNode::new(queue[queue_size - 1 - i].borrow_mut().val);
                if head.is_none() {
                    head = Some(Box::new(list_node));
                } else {
                    list_node.next = head;
                    head = Some(Box::new(list_node));
                }
                if tree_node.borrow_mut().left.is_some() {
                    queue.push(tree_node.borrow_mut().left.clone().unwrap());
                }
                if tree_node.borrow_mut().right.is_some() {
                    queue.push(tree_node.borrow_mut().right.clone().unwrap());
                }
            }
            queue = queue[queue_size..].to_vec();
            res.push(head);
        }
        res
    }

    //200. 岛屿数量
    pub fn num_islands(grid: Vec<Vec<char>>) -> i32 {
        //思路：先找到一个1，然后，将与它相连的都变为0，继续走，直到结束。
        let mut res = 0;
        let mut grid2 = grid.clone();
        fn set_zero(i: usize, j: usize, grid2: &mut Vec<Vec<char>>) {
            grid2[i][j] = '0';
            if i + 1 < grid2.len() && grid2[i + 1][j] == '1' {
                set_zero(i + 1, j, grid2);
            }
            if i as i32 - 1 >= 0 && grid2[i - 1][j] == '1' {
                set_zero(i - 1, j, grid2);
            }
            if j as i32 - 1 >= 0 && grid2[i][j - 1] == '1' {
                set_zero(i, j - 1, grid2);
            }
            if j + 1 < grid2[0].len() && grid2[i][j + 1] == '1' {
                set_zero(i, j + 1, grid2);
            }
        }
        for i in 0..grid.len() {
            for j in 0..grid[0].len() {
                if grid[i][j] == '1' && grid2[i][j] == '1' {
                    res += 1;
                    set_zero(i, j, &mut grid2);
                }
            }
        }
        res
    }

    //剑指 Offer 42. 连续子数组的最大和
    pub fn max_sub_array(nums: Vec<i32>) -> i32 {
        //动态规划算法，以该数字结尾的最大值
        let mut dp = nums.clone();
        let mut res = std::i32::MIN;
        for i in 1..nums.len() {
            if dp[i - 1] > 0 {
                dp[i] = dp[i - 1] + nums[i];
            }
        }
        for i in dp {
            if res < i {
                res = i
            }
        }
        res
    }

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
        if nums_size < 2 {
            return;
        }
        let mut i = nums_size as i32 - 2;
        while i >= 0 {
            //从后面找到一个比它大的最小的一个数
            let mut more_min = i as usize;
            for j in i as usize..nums_size {
                if nums[j] > nums[i as usize] {
                    if more_min == i as usize {
                        more_min = j;
                    } else if nums[j] < nums[more_min] {
                        more_min = j;
                    }
                }
            }
            if more_min != i as usize {
                //把该数放置在i位置上，后面按顺序摆放即可
                let temp = nums[i as usize];
                nums[i as usize] = nums[more_min];
                nums[more_min] = temp;
                //后面的进行排序
                for j in i as usize + 1..nums_size - 1 {
                    for k in j + 1..nums_size {
                        if nums[k] < nums[j] {
                            let temp = nums[j];
                            nums[j] = nums[k];
                            nums[k] = temp;
                        }
                    }
                }
                return;
            } else {
                //若不存在这样的数，则 i -= 1;如果直到结束都不存在,说明数组是递减的，反转即可。
                i -= 1;
            }
        }
        nums.reverse();
    }
}
