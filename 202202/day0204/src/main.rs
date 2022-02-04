fn main() {
    println!("Hello, world!");
}

//300. 最长递增子序列
pub fn length_of_lis(nums: Vec<i32>) -> i32 {
    // o (n ^2)
    //以 nums[i]结尾的最长递增子序列
    let mut dp = nums.clone();
    dp[0] = 1;
    for i in 1..nums.len() {
        //找到它前面的 比nums[i]小的 最大 dp[i]
        let mut max = 0;
        for j in 0..i {
            if nums[j] < nums[i] {
                max = max.max(dp[j]);
            }
        }
        dp[i] = max + 1;
    }
    dp.into_iter().max().unwrap()
}