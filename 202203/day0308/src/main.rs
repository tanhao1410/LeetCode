fn main() {
    println!("Hello, world!");
}

impl Solution {
    //2055. 蜡烛之间的盘子
    pub fn plates_between_candles(s: String, queries: Vec<Vec<i32>>) -> Vec<i32> {
        //前面盘子的数量的，下一个蜡烛的为主，上一个蜡烛的位置
        let mut dp = vec![(0, s.len() - 1, 0); s.len()];
        let bytes = s.as_bytes();
        if bytes[0] == b'*' {
            dp[0] = (1, 0, 0);
        }
        if bytes[bytes.len() - 1] == b'|' {
            dp[bytes.len() - 1] = (0, bytes.len() - 1, bytes.len() - 1);
        }
        for i in 1..s.len() {
            if bytes[i] == b'|' {
                dp[i] = (dp[i - 1].0, i, i);
            } else {
                dp[i].0 = dp[i - 1].0 + 1;
                dp[i].2 = dp[i - 1].2;
            }
            if bytes[s.len() - 1 - i] != b'|' {
                dp[s.len() - 1 - i].1 = dp[s.len() - i].1
            } else {
                dp[s.len() - 1 - i].1 = s.len() - 1 - i;
            }
        }
        //println!("{:?}",dp);
        queries
            .into_iter()
            .map(|v| {
                let start = v[0] as usize;
                let end = v[1] as usize;
                0.max(dp[dp[end].2].0 - dp[dp[start].1].0)
            })
            .collect()
    }
}

struct Solution;