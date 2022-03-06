fn main() {
    println!("Hello, world!");
}

impl Solution {
    //347. 前 K 个高频元素
    pub fn top_k_frequent(nums: Vec<i32>, k: i32) -> Vec<i32> {
        use std::collections::HashMap;
        let mut map = HashMap::new();
        for &num in &nums {
            let entry = map.entry(num).or_insert(0);
            *entry += 1;
        }
        let mut nums = map.into_iter().collect::<Vec<_>>();
        nums.sort_unstable_by_key(|&e| -e.1);
        nums.into_iter().take(k as usize).map(|e| e.0).collect()
    }
}

struct Solution;