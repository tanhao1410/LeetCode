fn main() {
    println!("Hello, world!");
}

impl Solution {
    //139. 单词拆分
    pub fn word_break(s: String, word_dict: Vec<String>) -> bool {
        use std::collections::HashSet;
        let set = word_dict.iter().map(|e| e.as_str()).collect::<HashSet<_>>();
        let mut dp = vec![false; s.len()];
        for i in 0..dp.len() {
            if set.contains(&s[..=i]) {
                dp[i] = true;
            } else {
                for j in 0..i {
                    if dp[j] && set.contains(&s[j + 1..=i]) {
                        dp[i] = true;
                        break;
                    }
                }
            }
        }
        dp[dp.len() - 1]
    }
}