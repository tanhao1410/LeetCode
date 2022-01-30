fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //120. 三角形最小路径和
    pub fn minimum_total(triangle: Vec<Vec<i32>>) -> i32 {
        //动态规划：dp[i] ,只记录某一行的最小值
        let mut dp = vec![0; triangle.len()];
        dp[0] = triangle[0][0];
        for i in 1..triangle.len() {
            for j in 0..triangle[i].len() {
                let temp = dp.clone();
                if j == 0 {
                    dp[0] += triangle[i][j];
                } else if j > 0 && j < triangle.len() - 1 {
                    dp[j] = triangle[i][j] + temp[j].min(temp[j - 1]);
                } else {
                    dp[j] = triangle[i][j] + temp[j - 1];
                }
            }
        }
        dp.into_iter().min().unwrap()
    }

    //931. 下降路径最小和
    pub fn min_falling_path_sum(matrix: Vec<Vec<i32>>) -> i32 {
        //动态规划:dp[i][j] 到达matirix[i][j] 位置路径和
        let mut dp = matrix.clone();
        for i in 1..dp.len() {
            for j in 0..dp[0].len() {
                dp[i][j] = matrix[i][j];
                if j > 0 && j < dp[0].len() - 1 {
                    dp[i][j] += dp[i - 1][j].min(dp[i - 1][j - 1]).min(dp[i - 1][j + 1]);
                } else if j == 0 {
                    dp[i][j] += dp[i - 1][j].min(dp[i - 1][j + 1]);
                } else {
                    dp[i][j] += dp[i - 1][j].min(dp[i - 1][j - 1]);
                }
            }
        }
        *dp[dp.len() - 1].iter().min().unwrap()
    }
}