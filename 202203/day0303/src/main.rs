fn main() {
    println!("Hello, world!");
}

impl Solution {
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

struct Solution;