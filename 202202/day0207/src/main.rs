fn main() {
    println!("Hello, world!");
}

//279. 完全平方数
pub fn num_squares(n: i32) -> i32 {
    //完全平方数有哪些【1,4,9,16,...,10000】共100个。
    // dp[上述] = 1；
    //for [上述] 遍历 完全平方数 + 1 dp[]
    //dp[j] = dp[j - nums[i]] + 1 .min(self);
    //从小往大的开始算
    //dp[j + nums[i]] = dp[j] + 1 .min(self)
    let mut nums = vec![];
    for i in 1..10000 {
        if i * i > n {
            break;
        }
        nums.push(i * i);
    }
    let mut dp = vec![i32::MAX; n as usize + 1];
    for i in 0..nums.len() {
        dp[nums[i] as usize] = 1;
    }
    for j in 1..=n {
        for i in 0..nums.len() {
            if j + nums[i] <= n {
                dp[j as usize + nums[i] as usize] = dp[j as usize + nums[i] as usize].min(dp[j as usize] + 1);
            }
        }
    }
    dp[n as usize]
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