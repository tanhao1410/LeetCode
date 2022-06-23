fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //1387. 将整数按权重排序
    pub fn get_kth(lo: i32, hi: i32, k: i32) -> i32 {
        use std::collections::HashMap;
        let map = (0..31).map(|i| (1 << i, i)).collect::<HashMap<i32, i32>>();
        let become_one_count = |num: &i32| -> i32{
            let mut num = *num;
            let mut res = 0;
            while !map.contains_key(&num) {
                if num % 2 == 0 {
                    num /= 2;
                } else {
                    num = num * 3 + 1;
                }
                res += 1;
            }
            res + map.get(&num).unwrap()
        };
        let mut nums = (lo..=hi).collect::<Vec<i32>>();
        nums.sort_by_key(become_one_count);
        nums[k as usize - 1]
    }
}