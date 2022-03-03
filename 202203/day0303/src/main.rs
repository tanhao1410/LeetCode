use std::cell::RefCell;
use std::rc::Rc;

fn main() {
    println!("Hello, world!");
}

impl Solution {
    //973. 最接近原点的 K 个点
    pub fn k_closest(mut points: Vec<Vec<i32>>, k: i32) -> Vec<Vec<i32>> {
        points.sort_unstable_by_key(|v| v[0] * v[0] + v[1] * v[1]);
        points
            .into_iter()
            .take(k as usize)
            .collect()
    }
    //122. 买卖股票的最佳时机 II
    pub fn max_profit2(prices: Vec<i32>) -> i32 {
        let mut res = 0;
        let mut buy_res = -prices[0];
        for i in 0..prices.len() {
            buy_res = buy_res.max(res - prices[i]);
            res = res.max(buy_res + prices[i]);
        }
        res
    }
    //121. 买卖股票的最佳时机
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        //写一个dp,表示后面的最大值
        let mut dp = vec![0; prices.len()];
        let mut res = 0;
        for i in (0..prices.len() - 1).rev() {
            dp[i] = dp[i + 1].max(prices[i + 1]);
            res = res.max(dp[i] - prices[i]);
        }
        res
    }
    //48. 旋转图像
    pub fn rotate(matrix: &mut Vec<Vec<i32>>) {
        let n = matrix.len();
        for i in 0..n / 2 {
            for j in 0..(n + 1) / 2 {
                let temp = matrix[i][j];
                matrix[i][j] = matrix[n - 1 - j][i];
                matrix[n - 1 - j][i] = matrix[n - 1 - i][n - 1 - j];
                matrix[n - 1 - i][n - 1 - j] = matrix[j][n - 1 - i];
                matrix[j][n - 1 - i] = temp;
            }
        }
    }
    //258. 各位相加
    pub fn add_digits(num: i32) -> i32 {
        if num < 9 {
            return num;
        }
        if num % 9 == 0 {
            return 9;
        }
        return num % 9;
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

//173. 二叉搜索树迭代器
struct BSTIterator {
    index: usize,
    datas: Vec<i32>,
}

impl BSTIterator {
    fn new(root: Option<Rc<RefCell<TreeNode>>>) -> Self {
        let mut res = Self {
            index: 0,
            datas: vec![],
        };
        res.dfs(&root);
        res
    }

    fn dfs(&mut self, root: &Option<Rc<RefCell<TreeNode>>>) {
        if root.is_some() {
            self.dfs(&root.as_ref().unwrap().borrow().left);
            self.datas.push(root.as_ref().unwrap().borrow().val);
            self.dfs(&root.as_ref().unwrap().borrow().right);
        }
    }


    fn next(&mut self) -> i32 {
        self.index += 1;
        self.datas[self.index - 1]
    }

    fn has_next(&self) -> bool {
        self.index < self.datas.len()
    }
}

struct Solution;