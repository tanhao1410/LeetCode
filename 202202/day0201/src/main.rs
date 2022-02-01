fn main() {
    println!("Hello, world!");
}

//63. 不同路径 II
pub fn unique_paths_with_obstacles(obstacle_grid: Vec<Vec<i32>>) -> i32 {
    let mut dp = obstacle_grid.clone();
    dp[0][0] ^= 1;
    for i in 1..dp.len() {
        if obstacle_grid[i][0] == 1 {
            dp[i][0] = 0;
        } else {
            dp[i][0] = dp[i - 1][0];
        }
    }
    for i in 1..dp[0].len() {
        if obstacle_grid[0][i] == 1 {
            dp[0][i] = 0;
        } else {
            dp[0][i] = dp[0][i - 1];
        }
    }
    for i in 1..dp.len() {
        for j in 1..dp[0].len() {
            if obstacle_grid[i][j] == 1 {
                dp[i][j] = 0;
            } else {
                dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
            }
        }
    }
    dp[dp.len() - 1][dp[0].len() - 1]
}