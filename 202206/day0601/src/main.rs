#![feature(is_sorted)]

fn main() {
    println!("Hello, world!");
    println!("{}", Solution::has_groups_size_x(vec![1, 1, 1, 2, 2, 2, 3, 3]));
}

struct Solution;

impl Solution {
    //914. 卡牌分组
    pub fn has_groups_size_x(deck: Vec<i32>) -> bool {
        use std::collections::HashMap;
        let mut map = HashMap::new();
        for num in deck {
            *map.entry(num).or_insert(0) += 1;
        }
        for i in 2.. {
            if map.values().all(|&v| v % i == 0) {
                return true;
            }
            if map.values().any(|&v| i > v) {
                return false;
            }
        }
        unreachable!()
    }

    //953. 验证外星语词典
    pub fn is_alien_sorted(words: Vec<String>, order: String) -> bool {
        use std::collections::HashMap;
        let map = order.into_bytes().into_iter().zip(b'a'..=b'z').collect::<HashMap<_, _>>();
        //每一个字母对应的是什么呢？
        let mut new_words = words
            .into_iter()
            .map(|word| String::from_utf8(word
                .as_bytes()
                .into_iter()
                .map(|l| *map.get(l).unwrap())
                .collect::<Vec<_>>()
            ).unwrap())
            .collect::<Vec<_>>();
        // for i in 1..new_words.len() {
        //     if new_words[i] < new_words[i - 1] {
        //         return false;
        //     }
        // }
        // true
        new_words.is_sorted()
    }
    //1071. 字符串的最大公因子
    pub fn gcd_of_strings(str1: String, str2: String) -> String {
        let is_child_str = |parent: &str, child: &str| {
            if parent.len() % child.len() != 0 {
                return false;
            }
            for i in 0..parent.len() / child.len() {
                if &parent[i * child.len()..(i + 1) * child.len()] != child {
                    return false;
                }
            }
            true
        };
        for i in 1..str1.len() + 1 {
            if str1.len() % i == 0
                && is_child_str(str1.as_str(), &str1[..str1.len() / i])
                && is_child_str(str2.as_str(), &str1[..str1.len() / i]) {
                return str1[..str1.len() / i].to_string();
            }
        }
        "".to_string()
    }
    //473. 火柴拼正方形
    pub fn makesquare(matchsticks: Vec<i32>) -> bool {
        let all_len = matchsticks.iter().sum::<i32>();
        if all_len % 4 == 0 && matchsticks.iter().all(|&l| l <= all_len / 4) {
            return Self::select_next(matchsticks.as_slice(), all_len / 4, &mut vec![0; 4], 0);
        }
        false
    }
    fn select_next(matchsticks: &[i32], res_len: i32, edges: &mut Vec<i32>, cur: usize) -> bool {
        if cur == matchsticks.len() {
            return edges.iter().all(|&l| l == res_len);
        }
        for i in 0..4 {
            if edges[i] + matchsticks[cur] <= res_len {
                edges[i] += matchsticks[cur];
                let inner_res = Self::select_next(matchsticks, res_len, edges, cur + 1);
                if inner_res {
                    return true;
                }
                edges[i] -= matchsticks[cur]
            }
        }
        false
    }
}