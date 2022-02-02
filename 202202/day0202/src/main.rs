fn main() {
    println!("Hello, world!");
}

//64. 最小路径和
pub fn min_path_sum(grid: Vec<Vec<i32>>) -> i32 {
    let mut dp = grid.clone();
    for i in 0..dp.len() {
        for j in 0..dp[0].len() {
            if i > 0 && j > 0 {
                dp[i][j] += dp[i - 1][j].min(dp[i][j - 1]);
            } else if i > 0 {
                dp[i][j] += dp[i - 1][j];
            } else if j > 0 {
                dp[i][j] += dp[i][j - 1]
            }
        }
    }

    dp[dp.len() - 1][dp[0].len() - 1]
}