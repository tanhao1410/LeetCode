fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //56. 合并区间
    pub fn merge(mut intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        intervals.sort_by_key(|e| e[0]);
        let mut res = vec![];
        let mut pre_interval = (intervals[0][0], intervals[0][1]);
        for i in 1..intervals.len() {
            let cur_interval = (intervals[i][0], intervals[i][1]);
            //合并当前区间与前一个区间
            if cur_interval.0 > pre_interval.1 {
                //两个区间不相交
                res.push(vec![pre_interval.0, pre_interval.1]);
                pre_interval = cur_interval;
            } else if cur_interval.1 > pre_interval.1 {
                //合并
                pre_interval = (pre_interval.0, cur_interval.1)
            }
        }
        //合并最后的区间
        res.push(vec![pre_interval.0,pre_interval.1]);
        res
    }

    //22. 括号生成
    pub fn generate_parenthesis(n: i32) -> Vec<String> {
        Self::generate_parenthesis2(n, n, String::new())
    }

    fn generate_parenthesis2(left: i32, right: i32, mut pre: String) -> Vec<String> {
        if left == 0 {
            for _ in 0..right {
                pre.push(')');
            }
            return vec![pre];
        }
        let mut res = vec![];
        if left < right {
            let mut new_pre = pre.clone();
            new_pre.push(')');
            res.append(&mut Self::generate_parenthesis2(left, right - 1, new_pre));
        }
        pre.push('(');
        res.append(&mut Self::generate_parenthesis2(left - 1, right, pre));
        res
    }

    //17. 电话号码的字母组合
    pub fn letter_combinations(digits: String) -> Vec<String> {
        // 2-9
        if digits.len() == 0 {
            return vec![];
        }
        let mut map = vec!["abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"];
        let mut res = vec!["".to_string()];
        for digit in digits.chars() {
            let bytes = map[(digit as u8 - b'2') as usize].as_bytes();
            //扩大chars.len()倍数
            let count = res.len();
            for i in 0..bytes.len() {
                if i == 0 {
                    for j in 0..count {
                        res[j].push(bytes[i] as char);
                    }
                } else {
                    for j in 0..count {
                        let mut new_str = res[j][..res[j].len() - 1].to_string();
                        //此时clone出现了问题，因为前面已经被更改过了。
                        new_str.push(bytes[i] as char);
                        res.push(new_str);
                    }
                }
            }
        }
        res
    }
    //365. 水壶问题
    pub fn can_measure_water(jug1_capacity: i32, jug2_capacity: i32, target_capacity: i32) -> bool {
        //广度优先策略
        use std::collections::HashSet;
        let mut read = HashSet::new();
        read.insert((0, 0));
        let mut stack = vec![(0, 0)];
        while let Some((jug1, jug2)) = stack.pop() {
            if jug1 == target_capacity || jug2 == target_capacity || jug2 + jug1 == target_capacity {
                return true;
            }
            //8种操作
            let operators = vec![(jug1_capacity, jug2), (jug1, jug2_capacity), (0, jug2), (jug1, 0),
                                 (0, jug1 + jug2), (jug1 - jug2_capacity + jug2, jug2_capacity), (jug1 + jug2, 0), (jug1_capacity, jug1 + jug2 - jug1_capacity)];
            for state in operators {
                if state.0 >= 0 && state.0 <= jug1_capacity && state.1 >= 0 && state.1 <= jug2_capacity
                    && !read.contains(&state) {
                    stack.push(state);
                    read.insert(state);
                }
            }
        }
        false
    }
    //1306. 跳跃游戏 III
    pub fn can_reach(arr: Vec<i32>, start: i32) -> bool {
        //深度优先策略
        let mut read = vec![false; arr.len()];
        let mut stack = vec![start as usize];
        read[start as usize] = true;
        while let Some(i) = stack.pop() {
            if arr[i] == 0 {
                return true;
            }
            //两边走
            let dirs = vec![i as i32 + arr[i], i as i32 - arr[i]];
            for dir in dirs {
                if dir >= 0 && dir < arr.len() as i32 && !read[dir as usize] {
                    stack.push(dir as usize);
                    read[dir as usize] = true;
                }
            }
        }
        false
    }

    //802. 找到最终的安全状态
    pub fn eventual_safe_nodes(graph: Vec<Vec<i32>>) -> Vec<i32> {
        //只能通往安全节点的节点是安全节点。
        //最初的安全节点是没有通往下一个结点。
        // 2为通往自己的节点们
        let mut graph2 = vec![vec![]; graph.len()];
        let mut stack = vec![];
        let mut safe_nodes = vec![false; graph.len()];
        for i in 0..graph.len() {
            //i 通向 j
            for &j in &graph[i] {
                graph2[j as usize].push(i);
            }
            if graph[i].len() == 0 {
                stack.push(i);
                safe_nodes[i] = true;
            }
        }

        while let Some(i) = stack.pop() {
            //看通向它的节点是否都是 都是通向 安全节点。
            for &node in &graph2[i] {
                if graph[node as usize]
                    .iter()
                    .all(|e| safe_nodes[*e as usize]) {
                    stack.push(node as usize);
                    safe_nodes[node as usize] = true;
                }
            }
        }

        safe_nodes
            .into_iter()
            .enumerate()
            .filter_map(|e| match e.1 {
                true => Some(e.0 as i32),
                false => None
            })
            .collect()
    }

    //240. 搜索二维矩阵 II
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        let mut x = 0i32;
        let mut y = matrix[0].len() as i32 - 1;
        while x < matrix.len() as i32 && y >= 0 {
            if matrix[x as usize][y as usize] == target {
                return true;
            }
            if matrix[x as usize][y as usize] > target {
                y -= 1;
            } else {
                x += 1;
            }
        }
        false
    }
}