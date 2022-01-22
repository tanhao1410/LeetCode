fn main() {
    println!("Hello, world!");
}

//53. 最大子数组和
pub fn max_sub_array(nums: Vec<i32>) -> i32 {
    //dp[i]以nums[i]结尾的最大子数组和，if dp[i] > 0 则 dp[i-1] = dp[i] + nums[i]
    // nums
    //     .into_iter()
    //     .fold((i32::MIN,-1),|(res,pre),num|{
    //         let pre = 0.max(pre) + num;
    //         (res.max(pre),pre)
    //     }) // res,dp[i-1],
    //     .0
    let mut res = nums[0];
    let mut pre =nums[0];
    for i in 1..nums.len() {
        if pre > 0{
            pre += nums[i];
        }else{
            pre = nums[i];
        }
        res = res.max(pre);
    }
    res
}

//1332. 删除回文子序列
pub fn remove_palindrome_sub(s: String) -> i32 {
    //思路：先删除最大的回文序列,dp[i]代表以i为中心最大的回文序列
    let mut dp = vec![1;s.len()];

    0
}