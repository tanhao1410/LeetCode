use std::collections::HashMap;

fn main() {
    println!("Hello, world!");
}

//567. 字符串的排列
pub fn check_inclusion(s1: String, s2: String) -> bool {
    let mut s1_map = vec![0; 26];
    s1
        .as_bytes()
        .iter()
        .for_each(|&l| s1_map[(l - b'a') as usize] += 1);
    //寻找一个大小相等的窗口
    if s2.len() < s1.len() {
        return false;
    }
    let mut s2_map = vec![0; 26];
    s2
        .as_bytes()
        .iter()
        .take(s1.len())
        .for_each(|&l| s2_map[(l - b'a') as usize] += 1);
    let is_equal = |s2_map: &[i32]| {
        for i in 0..26 {
            if s2_map[i] != s1_map[i] {
                return false;
            }
        }
        true
    };
    //先判断
    if is_equal(&s2_map) {
        return true;
    }
    let s2_bytes = s2.as_bytes();
    for i in s1.len()..s2.len() {
        //加进来一个，减去一个
        let add = s2_bytes[i] - b'a';
        let remove = s2_bytes[i - s1.len()] - b'a';
        s2_map[add as usize] += 1;
        s2_map[remove as usize] -= 1;
        if is_equal(&s2_map) {
            return true;
        }
    }
    false
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