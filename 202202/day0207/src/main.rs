fn main() {
    println!("Hello, world!");
}

//377. 组合总和 Ⅳ
pub fn combination_sum4(nums: Vec<i32>, target: i32) -> i32 {
    // let mut dp = vec![0; target as usize + 1];
    // dp[0] = 1;
    // for i in 0..nums.len() {
    //     for j in nums[i]..=target {
    //         dp[j] += dp[(j - nums[i]) as usize];
    //     }
    // }
    // dp[target as usize]
    //上面是不在乎顺序的组合，对于与顺序相关的组合？
    let mut dp = vec![0; target as usize + 1];
    dp[0] = 1;
    for i in 0..=target as usize {
        for j in 0..nums.len() {
            if i >= nums[j] as usize {
                dp[i] += dp[i - nums[j] as usize];
            }
        }
    }
    dp[target as usize]
}