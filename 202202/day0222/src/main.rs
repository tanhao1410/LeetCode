fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //剑指 Offer II 008. 和大于等于 target 的最短子数组
    pub fn min_sub_array_len(target: i32, nums: Vec<i32>) -> i32 {
        let mut res = i32::MAX;
        let mut start = 0;
        let mut end = 0;
        let mut sum = nums[start];
        while end < nums.len() {
            //end往前走
            while end < nums.len() && sum < target {
                end += 1;
                if end < nums.len() {
                    sum += nums[end];
                }
            }
            if sum >= target {
                res = res.min((end - start) as i32 + 1);
            }
            //start 往前走
            while start <= end && sum >= target {
                sum -= nums[start];
                start += 1;
                if sum >= target {
                    res = res.min((end - start) as i32 + 1);
                }
            }
        }
        if res == i32::MAX { 0 } else { res }
    }
    //剑指 Offer II 007. 数组中和为 0 的三个数
    pub fn three_sum(mut nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut res = vec![];
        if nums.len() < 3 {
            return res;
        }
        nums.sort_unstable();
        for i in 0..nums.len() - 2 {
            //双指针法，从剩下的区间中找到目标 -nums[i]
            if i > 0 && nums[i] == nums[i - 1] {
                continue;
            }
            let mut start = i + 1;
            let mut end = nums.len() - 1;
            while end > start {
                if nums[end] + nums[start] == -nums[i] {
                    res.push(vec![nums[i], nums[start], nums[end]]);
                    let mut next_start = start + 1;
                    while next_start < end && nums[next_start] == nums[start] {
                        next_start += 1;
                    }
                    start = next_start;
                    end -= 1;
                } else if nums[end] + nums[start] > -nums[i] {
                    end -= 1;
                } else {
                    start += 1;
                }
            }
        }
        res
    }
    //70. 爬楼梯
    pub fn climb_stairs(n: i32) -> i32 {
        //f(i) = f(i - 1) + f(i - 2)
        let mut pre = 1;
        let mut pre_pre = 0;
        for _ in 1..=n {
            pre += pre_pre;
            pre_pre = pre - pre_pre;
        }
        pre
    }
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