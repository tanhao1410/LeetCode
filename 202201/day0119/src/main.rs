fn main() {
    //重量是c,价值是w
    let w = vec![1, 2, 1, 3, 4, 1];
    let c = vec![4, 5, 3, 1, 3, 1];
    let v = 15;
    println!("{}", one_zero_problem(c.clone(), w.clone(), v));
    println!("{}", one_zero_problem_space_on(c.clone(), w.clone(), v));
    println!("{}", one_zero_problem_recursion(c, w, v));

    // println!("{}", complete_bagger_problem(c.clone(), w.clone(), v));
    // println!("{}", complete_bagger_recursion(c.clone(), w.clone(), v));
}

//416. 分割等和子集
pub fn can_partition(nums: Vec<i32>) -> bool {
    let count = nums.iter().sum::<i32>();
    if count % 2 == 1 {
        return false;
    }
    let target = count / 2;

    //求组合总数
    // 背包的大小是target 代表容量,商品是nums,为重量, 0,1 背包问题，唯一不同的在于，价值？
    //用一维数组来做
    let mut dp = vec![false; target as usize + 1];
    for i in 0..nums.len() {
        for v in (0..=target).rev() {
            if i == 0 {
                if v == nums[0] {
                    dp[v as usize] = true;
                }
            }
            let v = v as usize;
            if v >= nums[i] as usize {
                dp[v] = dp[v] || dp[v - nums[i] as usize]
            }
        }
    }
    dp[target as usize]
}

//完全背包问题
pub fn complete_bagger_problem(c: Vec<i32>, w: Vec<i32>, v: i32) -> i32 {
    let mut dp = vec![0; v as usize + 1];
    for i in 0..c.len() {
        for v in 0..=v {
            let v = v as usize;
            if v >= c[i] as usize {
                dp[v] = dp[v].max(dp[v - c[i] as usize] + w[i]);
            }
        }
    }
    dp[v as usize]
}

pub fn complete_bagger_recursion(c: Vec<i32>, w: Vec<i32>, v: i32) -> i32 {
    //每一个物品都可以选择放或不放
    complete_bagger_recursion_inner(&c, &w, v)
}

pub fn complete_bagger_recursion_inner(c: &[i32], w: &[i32], v: i32) -> i32 {
    //每一个物品都可以选择放或不放
    if c.len() < 2 {
        //就剩一个物品了
        if c[0] > v {
            return 0;
        }
        return w[0];
    } else {
        //放或不放
        let mut res = 0;
        if c[0] <= v {
            //放
            res = complete_bagger_recursion_inner(&c[..], &w[..], v - c[0]) + w[0];
        }
        //不放，取大的
        res.max(complete_bagger_recursion_inner(&c[1..], &w[1..], v))
    }
}


pub fn one_zero_problem_recursion(c: Vec<i32>, w: Vec<i32>, v: i32) -> i32 {
    //每一个物品都可以选择放或不放
    one_zero_problem_recursion_inner(&c, &w, v)
}

pub fn one_zero_problem_recursion_inner(c: &[i32], w: &[i32], v: i32) -> i32 {
    //每一个物品都可以选择放或不放
    if c.len() < 2 {
        //就剩一个物品了
        if c[0] > v {
            return 0;
        }
        return w[0];
    } else {
        //放或不放
        let mut res = 0;
        if c[0] <= v {
            //放
            res = one_zero_problem_recursion_inner(&c[1..], &w[1..], v - c[0]) + w[0];
        }
        //不放，取大的
        res.max(one_zero_problem_recursion_inner(&c[1..], &w[1..], v))
    }
}

pub fn one_zero_problem(c: Vec<i32>, w: Vec<i32>, v: i32) -> i32 {
    let mut dp = vec![vec![0; v as usize + 1]; c.len()];
    //二维方式解决  0 1背包问题
    for i in 0..c.len() {
        for v in 0..=v {
            let v = v as usize;
            if i == 0 {
                //只有第一个能选择
                if v >= c[0] as usize {
                    dp[0][v] = w[0];
                }
            } else {
                if c[i] <= v as i32 {
                    //能放下的情况
                    dp[i][v] = dp[i - 1][v].max(dp[i - 1][v - c[i] as usize] + w[i]);
                } else {
                    dp[i][v] = dp[i - 1][v];
                }
            }
        }
    }

    dp[c.len() - 1][v as usize]
}

// 0 1 背包问题 优化空间
pub fn one_zero_problem_space_on(c: Vec<i32>, w: Vec<i32>, v: i32) -> i32 {
    let mut dp = vec![0; v as usize + 1];
    //二维方式解决  0 1背包问题
    for i in 0..c.len() {
        for v in (0..=v).rev() {
            let v = v as usize;
            if v >= c[i] as usize {
                dp[v] = dp[v].max(dp[v - c[i] as usize] + w[i]);
            }
        }
    }

    dp[v as usize]
}




//766. 托普利茨矩阵
pub fn is_toeplitz_matrix(matrix: Vec<Vec<i32>>) -> bool {
    //每一行矩阵，去掉最后一个元素，应该与 下一行矩阵 去掉第一个元素相等
    let row = matrix.len();
    let col = matrix[0].len();
    for i in 0..row - 1 {
        let cur_row = &matrix[i];
        let next_row = &matrix[i + 1];
        //判断两者是否相等
        for j in 0..col - 1 {
            if cur_row[j] != next_row[j + 1] {
                return false;
            }
        }
    }
    true
}

//598. 范围求和 II
pub fn max_count(m: i32, n: i32, ops: Vec<Vec<i32>>) -> i32 {
    //matrix[0][0]肯定是最大的整数，其它的只要根据ops来不断的缩小范围即可。
    ops
        .iter()
        .map(|v| (v[0], v[1]))
        .fold((m, n, m * n), |(m, n, mul), (i, j)| {
            let (m, n) = (i.min(m), j.min(n));
            (m, n, m * n)
        })
        .2
}

pub fn contains_nearby_duplicate2(nums: Vec<i32>, k: i32) -> bool {
    use std::collections::HashSet;
    //滑动窗口的大小为k,每一次都往里添加元素，如果添加元素后，窗口小于k，说明又重复的了。
    let mut set = HashSet::new();
    for i in 0..nums.len() {
        if set.len() > k as usize {
            set.remove(&nums[i - 1 - k as usize]);
        }
        if set.contains(&nums[i]) {
            return true;
        }
        set.insert(nums[i]);
    }
    false
}

//219. 存在重复元素 II
pub fn contains_nearby_duplicate(nums: Vec<i32>, k: i32) -> bool {
    use std::collections::HashMap;
    let mut map = HashMap::new();
    nums
        .iter()
        .enumerate()
        .for_each(|(i, n)| {
            let vec = map.entry(n).or_insert(Vec::new());
            vec.push(i)
        });
    map
        .iter()
        .any(|(_, v)| {
            v.len() > 1
                && v.iter()
                .fold((-k - 1, false), |(pre, res), &n| (n as i32, res || n as i32 - pre <= k))
                .1
        })
}