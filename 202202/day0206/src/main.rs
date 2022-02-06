fn main() {
    println!("Hello, world!");
    println!("{}", coin_change(vec![1, 2, 5], 10));
    println!("{}", change(50, vec![1, 2, 5]));
}

//518. 零钱兑换 II
pub fn change(amount: i32, coins: Vec<i32>) -> i32 {
    //组合总数
    let mut dp = vec![0;amount as usize + 1];
    dp[0] = 1;
    for i in 0..coins.len(){
        for j in coins[i]..=amount{
            dp[j as usize] += dp[j as usize- coins[i] as usize];
        }
    }
    dp[dp.len() - 1]
}

//322. 零钱兑换
pub fn coin_change(coins: Vec<i32>, amount: i32) -> i32 {
    //dp[i][j] ,对于coins[i]来说，组合成j钱的最小数量。
    let mut dp = vec![vec![-1; amount as usize + 1]; coins.len()];
    //dp[0][coins[0]] = 1;
    //只有一个硬币的情况
    dp[0][0] = 0;
    for i in 1..=amount as usize {
        if i >= coins[0] as usize && dp[0][i - coins[0] as usize] != -1 {
            dp[0][i] = dp[0][i - coins[0] as usize] + 1;
        }
    }
    //每次新增加一种硬币
    for i in 1..coins.len() {
        for j in 0..=amount as usize {
            //每次更新dp[i][j]
            if j == 0 {
                dp[i][0] = 0;
            } else {
                //用本轮的硬币，不用本轮的硬币两种情况
                let mut count = -1;
                if j >= coins[i] as usize && dp[i][j - coins[i] as usize] != -1 {
                    count = dp[i][j - coins[i] as usize] + 1;
                }
                if count == -1 {
                    count = dp[i - 1][j];
                } else if dp[i - 1][j] != -1 {
                    count = count.min(dp[i - 1][j]);
                }
                dp[i][j] = count;
            }
        }
    }
    dp[coins.len() - 1][amount as usize]
}

//1748. 唯一元素的和
pub fn sum_of_unique(nums: Vec<i32>) -> i32 {
    use std::collections::HashMap;
    let mut map = HashMap::new();
    for num in nums {
        let entry = map.entry(num).or_insert(0);
        *entry += 1;
    }
    map
        .into_iter()
        .filter(|(_, v)| *v < 2)
        .map(|e| e.0)
        .sum()
}
