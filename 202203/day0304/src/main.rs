fn main() {
    println!("Hello, world!");
}

use std::rc::Rc;
use std::cell::RefCell;

impl Solution {
    //1630. 等差子数组
    pub fn check_arithmetic_subarrays(nums: Vec<i32>, l: Vec<i32>, r: Vec<i32>) -> Vec<bool> {
        let is_subarrays = |mut nums: Vec<i32>| {
            nums.sort_unstable();
            let mut dis = nums[1] - nums[0];
            for i in 2..nums.len() {
                if nums[i] - nums[i - 1] != dis {
                    return false;
                }
            }
            true
        };
        l.into_iter()
            .zip(r.into_iter())
            .map(|(i, j)| is_subarrays((&nums[i as usize..=j as usize]).to_vec()))
            .collect()
    }
    //2104. 子数组范围和
    pub fn sub_array_ranges(nums: Vec<i32>) -> i64 {
        let mut res = 0;
        for i in 0..nums.len() - 1 {
            let mut min = nums[i];
            let mut max = nums[i];
            for j in i + 1..nums.len() {
                min = min.min(nums[j]);
                max = max.max(nums[j]);
                res += (max - min) as i64;
            }
        }
        res
    }
    //1987. 不同的好子序列数目
    pub fn number_of_unique_good_subsequences(binary: String) -> i32 {
        //动态规划，dp[i] 以binary[i] 以0开头的子序列数目。以1开头的子序列数目
        let mut dp = vec![(0, 0); binary.len()];
        const MOD: i32 = 1000000007;
        let bytes = binary.as_bytes();
        let mut contains_zero = 0;
        for i in (0..bytes.len()).rev() {
            let pre = *dp.get(i + 1).unwrap_or(&(0, 0));
            dp[i] = pre;
            if bytes[i] == b'0' {
                contains_zero = 1;
                dp[i].0 = (pre.0 + pre.1 + 1) % MOD;
            } else {
                dp[i].1 = (pre.0 + pre.1 + 1) % MOD;
            }
        }
        //以1开头的子序列 + 0
        dp[0].1 + contains_zero
    }
    //236. 二叉树的最近公共祖先
    pub fn lowest_common_ancestor(root: Option<Rc<RefCell<TreeNode>>>, p: Option<Rc<RefCell<TreeNode>>>, q: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        //思路：如果根与p,或q相等，则返回root。否则，可能在于root的左子树或右子树。如果左与右都没有，说明左右各一样，返回root。否则返回对应的即可。
        if root.is_none() {
            return None;
        }
        let root_val = root.as_ref().unwrap().borrow().val;
        if root_val == p.as_ref().unwrap().borrow().val || root_val == q.as_ref().unwrap().borrow().val {
            return root;
        }
        let pp = Some(p.as_ref().unwrap().clone());
        let qq = Some(q.as_ref().unwrap().clone());
        let left = Self::lowest_common_ancestor(root.as_ref().unwrap().borrow_mut().left.take(), pp, qq);
        let right = Self::lowest_common_ancestor(root.as_ref().unwrap().borrow_mut().right.take(), p, q);
        if left.is_some() && right.is_some() {
            return root;
        } else if left.is_none() {
            return right;
        } else {
            return left;
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