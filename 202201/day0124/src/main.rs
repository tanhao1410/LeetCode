fn main() {
    println!("Hello, world!");
}

//121. 买卖股票的最佳时机
pub fn max_profit(prices: Vec<i32>) -> i32 {
    prices
        .iter()
        .fold((i32::MAX, 0), |(min, res), &cur| (min.min(cur), res.max(cur - min)))
        .1

    //思路：dp记录前面最便宜的价格
    // let mut dp = prices[0];
    // let mut res = 0;
    // for i in 1..prices.len(){
    //     res = res.max(prices[i] - dp);
    //     dp = dp.min(prices[i]);
    // }
    // res
}

//1014. 最佳观光组合
pub fn max_score_sightseeing_pair(values: Vec<i32>) -> i32 {
    //思路:dp[i]以values[i]结尾的组合，最大的值
    //dp[i] = dp[i-1] - 1 - values[i - 1] + values[i] 或 values[i] + values[i -1] - 1
    let mut dp = values[0] + values[1] - 1;
    let mut res = dp;
    for i in 2..values.len() {
        if dp > 2 * values[i - 1] {
            dp = dp - 1 - values[i - 1] + values[i];
        } else {
            dp = values[i - 1] - 1 + values[i];
        }
        res = res.max(dp);
    }
    res
}

//557. 反转字符串中的单词 III
pub fn reverse_words(s: String) -> String {
    s
        .split(" ")
        .map(|s| s.chars().rev().collect::<String>())
        .collect::<Vec<_>>()
        .join(" ")
}

//519. 随机翻转矩阵
pub fn reverse_string(s: &mut Vec<char>) {
    let mut i = 0;
    let mut j = s.len() - 1;
    while j > i {
        let temp = s[i];
        s[i] = s[j];
        s[j] = temp;
        i += 1;
        j -= 1;
    }
}