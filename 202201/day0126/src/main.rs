use std::collections::HashMap;

fn main() {
    println!("Hello, world!");
}

//3. 无重复字符的最长子串
pub fn length_of_longest_substring(s: String) -> i32 {
    //思路：dp[i]以s[i]结尾的最长无重复子串。
    //dp[i] = if map.get[s[i]] isnone i + 1 or i - map.get[s[i]] .min(dp[i - 1] + 1)
    let mut dp = vec![1; s.len()];
    let mut map = HashMap::new();
    let bytes = s.as_bytes();
    for i in 0..s.len() {
        let cur_letter = bytes[i];
        //得到它的前一个位置
        if let Some(l) = map.get(&cur_letter) {
            dp[i] = i - l;
        } else {
            dp[i] = i + 1;
        }
        if i > 0 {
            dp[i] = dp[i].min(dp[i - 1] + 1);
        }
        //更新该字母位置
        map.insert(cur_letter, i);
    }
    dp
        .into_iter()
        .max()
        .unwrap_or(0) as i32
}