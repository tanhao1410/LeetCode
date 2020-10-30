fn main() {
    println!("Hello, world!");
    //println!("{}", Solution::translate_num(12323));
}

struct Solution {}


impl Solution {
    //剑指 Offer 66. 构建乘积数组
    pub fn construct_arr(a: Vec<i32>) -> Vec<i32> {
        let mut res = a.clone();
        //不能采用除法，必然会存在许多相同的数字，特别是1，不然肯定会突破32位最大限制
        //那就用一个map来做记录
        let mut m = std::collections::HashMap::new();

        fn get_construct(a: &Vec<i32>, index: usize) -> i32 {
            let mut res = 1;
            for i in 0..a.len() {
                if i != index {
                    res *= a[i]
                }
            }
            res
        }

        for i in 0..res.len() {
            if let Some(value) = m.get(&a[i]) {
                res[i] = *value
            } else {
                let value = get_construct(&a, i);
                res[i] = value;
                m.insert(a[i], value);
            }
        }
        res
    }

    //剑指 Offer 46. 把数字翻译成字符串
    pub fn translate_num(num: i32) -> i32 {
        //直接递归就行了，把数字转换成字符串进行判别即可，或者不进行转换，直接递归也行
        let s = num.to_string();
        fn translate(s: String) -> i32 {
            if s.len() < 2 {
                return 1;
            }
            if s.chars().nth(0).unwrap() == '1' || (s.chars().nth(0).unwrap() == '2' && s.chars().nth(1).unwrap() < '6') {
                return translate(s[1..].to_string()) + translate(s[2..].to_string());
            } else {
                return translate(s[1..].to_string());
            }
        }
        translate(s)
    }

    //剑指 Offer 47. 礼物的最大价值
    pub fn max_value(grid: Vec<Vec<i32>>) -> i32 {
        //思路：普通动态规划题目
        let mut dp = grid.clone();
        let (m, n) = (grid.len() - 1, grid[0].len() - 1);
        for i in (0..m + 1).rev() {
            for j in (0..n + 1).rev() {
                if j == n && i == m {} else if i == m && j < n {
                    //最后一行
                    dp[m][j] = dp[m][j + 1] + grid[i][j];
                } else if i < m && j == n {
                    //最后一列
                    dp[i][j] = dp[i + 1][j] + grid[i][j];
                } else {
                    if dp[i + 1][j] > dp[i][j + 1] {
                        dp[i][j] = dp[i + 1][j] + grid[i][j];
                    } else {
                        dp[i][j] = dp[i][j + 1] + grid[i][j];
                    }
                }
            }
        }
        dp[0][0]
    }

    //剑指 Offer 14- I. 剪绳子
    pub fn cutting_rope(n: i32) -> i32 {
        let mut dp = vec![0, 0, 1, 2, 4];
        for i in 5..n + 1 {
            let mut max = 0;
            for j in 1..i / 2 + 1 {
                if j * dp[(i - j) as usize] > max {
                    max = j * dp[(i - j) as usize];
                }
                if j * (i - j) > max {
                    max = j * (i - j);
                }
            }
            dp.push(max);
        }
        dp[n as usize]
    }

    //剑指 Offer 12. 矩阵中的路径
    pub fn exist(board: Vec<Vec<char>>, word: String) -> bool {
        let mut board = board;
        //回溯法，先找到第一个可以走的，再走，遇到不满足的就回退
        fn next(board: &mut Vec<Vec<char>>, m: usize, n: usize, word: &str) -> bool {
            if word.len() == 0 {
                return true;
            }
            //找下一个可以走的地方符合的
            let mut can = vec![];
            //往下
            if m + 1 < board.len() && board[m + 1][n] == word.chars().nth(0).unwrap() {
                can.push((m + 1, n));
            }
            if m as i32 - 1 >= 0 && board[m - 1][n] == word.chars().nth(0).unwrap() {
                can.push((m - 1, n));
            }
            if n + 1 < board[0].len() && board[m][n + 1] == word.chars().nth(0).unwrap() {
                can.push((m, n + 1));
            }
            if n as i32 - 1 >= 0 && board[m][n - 1] == word.chars().nth(0).unwrap() {
                can.push((m, n - 1));
            }
            if can.len() == 0 {
                return false;
            }

            for i in can.into_iter() {
                board[i.0][i.1] = ' ';
                if next(board, i.0, i.1, &word[1..]) {
                    return true;
                }
                board[i.0][i.1] = word.chars().nth(0).unwrap();
            }
            false
        }

        let (m, n) = (board.len(), board[0].len());
        for i in 0..m {
            for j in 0..n {
                if board[i][j] == word.chars().nth(0).unwrap() {
                    board[i][j] = ' ';
                    if next(&mut board, i, j, &word[1..]) {
                        return true;
                    }
                    board[i][j] = word.chars().nth(0).unwrap();
                }
            }
        }
        false
    }

    //每日一题；463 岛屿的周长:新思路：左右和上下其实是相等的，减少一遍求。2.也可以是采用总的岛屿数*4 - 重合的边*2
    pub fn island_perimeter(grid: Vec<Vec<i32>>) -> i32 {
        if grid.len() == 0 || grid[0].len() == 0 {
            return 0;
        }
        //思路：从四个方向分别求出对应的边长，然后相加
        let (mut left_num, mut right_num, mut up_num, mut down_num) = (0, 0, 0, 0);
        for i in 0..grid.len() {
            let (mut left_flag, mut right_flag) = (true, true);
            for j in 0..grid[0].len() {
                if grid[i][j] == 1 && left_flag {
                    left_num += 1;
                    left_flag = false;
                }
                if grid[i][j] == 0 {
                    left_flag = true;
                }
            }
            for j in (0..grid[0].len()).rev() {
                if grid[i][j] == 1 && right_flag {
                    right_num += 1;
                    right_flag = false;
                }
                if grid[i][j] == 0 {
                    right_flag = true;
                }
            }
        }
        for i in 0..grid[0].len() {
            let (mut up_flag, mut down_flag) = (true, true);
            for j in 0..grid.len() {
                if grid[j][i] == 1 && up_flag {
                    up_num += 1;
                    up_flag = false;
                }
                if grid[j][i] == 0 {
                    up_flag = true;
                }
            }
            for j in 0..grid.len() {
                if grid[j][i] == 1 && down_flag {
                    down_num += 1;
                    down_flag = false;
                }
                if grid[j][i] == 0 {
                    down_flag = true;
                }
            }
        }
        left_num + right_num + up_num + down_num
    }
}