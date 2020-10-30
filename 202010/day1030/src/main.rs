fn main() {
    println!("Hello, world!");
    println!("{}",Solution::cutting_rope2(1000));
}

struct Solution{

}

impl Solution {

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
                can.push((m , n+1));
            }
            if n as i32 - 1 >= 0 && board[m][n - 1] == word.chars().nth(0).unwrap() {
                can.push((m , n-1));
            }
            if can.len()==0{
                return false;
            }

            for i in can.into_iter(){
                board[i.0][i.1] = ' ';
                if next(board,i.0,i.1,&word[1..]){
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
        if grid.len() == 0 || grid[0].len() == 0{
            return 0
        }
        //思路：从四个方向分别求出对应的边长，然后相加
        let  (mut left_num,mut right_num,mut up_num,mut down_num) = (0,0,0,0);
        for i in 0..grid.len(){
            let (mut left_flag,mut right_flag) = (true,true);
            for j in 0..grid[0].len(){
                if grid[i][j] == 1 && left_flag{
                    left_num +=1;
                    left_flag = false;
                }
                if grid[i][j] == 0{
                    left_flag = true;
                }
            }
            for j in (0..grid[0].len()).rev(){
                if grid[i][j] == 1 && right_flag{
                    right_num +=1;
                    right_flag = false;
                }
                if grid[i][j] == 0{
                    right_flag = true;
                }
            }
        }
        for i in 0..grid[0].len(){
            let (mut up_flag,mut down_flag) = (true,true);
            for j in 0..grid.len(){
                if grid[j][i] == 1 && up_flag{
                    up_num +=1;
                    up_flag = false;
                }
                if grid[j][i] == 0{
                    up_flag = true;
                }
            }
            for j in 0..grid.len(){
                if grid[j][i] == 1 && down_flag{
                    down_num +=1;
                    down_flag = false;
                }
                if grid[j][i] == 0{
                    down_flag = true;
                }
            }
        }
        left_num+right_num+up_num+down_num
    }
}