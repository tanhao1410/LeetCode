fn main() {
    println!("Hello, world!");
}

use std::rc::Rc;
use std::cell::RefCell;

impl Solution {
    //剑指 Offer II 054. 所有大于等于节点的值之和
    pub fn convert_bst(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        //如果节点左右都无，值不变。需要看上面是否有传递的值。
        //右节点有值的话，先改变右节点。然后，将右子树的和+上
        //对于左节点，每一个左节点。开始时加上上面传递来的值。
        if root.is_none() {
            return None;
        }
        Self::convert_bst_preval(Some(root.as_ref().unwrap().clone()), 0);
        root
    }
    fn convert_bst_preval(root: Option<Rc<RefCell<TreeNode>>>, pre: i32) {
        if let Some(node) = root {
            if node.borrow().left.is_none() && node.borrow().right.is_none() {
                node.borrow_mut().val += pre;
                return;
            }
            if node.borrow().right.is_some() {
                Self::convert_bst_preval(Some(node.borrow().right.as_ref().unwrap().clone()), pre);
                //
                let right_sum = Self::tree_max(&node.borrow().right);
                node.borrow_mut().val += right_sum;
            } else {
                node.borrow_mut().val += pre;
            }
            let cur_val = node.borrow().val;
            if node.borrow().left.is_some() {
                Self::convert_bst_preval(Some(node.borrow().left.as_ref().unwrap().clone()), cur_val);
                //求出left后，
            }
            //一个节点，它的值应该是自己+右子树的最大树
        }
    }

    fn tree_max2(root: &Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut sum = 0;
        if root.is_some() {
            if root.as_ref().unwrap().borrow().left.is_some() {
                return Self::tree_max2(&root.as_ref().unwrap().borrow().left);
            }
            sum = root.as_ref().unwrap().borrow().val;
        }
        sum
    }

    //剑指 Offer II 053. 二叉搜索树中的中序后继
    pub fn inorder_successor(root: Option<Rc<RefCell<TreeNode>>>, p: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        let root = root.as_ref().unwrap();
        let p_val = p.as_ref().unwrap().borrow().val;
        if p_val == root.borrow().val {
            //返回右子树的最小值
            return Self::tree_min(root.borrow_mut().right.take());
        } else if p_val < root.borrow().val {
            //找左子树的最大值
            if root.borrow().left.is_none() {
                return None;
            }
            let left_max = Self::tree_max(Some(root.borrow().left.as_ref().unwrap().clone()));
            if left_max.is_some() && left_max.as_ref().unwrap().borrow().val == p_val {
                return Some(root.clone());
            }
            return Self::inorder_successor(root.borrow_mut().left.take(), p);
        } else {
            return Self::inorder_successor(root.borrow_mut().right.take(), p);
        }
    }

    fn tree_max(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        if root.is_none() {
            return None;
        }
        if root.as_ref().unwrap().borrow().right.is_none() {
            return root;
        }
        Self::tree_max(Some(root.as_ref().unwrap().borrow().right.as_ref().unwrap().clone()))
    }

    fn tree_min(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        if root.is_none() {
            return None;
        }
        if root.as_ref().unwrap().borrow().left.is_none() {
            return root;
        }
        Self::tree_min(Some(root.as_ref().unwrap().borrow().left.as_ref().unwrap().clone()))
    }


    //剑指 Offer II 052. 展平二叉搜索树
    pub fn increasing_bst(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        Self::create_tree(&Self::mid_read_tree(&root))
    }

    fn create_tree(nums: &[i32]) -> Option<Rc<RefCell<TreeNode>>> {
        if nums.len() == 0 {
            return None;
        }
        let mut root = TreeNode::new(nums[0]);
        root.right = Self::create_tree(&nums[1..]);
        Some(Rc::new(RefCell::new(root)))
    }

    fn mid_read_tree(root: &Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut res = vec![];
        if let Some(node) = root {
            res.append(&mut Self::mid_read_tree(&node.borrow_mut().left));
            res.push(node.borrow().val);
            res.append(&mut Self::mid_read_tree(&node.borrow().right))
        }
        res
    }
    //841. 钥匙和房间
    pub fn can_visit_all_rooms(rooms: Vec<Vec<i32>>) -> bool {
        //深度优先遍历
        let mut stack = vec![0];
        let mut used = vec![false; rooms.len()];
        used[0] = true;
        while let Some(room) = stack.pop() {
            let new_rooms = &rooms[room];
            for new_room in new_rooms {
                if !used[*new_room as usize] {
                    used[*new_room as usize] = true;
                    stack.push(*new_room as usize);
                }
            }
        }
        used.into_iter().all(|e| e)
    }
    //11. 盛最多水的容器
    pub fn max_area(height: Vec<i32>) -> i32 {
        //双指针
        let mut start = 0;
        let mut end = height.len() - 1;
        let mut res = 0;
        while end > start {
            res = res.max((end - start) as i32 * height[start].min(height[end]));
            if height[start] < height[end] {
                start += 1;
            } else {
                end -= 1;
            }
        }
        res
    }
    //1557. 可以到达所有点的最少点数目
    pub fn find_smallest_set_of_vertices(n: i32, edges: Vec<Vec<i32>>) -> Vec<i32> {
        let mut reached = vec![false; n as usize];
        for edge in &edges {
            //代表这个节点可以通过其他节点到达
            let end = edge[1] as usize;
            reached[end] = true;
        }
        reached
            .into_iter()
            .enumerate()
            .filter(|&v| !v.1)
            .map(|v| v.0 as i32)
            .collect()
    }
    //42. 接雨水
    pub fn trap(height: Vec<i32>) -> i32 {
        //它前面的最大值，它后面的最大值
        let mut pre_max = vec![0; height.len()];
        let mut tail_max = vec![0; height.len()];
        for i in 0..height.len() {
            if i > 0 {
                pre_max[i] = pre_max[i - 1].max(height[i]);
                tail_max[height.len() - 1 - i] = tail_max[height.len() - i].max(height[height.len() - 1 - i]);
            } else {
                pre_max[i] = height[i];
                tail_max[height.len() - 1 - i] = height[height.len() - 1 - i];
            }
        }
        //每一个高度柱能放多少水呢，取决于，它左右两边的最高值的最小值
        pre_max.into_iter()
            .zip(tail_max.into_iter())
            .map(|e| e.0.min(e.1))
            .zip(height.into_iter())
            .map(|v| v.0 - v.1)
            .sum()
    }
    //503. 下一个更大元素 II
    pub fn next_greater_elements(nums: Vec<i32>) -> Vec<i32> {
        let mut stack = vec![];
        let mut res = vec![0; nums.len()];
        //单调栈,先进行一圈，好让后面的知道有哪些比它大
        for i in (0..nums.len()).rev() {
            while !stack.is_empty() && stack[stack.len() - 1] <= nums[i] {
                stack.pop();
            }
            stack.push(nums[i]);
        }
        for i in (0..nums.len()).rev() {
            while !stack.is_empty() && stack[stack.len() - 1] <= nums[i] {
                stack.pop();
            }
            if stack.is_empty() {
                res[i] = -1;
            } else {
                res[i] = stack[stack.len() - 1];
            }
            stack.push(nums[i]);
        }
        res
    }
    //521. 最长特殊序列 Ⅰ
    pub fn find_lu_slength(a: String, b: String) -> i32 {
        if a == b {
            -1
        } else {
            a.len().max(b.len()) as i32
        }
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