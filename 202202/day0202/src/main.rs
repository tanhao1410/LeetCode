fn main() {
    println!("Hello, world!");
}

//231. 2 的幂
pub fn is_power_of_two(n: i32) -> bool {
    n > 0 && n.count_ones() == 1
}

//221. 最大正方形
pub fn maximal_square(matrix: Vec<Vec<char>>) -> i32 {
    let mut dp = vec![vec![0; matrix[0].len()]; matrix.len()];
    for i in 0..dp.len() {
        for j in 0..dp[0].len() {
            if matrix[i][j] == '0' {} else if i > 0 && j > 0 {
                dp[i][j] = 1 + dp[i - 1][j].min(dp[i][j - 1].min(dp[i - 1][j - 1]))
            } else {
                dp[i][j] = 1;
            }
        }
    }
    let max_edge = *dp.iter().flat_map(|v| v.iter()).max().unwrap();
    max_edge * max_edge
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