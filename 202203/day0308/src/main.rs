fn main() {
    println!("Hello, world!");
}

impl Solution {
    //剑指 Offer II 032. 有效的变位词
    pub fn is_anagram(s: String, t: String) -> bool {
        if s == t {
            return false;
        }
        let mut v = vec![0; 26];
        let bytes = s.as_bytes();
        for &b in bytes {
            v[(b - 'a') as usize] += 1;
        }
        let bytes = t.as_bytes();
        for &b in bytes {
            v[(b - b'a') as usize] -= 1;
        }
        v.into_iter().all(|e| e == 0)
    }
    //剑指 Offer II 018. 有效的回文
    pub fn is_palindrome(s: String) -> bool {
        let mut bytes = s.into_bytes();
        let mut i = 0;
        let mut j = bytes.len() - 1;
        while j > i {
            if !((bytes[i] >= b'a' && bytes[i] <= b'z') || (bytes[i] >= b'A' && bytes[i] <= b'Z') || (bytes[i] >= b'0' && bytes[i] <= b'9')) {
                i += 1;
                continue;
            }
            if !((bytes[j] >= b'a' && bytes[j] <= b'z') || (bytes[j] >= b'A' && bytes[j] <= b'Z') || (bytes[j] >= b'0' && bytes[j] <= b'9')) {
                j -= 1;
                continue;
            }
            if bytes[i] >= b'A' && bytes[i] <= b'Z' {
                bytes[i] += b'a' - b'A';
            }
            if bytes[j] >= b'A' && bytes[j] <= b'Z' {
                bytes[j] += b'a' - b'A';
            }
            if bytes[i] != bytes[j] {
                return false;
            }
            i += 1;
            j -= 1;
        }
        true
    }
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