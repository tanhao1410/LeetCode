fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //433. 最小基因变化
    pub fn min_mutation(start: String, end: String, bank: Vec<String>) -> i32 {
        //广度优先策略？每一次改变一次。遍历过的不再遍历了
        use std::collections::{HashSet, VecDeque};
        let mut map = HashSet::new();
        let is_one_change = |one: &String, two: &String| {
            one
                .as_bytes()
                .iter()
                .zip(two.as_bytes().iter())
                .filter(|(i, j)| **i != **j)
                .count()
                == 1
        };
        map.insert(&start);
        let mut queue = VecDeque::new();
        queue.push_back(&start);
        let mut layer = 0;
        while !queue.is_empty() {
            //突变产生了end
            if map.contains(&end) {
                return layer;
            }
            let mut len = queue.len();
            for _ in 0..len {
                let cur = queue.pop_front().unwrap();
                //测试cur能突变成哪一个
                for s in &bank {
                    if !map.contains(s) && is_one_change(s, cur) {
                        queue.push_back(s);
                        map.insert(s);
                    }
                }
            }
            layer += 1;
        }
        -1
    }
    //969. 煎饼排序
    pub fn pancake_sort(mut arr: Vec<i32>) -> Vec<i32> {
        let mut res = vec![];
        let reverse_arr = |arr: &mut Vec<i32>, mut start: usize, mut end: usize| {
            while start < end {
                let temp = arr[start];
                arr[start] = arr[end];
                arr[end] = temp;
                start += 1;
                end -= 1;
            }
        };
        for i in (1..=arr.len()).rev() {
            if arr[i - 1] != i as i32 {
                //需要翻转
                for j in 0..i {
                    if arr[j] == i as i32 {
                        if j != 0 {
                            //先翻转一次，将这个数翻转到开头
                            res.push(j as i32 + 1);
                            reverse_arr(&mut arr, 0, j);
                        }
                        res.push(i as i32);
                        reverse_arr(&mut arr, 0, i - 1);
                        break;
                    }
                }
            }
        }
        res
    }
}