fn main() {
    println!("Hello, world!");
}

impl Solution {
    //322. 零钱兑换
    pub fn coin_change(coins: Vec<i32>, amount: i32) -> i32 {
        let mut dp = vec![vec![-1; amount as usize + 1]; coins.len()];
        for i in 0..dp.len() {
            dp[i][0] = 0;
            for j in 1..dp[0].len() {
                if j >= coins[i] as usize && dp[i][j - coins[i] as usize] != -1 {
                    dp[i][j] = dp[i][j - coins[i] as usize] + 1;
                    if i > 0 && dp[i - 1][j] != -1 && dp[i - 1][j] < dp[i][j] {
                        dp[i][j] = dp[i - 1][j]
                    }
                } else if i > 0 {
                    dp[i][j] = dp[i - 1][j];
                }
            }
        }
        *dp.last().unwrap().last().unwrap()
    }
    //537. 复数乘法
    pub fn complex_number_multiply(num1: String, num2: String) -> String {
        // a b c d
        let transfer = |num: &str| {
            let i = num.find('+').unwrap();
            (num[..i].parse::<i32>().unwrap(), num[i + 1..num1.len() - 1].parse::<i32>().unwrap())
        };
        let (a, b) = transfer(&num1);
        let (c, d) = transfer(&num2);
        // let mut res = String::new();
        // res.push_str(&(a * c - b * d).to_string());
        // res.push('+');
        // res.push_str(&(a * d + b * c).to_string());
        // res.push('i');
        // res
        format!("{}+{}i", (a * c - b * d), (a * d + b * c))
    }
}

struct Solution;
