use std::collections::HashSet;

fn main() {
    println!("Hello, world!");
}

//1219. 黄金矿工
pub fn get_maximum_gold(grid: Vec<Vec<i32>>) -> i32 {
    // 思路： 递归法可以解决。深度优先遍历。
    // 开始时 最多有25个选择
    // 从每一个选择处出发，怎么确定最大收益？
    let mut starts = vec![];
    for i in 0..grid.len() {
        for j in 0..grid[0].len() {
            if grid[i][j] != 0 {
                starts.push((i, j));
            }
        }
    }
    let mut res = 0;
    let mut set = HashSet::new();
    for (x, y) in starts {
        res = res.max(get_maximum_gold2(&grid, x, y, &mut set));
    }
    res
}

pub fn get_maximum_gold2(grid: &Vec<Vec<i32>>, x: usize, y: usize, already: &mut HashSet<(usize, usize)>) -> i32 {
    //递归的形式去找最大收益，
    if grid[x][y] == 0 || already.contains(&(x, y)) {
        return 0;
    }
    let mut res = grid[x][y];
    already.insert((x, y));
    //四个方向，哪个最大去哪个
    let mut other = 0;
    if x > 0 {
        other = other.max(get_maximum_gold2(grid, x - 1, y, already));
    }
    if y > 0 {
        other = other.max(get_maximum_gold2(grid, x, y - 1, already));
    }
    if x < grid.len() - 1 {
        other = other.max(get_maximum_gold2(grid, x + 1, y, already));
    }
    if y < grid[0].len() - 1 {
        other = other.max(get_maximum_gold2(grid, x, y + 1, already));
    }
    already.remove(&(x, y));
    res + other
}

//1143. 最长公共子序列
pub fn longest_common_subsequence(text1: String, text2: String) -> i32 {
    //思路：dp[i][j] , text1[..i]与text2[..j]的最长公共子序列
    // dp[0][j] = 0 或1 ， if dp[0][j] == 1 ,则 dp[0][j++] = 1
    // if text1[i] == text2[j] ,则 的dp[i][j] = dp[i - 1][j-1] + 1
    let bytes1 = text1.as_bytes();
    let bytes2 = text2.as_bytes();
    let mut dp = vec![vec![0; text2.len()]; text1.len()];
    //求两者第一个字母相遇的位置
    let mut l1 = 0;
    while l1 < text2.len() && bytes2[l1] != bytes1[0] {
        l1 += 1;
    }
    for i in l1..text2.len() {
        dp[0][i] = 1;
    }
    let mut l2 = 0;
    while l2 < text1.len() && bytes1[l2] != bytes2[0] {
        l2 += 1;
    }
    for i in l2..text1.len() {
        dp[i][0] = 1;
    }

    for i in 1..bytes1.len() {
        for j in 1..bytes2.len() {
            if bytes1[i] == bytes2[j] {
                dp[i][j] = dp[i - 1][j - 1] + 1;
            } else {
                dp[i][j] = dp[i - 1][j].max(dp[i][j - 1]);
            }
        }
    }

    dp[text1.len() - 1][text2.len() - 1]
}

//392. 判断子序列
pub fn is_subsequence(s: String, t: String) -> bool {
    //思路：双指针法，对于任意一个s中的字母，要在t中找到，找到后，s指针往前走
    let s_bytes = s.as_bytes();
    let t_byts = t.as_bytes();
    let mut s = 0;
    let mut t = 0;
    while s < s_bytes.len() && t < t_byts.len() {
        while t < t_byts.len() && s_bytes[s] != t_byts[t] {
            t += 1;
        }
        if t < t_byts.len() {
            s += 1;
            t += 1;
        }
    }
    s == s_bytes.len()
}