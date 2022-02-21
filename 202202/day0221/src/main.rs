fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //1615. 最大网络秩
    pub fn maximal_network_rank(n: i32, roads: Vec<Vec<i32>>) -> i32 {
        //改善方法，不一定要两层循环。按每一个点能连接的数量进行排序，返回的结果中肯定是最大的两个点进行组合。
        //如果最大的两个点都是两两相连，如果最大的两个点不相连
        use std::collections::HashSet;
        let mut graph = vec![(0, 0); n as usize];
        let mut set = HashSet::new();
        for road in &roads {
            graph[road[0] as usize].1 += 1;
            graph[road[1] as usize].1 += 1;
            set.insert((road[0], road[1]));
        }
        for i in 0..n as usize {
            graph[i].0 = i as i32;
        }
        graph.sort_unstable_by_key(|e| (*e).1);
        let mut res = 0;
        //寻找最大的几个值
        let max_conn = graph[n as usize - 1].1;
        for i in (1..n as usize).rev() {
            //肯定能用到最大的连接数
            if graph[i].1 < max_conn {
                break;
            }
            //第二个最多用到一个比最大连接数小的数
            for j in (0..i) {
                let mut cur_count = max_conn + graph[j].1;
                //看两者是否相连
                if set.contains(&(graph[i].0, graph[j].0)) || set.contains(&(graph[j].0, graph[i].0)) {
                    cur_count -= 1;
                }
                res = res.max(cur_count);
            }
        }
        res
    }
}
