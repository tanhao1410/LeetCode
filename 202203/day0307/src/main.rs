fn main() {
    println!("Hello, world!");
}

impl Solution {
    //剑指 Offer II 012. 左右两边子数组的和相等
    pub fn pivot_index(nums: Vec<i32>) -> i32 {
        //先求总和，然后从零开始依次求和sum，若sum = (sum2 - cur )/2。返回结果
        let sum = nums.iter().sum::<i32>();
        let mut s = 0;
        for i in 0..nums.len() {
            let cur = nums[i];
            if s * 2 == sum - cur {
                return i as i32;
            }
            s += cur;
        }
        -1
    }
    //剑指 Offer II 038. 每日温度
    pub fn daily_temperatures(temperatures: Vec<i32>) -> Vec<i32> {
        let mut stack: Vec<(i32, usize)> = vec![];
        let mut res = vec![0; temperatures.len()];
        for i in (0..temperatures.len()).rev() {
            let cur = temperatures[i];
            while !stack.is_empty() && stack[stack.len() - 1].0 <= cur {
                stack.pop();
            }
            if !stack.is_empty() {
                res[i] = stack[stack.len() - 1].0 as i32 - i as i32;
            }
            stack.push((cur, i));
        }
        res
    }
    //剑指 Offer II 035. 最小时间差
    pub fn find_min_difference(time_points: Vec<String>) -> i32 {
        //每一分钟做个对应的话，也不过24*60个。需要注意的是存在循环。
        use std::collections::HashSet;
        let mut set = HashSet::new();
        let time_to_int = |time: &str| {
            let mut res = 0;
            let mut split = time.split(":");
            res += split.next().unwrap().parse::<i32>().unwrap() * 60;
            res += split.next().unwrap().parse::<i32>().unwrap();
            res
        };
        for time in time_points {
            let time = time_to_int(&time);
            if set.contains(&time) {
                return 0;
            }
            set.insert(time);
        }
        let mut times = set.into_iter().collect::<Vec<_>>();
        times.sort_unstable();
        let mut min = times[0] + 1440 - times[times.len() - 1];
        for i in 1..times.len() {
            min = min.min(times[i] - times[i - 1]);
        }
        min
    }
    //剑指 Offer II 036. 后缀表达式
    pub fn eval_rpn(tokens: Vec<String>) -> i32 {
        use std::ops::Add;
        use std::ops::Mul;
        use std::ops::Div;
        use std::ops::Sub;
        let mut stack = vec![];
        let opera = |stack: &mut Vec<i32>, op: fn(i32, i32) -> i32| {
            let mut a = stack.pop().unwrap();
            let mut b = stack.pop().unwrap();
            stack.push(op(b, a));
        };
        for token in tokens {
            //如果是数就进栈
            match token.as_str() {
                "+" => opera(&mut stack, i32::add),
                "-" => opera(&mut stack, i32::sub),
                "*" => opera(&mut stack, i32::mul),
                "/" => opera(&mut stack, i32::div),
                _ => stack.push(token.parse().unwrap())
            }
        }
        stack.pop().unwrap()
    }
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