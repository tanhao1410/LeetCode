fn main() {
    println!("Hello, world!");
}

//198. 打家劫舍
pub fn rob(mut nums: Vec<i32>) -> i32 {
    //dp[i] 偷了第i家的情况下的最大值,则第i-1家不能偷。dp[i] = max{nums[i] + dp[i - 2] or dp[i - 3]}
    for i in 2..nums.len() {
        nums[i] += nums[i - 2].max(*nums.get(i - 3).unwrap_or(&0));
    }
    nums.into_iter().rev().take(2).max().unwrap()
}

//2029. 石子游戏 IX
pub fn stone_game_ix(stones: Vec<i32>) -> bool {
    true
}