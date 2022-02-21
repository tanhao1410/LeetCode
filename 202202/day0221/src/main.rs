fn main() {
    println!("Hello, world!");
    println!("{}", Solution::possible_bipartition(4, vec![vec![1, 2], vec![1, 3], vec![2, 4]]));
}

struct Solution;

impl Solution {
    //785. 判断二分图
    pub fn is_bipartite(graph: Vec<Vec<i32>>) -> bool {
        use std::collections::HashSet;
        let mut used = vec![false; graph.len()];
        let mut stack1 = vec![];
        let mut stack2 = vec![];
        let mut set1 = HashSet::new();
        let mut set2 = HashSet::new();
        for i in 0..graph.len() {
            if !used[i] {
                stack1.push(i);
                set1.insert(i);
                while !stack1.is_empty() || !stack2.is_empty() {
                    while let Some(cur) = stack1.pop() {
                        for &next in &graph[cur] {
                            if !set2.contains(&(next as usize)) {
                                set2.insert(next as usize);
                                stack2.push(next as usize);
                                used[next as usize] = true;
                            }
                        }
                    }
                    while let Some(cur) = stack2.pop() {
                        for &next in &graph[cur] {
                            if !set1.contains(&(next as usize)) {
                                set1.insert(next as usize);
                                stack1.push(next as usize);
                                used[next as usize] = true;
                            }
                        }
                    }
                }
            }
        }
        set1.len() + set2.len() == graph.len()
    }

    //886. 可能的二分法
    pub fn possible_bipartition(n: i32, dislikes: Vec<Vec<i32>>) -> bool {
        //思路:1.用两个集合来分别遍历
        // 2.访问节点 i 时，将所有i不喜欢的放入到对方集合中
        // 3.同理，对方遍历时，将自己不喜欢的 放入到对方集合中
        // 4.最后，看两个集合是否有交集。如果有交集，返回false
        use std::collections::HashSet;
        let mut used = vec![false; n as usize];
        let mut map = vec![HashSet::new(); n as usize];
        for dislike in &dislikes {
            map[dislike[0] as usize - 1].insert(dislike[1] as usize - 1);
            map[dislike[1] as usize - 1].insert(dislike[0] as usize - 1);
        }
        //分成两个集合。
        let mut set1 = HashSet::new();
        let mut set2 = HashSet::new();

        //所有与set1内部有连接的都放入到set2中
        let mut queue1 = vec![];//里面放的都是应该放置到set2中的元素
        let mut queue2 = vec![];//里面放的都是应该放置到set1中的元素
        for i in 0..n as usize {
            if !used[i] {
                queue1.push(i as i32);
                used[i] = true;
                set1.insert(i as i32);
                while !queue2.is_empty() || !queue1.is_empty() {
                    while let Some(cur) = queue1.pop() {
                        for &next in &map[cur as usize] {
                            if !set2.contains(&(next as i32)) {
                                set2.insert(next as i32);
                                used[next] = true;
                                queue2.push(next as i32);
                            }
                        }
                    }
                    while let Some(cur) = queue2.pop() {
                        for &next in &map[cur as usize] {
                            if !set1.contains(&(next as i32)) {
                                set1.insert(next as i32);
                                used[next] = true;
                                queue1.push(next as i32);
                            }
                        }
                    }
                }
            }
        }
        //println!("{:?}",set1);
        //println!("{:?}",set2);
        set1.len() + set2.len() == n as usize
    }

    //763. 划分字母区间
    pub fn partition_labels(s: String) -> Vec<i32> {
        let mut last_indexs = vec![0; 26];
        let mut res = vec![];
        let bytes = s.as_bytes();
        for i in 0..bytes.len() {
            last_indexs[(bytes[i] - b'a') as usize] = i;
        }
        let mut last_index = 0;
        let mut pre_index = -1;
        for i in 0..bytes.len() {
            last_index = last_index.max(last_indexs[(bytes[i] - b'a') as usize]);
            if last_index == i {
                res.push(i as i32 - pre_index);
                pre_index = i as i32;
            }
        }
        res
    }
    //838. 推多米诺
    pub fn push_dominoes(dominoes: String) -> String {
        let mut bytes = dominoes.into_bytes();
        //从开始处往后找，直到找到一个L或R ，如果是L，则从开始处到L都要变成L否则，
        //从当前R处开始往后找，直到 找到一个L或R ，如果是L ，则开始与结束之间的要改变。如果是R ，则从开始到结束变成R
        let mut start = 0;
        while start < bytes.len() {
            let status = bytes[start] != b'R';
            let mut end = start + 1;
            while end < bytes.len() && bytes[end] == b'.' {
                end += 1;
            }
            if end == bytes.len() {
                //走到了终点
                if !status {
                    for i in start + 1..end {
                        bytes[i] = b'R';
                    }
                }
                break;
            }
            match bytes[end] {
                b'L' => {
                    if status {
                        //从开始到结束之间的都变成L
                        for i in start..end {
                            bytes[i] = b'L';
                        }
                    } else {
                        //L...R
                        for i in 1..=(end - start - 1) / 2 {
                            bytes[i + start] = b'R';
                            bytes[end - i] = b'L';
                        }
                    }
                }
                _ => {
                    if !status {
                        for i in start..end {
                            bytes[i] = b'R';
                        }
                    }
                }
            }
            start = end;
        }
        String::from_utf8(bytes).unwrap()
    }
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
