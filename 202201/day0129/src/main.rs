use std::collections::HashSet;

fn main() {
    println!("Hello, world!");
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