fn main() {
    println!("Hello, world!");
}

//122. 买卖股票的最佳时机 II
pub fn max_profit2(prices: Vec<i32>) -> i32 {
    //思路:什么时候买，买一个后面有比它大的，什么时候卖：已经买了，比买入贵，后面又降价了。
    let mut res = 0;
    let mut buy = false;
    let mut buy_num = 0;
    for i in 0..prices.len() - 1 {
        if !buy && prices[i + 1] > prices[i] {
            //如果还没买，可以买
            buy = true;
            buy_num = prices[i];
        }
        //什么时候可以卖呢，只有买了才能卖，只有，第二天降价了可以卖
        if buy && prices[i + 1] < prices[i] {
            buy = false;
            res += prices[i] - buy_num;
        }
    }
    if buy {
        res += prices[prices.len() - 1] - buy_num;
    }
    res
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