use std::arch::x86_64::_mm256_mask_compress_epi32;

fn main() {
    println!("Hello, world!");
}

//213. 打家劫舍 II
pub fn rob2(nums: Vec<i32>) -> i32 {
    //区别在于是环形的 dp[i] = max{dp[i-2] + nums[i],dp[i - 1]} 限制是，dp[0] 与dp[len - 1]不可共存
    //思路1：转化为198，dp[n-1] 或  去除掉nums[0] 后，求dp[n-1]。大的为返回结果。
    let rob = |nums: &[i32]| {
        if nums.len() < 2 {
            return nums[0];
        }
        let mut pre1 = nums[0];
        let mut pre2 = nums[1];
        for i in 2..nums.len() {
            let temp = pre2.max(pre1 + nums[i]);
            pre1 = pre2.max(pre1);
            pre2 = temp;
        }
        pre1.max(pre2)
    };
    if nums.len() == 1 {
        return nums[0];
    }
    rob(&nums[1..]).max(rob(&nums[..nums.len() - 1]))
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