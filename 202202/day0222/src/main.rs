fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //509. 斐波那契数
    pub fn fib(n: i32) -> i32 {
        if n < 2 {
            return n;
        }
        let mut pre = 1;
        let mut pre_pre = 0;
        for i in 2..=n {
            pre += pre_pre;
            pre_pre = pre - pre_pre;
        }
        pre
    }
    //49. 字母异位词分组
    pub fn group_anagrams(mut strs: Vec<String>) -> Vec<Vec<String>> {
        use std::collections::HashMap;
        let mut transfer = |word: &String| {
            let chars = word.as_bytes();
            let mut counts = vec![0; 26];
            for &c in chars {
                counts[(c - b'a') as usize] += 1;
            }
            let mut res = String::new();
            for i in 0..26 {
                for _ in 0..counts[i] {
                    res.push((b'a' + i as u8) as char);
                }
            }
            res
        };
        let mut map: HashMap<String, Vec<String>> = HashMap::new();
        while let Some(word) = strs.pop() {
            let word2 = transfer(&word);
            if let Some(v) = map.get_mut(&word2) {
                v.push(word)
            } else {
                map.insert(word2, vec![]);
            }
        }
        map.into_iter().map(|entry| entry.1).collect()
    }

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