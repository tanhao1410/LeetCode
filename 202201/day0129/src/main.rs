use std::collections::{HashSet, VecDeque};

fn main() {
    println!("Hello, world!");
}

//994. 腐烂的橘子
pub fn oranges_rotting(mut grid: Vec<Vec<i32>>) -> i32 {
    //广度优先
    let mut queue = VecDeque::new();
    let mut visit = vec![vec![true; grid[0].len()]; grid.len()];
    //从等于2处开始腐烂
    for x in 0..grid.len() {
        for y in 0..grid[0].len() {
            if grid[x][y] == 2 {
                queue.push_back((x, y));
            }
        }
    }
    let mut res = 0;
    //已经遍历过的需要不再遍历
    let mut size = queue.len();
    while size > 0 {
        for _ in 0..size {
            let (x, y) = queue.pop_front().unwrap();
            grid[x][y] = 2;
            //将它周围的好橘子腐烂
            if x > 0 && grid[x - 1][y] == 1 && visit[x - 1][y] {
                visit[x - 1][y] = false;
                queue.push_back((x - 1, y));
            }
            if x + 1 < grid.len() && grid[x + 1][y] == 1 && visit[x + 1][y] {
                visit[x + 1][y] = false;
                queue.push_back((x + 1, y));
            }
            if y > 0 && grid[x][y - 1] == 1 && visit[x][y - 1] {
                visit[x][y - 1] = false;
                queue.push_back((x, y - 1));
            }
            if y + 1 < grid[0].len() && grid[x][y + 1] == 1 && visit[x][y + 1] {
                visit[x][y + 1] = false;
                queue.push_back((x, y + 1));
            }
        }
        size = queue.len();
        res += 1;
    }
    let mut flag = true;
    for row in &grid {
        for &v in row {
            if v == 1 {
                return -1;
            }
            if v == 2 {
                flag = false;
            }
        }
    }
    if flag { 0 } else { res - 1 }
}

//542. 01 矩阵
pub fn update_matrix(mat: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
    let mut res = vec![vec![-1; mat[0].len()]; mat.len()];
    //广度优先遍历
    let mut set = HashSet::new();
    // 从所有为0的位置开始遍历，遍历第一次的时候，接着的都是1
    for row in 0..mat.len() {
        for col in 0..mat[0].len() {
            if mat[row][col] == 0 {
                set.insert((row, col));
            }
        }
    }

    let mut distance = 0;
    let mut size = set.len();
    while size > 0 {
        let mut new_set = HashSet::new();
        for (x, y) in set.iter() {
            let x = *x;
            let y = *y;
            res[x][y] = distance;
        }
        for (x, y) in set.iter() {
            let x = *x;
            let y = *y;
            //确定(x,y) 四周是否可以加入进来
            if x >= 1 && mat[x - 1][y] == 1 && res[x - 1][y] == -1 {
                new_set.insert((x - 1, y));
            }
            if x + 1 < mat.len() && mat[x + 1][y] == 1 && res[x + 1][y] == -1 {
                new_set.insert((x + 1, y));
            }
            if y >= 1 && mat[x][y - 1] == 1 && res[x][y - 1] == -1 {
                new_set.insert((x, y - 1));
            }
            if y + 1 < mat[0].len() && mat[x][y + 1] == 1 && res[x][y + 1] == -1 {
                new_set.insert((x, y + 1));
            }
        }
        set.clear();
        set.extend(new_set.iter());
        size = set.len();
        distance += 1;
    }
    res
}


//118. 杨辉三角
pub fn generate(num_rows: i32) -> Vec<Vec<i32>> {
    let mut res = vec![vec![1]];
    for _ in 1..num_rows as usize {
        let cur_line = &res[res.len() - 1];
        let mut next_line = vec![1; cur_line.len() + 1];
        for i in 1..next_line.len() - 1 {
            next_line[i] = cur_line[i - 1] + cur_line[i];
        }
        res.push(next_line);
    }
    res
}