fn main() {
    println!("Hello, world!");
    println!("答案：{}", Solution::open_lock(vec!["0201".to_string(),"0101".to_string(),"0102".to_string(),"1212".to_string(),"2002".to_string()], "0202".to_string()));
}

struct Solution;

impl Solution {
    //752. 打开转盘锁
    pub fn open_lock(deadends: Vec<String>, target: String) -> i32 {
        use std::collections::{HashSet, VecDeque};
        //广度优先算法，每次共有8种走法。0000 用数字来表示 ，走法变成 +-1 +-10 不带进位的加减法 (00 + 100 - 10)%100
        let set = deadends.into_iter().map(|s| s.parse::<i32>().unwrap()).collect::<HashSet<i32>>();
        let mut used = HashSet::new();
        if set.contains(&0) {
            return -1;
        }
        let mut queue = VecDeque::new();
        let mut step = 0;
        queue.push_back(0);
        used.insert(0);
        while !queue.is_empty() {
            if used.contains(&target.parse().unwrap()) {
                return step;
            }
            let len = queue.len();
            for _ in 0..len {
                let cur = queue.pop_front().unwrap();
                print!("{},",cur);
                // 8种走法
                for i in 0..4 {
                    // 9 + 10 + 1
                    for j in vec![1, -1] {
                        // 1200 + 10 + 1  1201
                        let k = 10i32.pow(i);
                        // 1201 =>1202
                        //得到中间的某位，得到
                        let num = (cur / k) % 10;
                        let next_num = (num + 10 + j) % 10;
                        //需要将原来的数的该位替换成它
                        let next = cur - num * k + next_num * k;
                        if !set.contains(&next) && !used.contains(&next){
                            queue.push_back(next);
                            used.insert(next);
                        }
                    }
                }
            }
            println!();
            step += 1;
        }
        -1
    }

    //1654. 到家的最少跳跃次数
    pub fn minimum_jumps(forbidden: Vec<i32>, a: i32, b: i32, x: i32) -> i32 {
        use std::collections::{HashSet, VecDeque};
        //解答：最远禁止点
        let max_distance = *forbidden.iter().max().unwrap() + a + b;
        let max_distance = max_distance.max(x + b);
        //(index,是前进过来的吗)
        let mut forbidden_set = forbidden.into_iter().collect::<HashSet<i32>>();
        let mut used = HashSet::new();
        let mut queue = VecDeque::new();
        queue.push_back((0, true));
        used.insert((0, true));
        let mut step = 0;
        while !queue.is_empty() {
            if used.contains(&(x, true)) || used.contains(&(x, false)) {
                return step;
            }
            step += 1;
            let len = queue.len();
            for _ in 0..len {
                let (cur, flag) = queue.pop_front().unwrap();
                //往前走
                if cur <= max_distance && !forbidden_set.contains(&(cur + a)) && !used.contains(&(cur + a, true)) {
                    used.insert((cur + a, true));
                    queue.push_back((cur + a, true));
                }
                //往后走
                if flag && cur - b > 0 && !forbidden_set.contains(&(cur - b)) && !used.contains(&(cur - b, false)) {
                    used.insert((cur - b, false));
                    queue.push_back((cur - b, false));
                }
            }
        }
        -1
    }

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