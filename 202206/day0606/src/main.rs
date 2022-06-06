fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //2094. 找出 3 位偶数
    pub fn find_even_numbers(digits: Vec<i32>) -> Vec<i32> {
        let counts = digits.into_iter().fold(vec![0; 10], |mut counts, n| {
            counts[n as usize] += 1;
            counts
        });
        let mut contains = |mut num: i32| {
            let mut num_counts = vec![0; 10];
            while num > 0 {
                num_counts[num as usize % 10] += 1;
                num /= 10;
            }
            num_counts.iter().zip(counts.iter()).all(|e| e.0 <= e.1)
        };

        (100..999)
            .step_by(2)
            .filter(|n| contains(*n))
            .collect()
    }
}

use std::collections::BTreeMap;

//732. 我的日程安排表 III
struct MyCalendarThree {
    items: BTreeMap<i32, i32>
}

impl MyCalendarThree {
    fn new() -> Self {
        Self { items: BTreeMap::new() }
    }
    fn book(&mut self, start: i32, end: i32) -> i32 {
        *self.items.entry(start).or_insert(0) += 1;
        *self.items.entry(end).or_insert(0) += 1;
        //每一次都加上该值，求最大的情况即答案
        self.items
            .values()
            .fold((0, 0), |(res, mut sum), i| {
                sum += *i;
                (res.max(sum), sum)
            }).0
    }
}