fn main() {
    println!("Hello, world!");
}

impl Solution {
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