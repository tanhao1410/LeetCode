use std::i32::MAX;

fn main() {
    println!("Hello, world!");
    Solution::k_closest(vec![vec![1, 3], vec![-2, 2]], 1);
}

struct Solution {}

impl Solution {

    //139. 单词拆分
    pub fn word_break(s: String, word_dict: Vec<String>) -> bool {
        //思路：递归法，先 截取一个可以截的，剩下的继续调用。
        //可能会存在多个可以截的，循环即可。
        fn word_break2(s: String, word_set: &std::collections::HashSet<String>) -> bool {
            if s.len() == 0 || word_set.contains(&s){
                return true;
            }
            for i in 0..s.len() {
                if word_set.contains(&s[..i].to_string()) && word_break2(s[i..].to_string(), word_set) {
                    return true;
                }
            }
            false
        }
        let mut word_set = std::collections::HashSet::new();
        for i in word_dict.clone() {
            word_set.insert(i);
        }
        //对于可以由其他单词组成的单词，不应该再放进去
        for i in word_dict{
            word_set.remove(&i);
            if !word_break2(i.clone(),&word_set){
                word_set.insert(i);
            }
        }

        word_break2(s,&word_set)
    }

    //130. 被围绕的区域
    pub fn solve(board: &mut Vec<Vec<char>>) {
        //思路：从最外圈的O开始走，将该字母换成1，走完后，再遍历矩阵，将O换成X，1换成o即可。
        //先找最外层的o
        if board.is_empty() || board[0].is_empty() {
            return;
        }
        let mut queue = vec![];
        //上下左右四块
        for i in 0..board[0].len() {
            if board[0][i] == 'O' {
                queue.push((0, i))
            }
            if board[board.len() - 1][i] == 'O' {
                queue.push((board.len() - 1, i))
            }
        }
        for i in 0..board.len() {
            if board[i][0] == 'O' {
                queue.push((i, 0))
            }
            if board[i][board[0].len() - 1] == 'O' {
                queue.push((i, board[0].len() - 1))
            }
        }

        while queue.len() > 0 {
            let x = queue.pop().unwrap();
            board[x.0][x.1] = 'o';
            //看它的上下左右是否为O，为的话，加入队列
            //上
            if x.0 > 0 && board[x.0 - 1][x.1] == 'O' {
                queue.push((x.0 - 1, x.1));
            }
            if x.0 < board.len() - 1 && board[x.0 + 1][x.1] == 'O' {
                queue.push((x.0 + 1, x.1));
            }
            if x.1 > 0 && board[x.0][x.1 - 1] == 'O' {
                queue.push((x.0, x.1 - 1));
            }
            if x.1 < board[0].len() - 1 && board[x.0][x.1 + 1] == 'O' {
                queue.push((x.0, x.1 + 1));
            }
        }

        for i in board {
            for j in i {
                if *j == 'O' {
                    *j = 'X';
                } else if *j == 'o' {
                    *j = 'O';
                }
            }
        }
    }

    //973. 最接近原点的 K 个点
    pub fn k_closest(points: Vec<Vec<i32>>, k: i32) -> Vec<Vec<i32>> {
        //思路：用一个vec记录最接近的点，发现更小的后，清空原来的，加入新的
        //找出k个，而不仅仅是最小的几个。插入排序法
        let (mut res, mut cur_dis) = (Vec::new(), std::i32::MAX);

        fn get_dis(point: &Vec<i32>) -> i32 {
            point[0] * point[0] + point[1] * point[1]
        }
        for i in 0..k {
            let mut index = 0;
            while index < res.len() && get_dis(&res[index]) < get_dis(&points[i as usize]) {
                index += 1;
            }
            res.insert(index, points[i as usize].clone());
        }
        for i in k..points.len() as i32 {
            let mut index = k - 1;
            while index >= 0 && get_dis(&res[index as usize]) > get_dis(&points[i as usize]) {
                index -= 1;
            }
            if index < 0 {
                res.insert(0  ,points[i as usize].clone());
                res.pop();
            } else if index < k - 1 {
                res.insert(index as usize + 1, points[i as usize].clone());
                res.pop();
            }
        }
        res
    }
}
