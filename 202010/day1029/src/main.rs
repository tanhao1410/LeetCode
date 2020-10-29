fn main() {
    println!("Hello, world!");
    let num = Solution::count_digit_one(17);
    println!("{}", num);
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
}

struct Solution {}

impl Solution {
    //每日一题：129.求根到叶子节点数字之和
    pub fn sum_numbers(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        //思路：深度优先遍历法或递归法
        fn dfs(root: Option<Rc<RefCell<TreeNode>>>, parent_val: i32) -> Vec<i32> {
            let mut nums = vec![];
            if let Some(node) = root {
                let val = node.borrow().val + 10 * parent_val;
                //得到左右根节点的
                if node.borrow().left.is_some() || node.borrow().right.is_some() {
                    nums.append(&mut dfs(node.borrow().left.clone(), val));
                    nums.append(&mut dfs(node.borrow().right.clone(), val));
                } else {
                    nums.push(val);
                }
            }
            nums
        }
        dfs(root, 0).iter().sum()
    }

    //剑指 Offer 43. 1～n 整数中 1 出现的次数
    //注意10,105等
    pub fn count_digit_one(n: i32) -> i32 {
        if n == 0 {
            return 0;
        }
        if n < 10 {
            return 1;
        }
        //求10的n次方
        fn getN(mut n: usize) -> i32 {
            let mut res = 1;
            while n > 0 {
                res *= 10;
                n -= 1;
            }
            res
        }
        let mut res = 0;
        //先来一个map记录对应关系
        let mut m = [0; 15];
        m[1] = 1;
        for i in 2..10 {
            m[i] = getN(i - 1) + 10 * m[i - 1];
        }
        //求n的位数-1
        let mut bit_num = 0;
        //求n的最高位和位数
        let mut nn = n;
        while nn >= 10 {
            nn = nn / 10;
            bit_num += 1;
        }
        let base = m[bit_num];
        //看n的最高一位是多少，如果是1，如果大于1，在上面nn即最高一位。
        let next = n % (getN(bit_num) * nn);
        if nn == 1 {
            res = base + next + 1 + Solution::count_digit_one(next);
        } else {
            res = base + getN(bit_num) + (nn - 1) * m[bit_num] + Solution::count_digit_one(next);
        }
        return res as i32;
    }
}

struct MedianFinder {
    nums: Vec<i32>,
}

//剑指 Offer 41. 数据流中的中位数
impl MedianFinder {
    /** initialize your data structure here. */
    fn new() -> Self {
        MedianFinder {
            nums: Vec::new()
        }
    }

    fn add_num(&mut self, num: i32) {
        //要排序的,直接放到最后一位
        if self.nums.len() == 0 || num >= *self.nums.last().unwrap() {
            self.nums.push(num);
        } else {
            let mut i = 0;
            while self.nums[i] < num {
                i += 1;
            }
            self.nums.insert(i, num);
        }
    }

    fn find_median(&self) -> f64 {
        let mut res = 0.0;
        let size = self.nums.len();
        if size == 0 {
            res
        } else if size % 2 == 0 {
            self.nums[size / 2 - 1] + self.nums[size / 2] as f64 / 2.0
        } else {
            self.nums[size / 2] as f64
        }
    }
}
