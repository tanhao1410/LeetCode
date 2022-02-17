fn main() {
    println!("Hello, world!");
}

//706. 设计哈希映射
struct MyHashMap {
    nums: Vec<Vec<(i32, i32)>>,
    len: usize,
}

impl MyHashMap {
    fn new() -> Self {
        Self { nums: vec![vec![]; 1], len: 0 }
    }

    fn resize(&mut self) {
        if self.len > self.nums.len() * 4 {
            let mut new_datas = vec![vec![]; self.nums.len() * 2];
            let mut new_len = 0;
            for v in &self.nums {
                for &(k, v) in v {
                    if v != -1 {
                        new_datas[k as usize % (self.nums.len() * 2)].push((k, v));
                        new_len += 1;
                    }
                }
            }
            self.nums = new_datas;
            self.len = new_len;
        }
    }

    fn get_hash(&self, key: i32) -> usize {
        key as usize % self.nums.len()
    }

    fn put(&mut self, key: i32, value: i32) {
        self.resize();
        //先看有没有,如果有，则直接更新
        let index = self.get_hash(key);
        for e in &mut self.nums[index] {
            if (*e).0 == key {
                (*e).1 = value;
                return;
            }
        }
        //没有则创建
        self.nums[self.get_hash(key)].push((key, value));
        self.len += 1;
    }

    fn get(&self, key: i32) -> i32 {
        for &e in &self.nums[self.get_hash(key)] {
            if e.0 == key {
                return e.1;
            }
        }
        -1
    }

    fn remove(&mut self, key: i32) {
        self.put(key, -1)
    }
}

struct Solution;

impl Solution {
    //1466. 重新规划路线
    pub fn min_reorder(n: i32, connections: Vec<Vec<i32>>) -> i32 {
        //广度优先遍历或深度优先遍历都可。从0开始
        //先将数组变成可到达的地方。先不管方向
        let mut map = vec![vec![]; n as usize];
        let mut already = vec![false; n as usize];
        for i in 0..connections.len() {
            let src = connections[i][0];
            let dst = connections[i][1];
            map[src as usize].push((dst, 1));
            map[dst as usize].push((src, 0));
        }
        //从0开始深度优先遍历。怎么确定路线是否需要更改呢？在map中所存路线中用一个bool来记录,为true的是需要更改的。
        let mut stack = vec![0];
        already[0] = true;
        let mut res = 0;
        while let Some(n) = stack.pop() {
            //得到它能访问的城市
            for (next, is_change) in &map[n] {
                if !already[*next as usize] {
                    already[*next as usize] = true;
                    stack.push(*next as usize);
                    res += is_change;
                }
            }
        }
        res
    }
    //688. 骑士在棋盘上的概率
    pub fn knight_probability(n: i32, k: i32, row: i32, column: i32) -> f64 {
        let mut dp = vec![vec![vec![0.0; n as usize]; n as usize]; k as usize + 1];
        for i in 0..n as usize {
            for j in 0..n as usize {
                dp[0][i][j] = 1.0;
            }
        }
        let direct = vec![-2i32, -1, 1, 2];
        for i in 1..=k as usize {
            for x in 0..n as usize {
                for y in 0..n as usize {
                    for &u in &direct {
                        for &v in &direct {
                            if u.abs() + v.abs() == 3
                                && x as i32 + u >= 0 && x as i32 + u < n && y as i32 + v >= 0 && y as i32 + v < n {
                                dp[i][x][y] += dp[i - 1][(x as i32 + u) as usize][(y as i32 + v) as usize];
                            }
                        }
                    }
                }
            }
        }
        let mut gross = 1.0;
        for _ in 0..k {
            gross *= 8.0;
        }
        dp[k as usize][row as usize][column as usize] / gross
    }

    pub fn knight_probability2(n: i32, k: i32, row: i32, column: i32) -> f64 {
        //递归算法超时：用动态规划算法呢？用java来实现
        //dp[k][x][y],k代表步数。dp[0].. = 1;
        // dp[1][x][y] = 从x,y 能到达哪
        //总个数.最大可能达到8^100
        let mut gross = 1.0;
        for _ in 0..k {
            gross *= 8.0;
        }
        Self::knight_gross(n, k, row, column) / gross
    }

    fn knight_gross(n: i32, k: i32, row: i32, column: i32) -> f64 {
        let mut gross = 0.0;
        if k == 0 {
            return 1.0;
        }
        for i in vec![-1i32, -2, 1, 2] {
            for j in vec![-1i32, -2, 1, 2] {
                if i.abs() + j.abs() == 3 {
                    if row + i >= 0 && row + i < n && column + j >= 0 && column + j < n {
                        gross += Self::knight_gross(n, k - 1, row + i, column + j);
                    }
                }
            }
        }
        gross
    }
    //40. 组合总和 II
    pub fn combination_sum2(mut candidates: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        //采用递归来做
        //怎么表示数字已经被使用了呢？用一个数组来记录
        candidates.sort_unstable();
        Self::combination_sum(&candidates, target, &mut vec![false; candidates.len()], 0)
    }

    fn combination_sum(candidates: &[i32], target: i32, already: &mut Vec<bool>, pre: i32) -> Vec<Vec<i32>> {
        let mut res = vec![];
        if target == 0 {
            let mut item = vec![];
            for i in 0..already.len() {
                if already[i] {
                    item.push(candidates[i]);
                }
            }
            res.push(item);
        } else {
            let mut pre_num = 0;
            //从剩下的元素中选择一个元素
            for i in 0..candidates.len() {
                //过滤掉已经选择过的元素
                if !already[i] {
                    if candidates[i] != pre_num && candidates[i] <= target && candidates[i] >= pre {
                        already[i] = true;
                        res.append(&mut Self::combination_sum(candidates, target - candidates[i], already, candidates[i]));
                        already[i] = false;
                    }
                    pre_num = candidates[i];
                }
            }
        }
        res
    }

    //1376 通知所有员工所需的时间
    pub fn num_of_minutes(n: i32, head_id: i32, manager: Vec<i32>, inform_time: Vec<i32>) -> i32 {
        //广度优先遍历，而应该用深度优先算法，计算最长路径
        // 先把manager数组变成一个vec[vec] v为该员工的下属
        let mut subs = vec![vec![]; n as usize];
        for i in 0..n as usize {
            //找到它的上级
            if manager[i] != -1 {
                subs[manager[i] as usize].push(i);
            }
        }
        //需要记录通知到某个员工的时间
        let mut reach_time = vec![0; n as usize];
        reach_time[head_id as usize] = inform_time[head_id as usize];
        let mut stack = vec![];
        stack.push(head_id as usize);
        //需要多少时间来通知它的下属
        while let Some(n) = stack.pop() {
            //如果没有下属了，代表这条线路走到了终点。
            //它的下属们进栈
            for sub in &subs[n] {
                stack.push(*sub);
                //上一级的时间 + 本机的通知时间
                reach_time[*sub] = reach_time[n] + inform_time[*sub];
            }
        }
        reach_time
            .into_iter()
            .max()
            .unwrap()
    }

    //47. 全排列 II
    pub fn permute_unique(mut nums: Vec<i32>) -> Vec<Vec<i32>> {
        nums.sort_unstable();
        //选择一个数，然后选择第二个数，如果接下来的数和自己相同，跳过该数。
        Self::select_next_num(nums, vec![])
    }

    fn select_next_num(nums: Vec<i32>, pre: Vec<i32>) -> Vec<Vec<i32>> {
        let mut res = vec![];
        if nums.len() == 0 {
            res.push(pre);
            return res;
        }
        let mut pre_num = 11;
        for i in 0..nums.len() {
            if nums[i] != pre_num {
                //选择当前数字
                let mut new_pre = pre.clone();
                new_pre.push(nums[i]);
                let mut new_nums = vec![0; nums.len() - 1];
                for j in 0..i {
                    new_nums[j] = nums[j];
                }
                for j in i + 1..nums.len() {
                    new_nums[j - 1] = nums[j];
                }
                res.append(&mut Self::select_next_num(new_nums, new_pre));
                pre_num = nums[i];
            }
        }
        res
    }
}