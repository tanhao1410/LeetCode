fn main() {
    println!("Hello, world!");
    println!("{}", Solution::add_strings("9999".to_string(), "1".to_string()));
}

struct Solution;

impl Solution {
    //1129. 颜色交替的最短路径
    pub fn shortest_alternating_paths(n: i32, red_edges: Vec<Vec<i32>>, blue_edges: Vec<Vec<i32>>) -> Vec<i32> {
        //广度优先策略；
        //走法，本次走了红边，下次就只能走绿边
        use std::collections::{VecDeque,HashSet};
        let mut used = HashSet::new();
        let mut queue = VecDeque::new();
        let mut res = vec![-1; n as usize];
        let mut step = 0;
        let mut red_set = vec![vec![]; n as usize];
        let mut blue_set = vec![vec![]; n as usize];
        for edge in red_edges {
            red_set[edge[0] as usize].push(edge[1] as usize);
        }
        for edge in blue_edges {
            blue_set[edge[0] as usize].push(edge[1] as usize);
        }
        queue.push_back((0, true));
        used.insert((0, true));
        used.insert((0, false));
        queue.push_back((0, false));
        while !queue.is_empty() {
            let len = queue.len();
            for _ in 0..len {
                let (cur, is_red) = queue.pop_front().unwrap();
                if res[cur as usize] == -1 {
                    res[cur as usize] = step;
                }
                //只能走绿色或红色
                for &next in if is_red { &blue_set[cur as usize] } else { &red_set[cur as usize] } {
                    if !used.contains(&(next, !is_red)) {
                        queue.push_back((next, !is_red));
                        used.insert((next, !is_red));
                    }
                }
            }
            step += 1;
        }
        res
    }

    //127. 单词接龙
    pub fn ladder_length(begin_word: String, end_word: String, word_list: Vec<String>) -> i32 {
        use std::collections::{HashMap, VecDeque};
        let mut word_map = word_list
            .iter()
            .map(|e| (e, false))
            .collect::<HashMap<&String, bool>>();
        let mut queue = VecDeque::new();
        let change_one = |s1: &String, s2: &String| {
            s1.as_bytes().iter().zip(s2.as_bytes().iter()).filter(|&e| e.0 != e.1).count() == 1
        };
        queue.push_back(&begin_word);
        word_map.insert(&begin_word, true);
        let mut step = 1;
        while !queue.is_empty() {
            let len = queue.len();
            if *word_map.get(&end_word).unwrap_or(&false) {
                return step;
            }
            for _ in 0..len {
                let cur = queue.pop_front().unwrap();
                for word in &word_list {
                    //没有被使用过，且差别只有一个字母
                    if !*word_map.get(word).unwrap() && change_one(word, cur) {
                        queue.push_back(word);
                        word_map.insert(word, true);
                    }
                }
            }
            step += 1;
        }
        0
    }
    //415. 字符串相加
    pub fn add_strings(num1: String, num2: String) -> String {
        let mut res = vec![];
        let mut num1 = num1.as_bytes();
        let mut num2 = num2.as_bytes();
        let mut index1 = num1.len() as i32 - 1;
        let mut index2 = num2.len() as i32 - 1;
        let mut flag = 0;
        while index1 >= 0 || index2 >= 0 {
            let mut bit_res = num1.get(index1 as usize).unwrap_or(&b'0') - b'0'
                + num2.get(index2 as usize).unwrap_or(&b'0') - b'0' + flag;
            flag = bit_res / 10;
            bit_res %= 10;
            res.push(bit_res + b'0');
            index1 -= 1;
            index2 -= 1;
        }
        if flag == 1 {
            res.push(b'1');
        }
        String::from_utf8(res.into_iter().rev().collect()).unwrap()
    }
    //1557. 可以到达所有点的最少点数目
    pub fn find_smallest_set_of_vertices(n: i32, edges: Vec<Vec<i32>>) -> Vec<i32> {
        let mut can_reach = vec![false; n as usize];
        for edge in &edges {
            can_reach[edge[1] as usize] = true;
        }
        can_reach
            .into_iter()
            .enumerate()
            .filter_map(|e| match e.1 {
                false => Some(e.0 as i32),
                _ => None
            })
            .collect()
    }
    //997. 找到小镇的法官
    pub fn find_judge(n: i32, trust: Vec<Vec<i32>>) -> i32 {
        let mut map = vec![(0, 0); n as usize];
        for v in trust {
            map[v[0] as usize - 1].0 += 1;
            map[v[1] as usize - 1].1 += 1;
        }
        for i in 0..n as usize {
            if map[i].0 == 0 && map[i].1 == n - 1 {
                return i as i32 + 1;
            }
        }
        return -1;
    }
}