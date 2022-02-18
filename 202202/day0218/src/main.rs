fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
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