fn main() {
    println!("Hello, world!");
}

impl Solution {
    //451. 根据字符出现频率排序
    pub fn frequency_sort(s: String) -> String {
        use std::collections::HashMap;
        let mut map = HashMap::new();
        s.chars().for_each(|c| *map.entry(c).or_insert(0) += 1);
        let mut v = map.into_iter().collect::<Vec<_>>();
        v.sort_unstable_by_key(|&e| -e.1);
        let mut res = String::new();
        for (c, count) in v {
            for _ in 0..count {
                res.push(c);
            }
        }
        res
    }
}

struct Solution;