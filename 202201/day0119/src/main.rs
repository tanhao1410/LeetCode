fn main() {
    println!("Hello, world!");
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