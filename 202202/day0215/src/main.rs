fn main() {
    println!("Hello, world!");
}

//1091. 二进制矩阵中的最短路径
pub fn shortest_path_binary_matrix(mut grid: Vec<Vec<i32>>) -> i32 {
    use std::collections::VecDeque;
    //广度优先遍历
    if grid[0][0] == 1 || grid[grid.len() - 1][grid.len() - 1] == 1 {
        return -1;
    }
    //从grid[0][0] 处开始广度优先遍历，
    let mut queue = VecDeque::new();
    queue.push_back((0, 0));
    grid[0][0] = 1;//访问过的不再访问
    let n = grid.len();
    let mut distance = 0;
    let mut len = queue.len();
    while len > 0 {
        distance += 1;
        for _ in 0..len {
            let (x, y) = queue.pop_front().unwrap();
            if x == grid.len() - 1 && x == y {
                return distance;
            }
            if x > 0 && grid[x - 1][y] == 0 {
                queue.push_back((x - 1, y));
                grid[x - 1][y] = 1;
            }
            if x < n - 1 && grid[x + 1][y] == 0 {
                queue.push_back((x + 1, y));
                grid[x + 1][y] = 1;
            }
            if y > 0 && grid[x][y - 1] == 0 {
                queue.push_back((x, y - 1));
                grid[x][y - 1] = 1;
            }
            if y < n - 1 && grid[x][y + 1] == 0 {
                queue.push_back((x, y + 1));
                grid[x][y + 1] = 1;
            }
            if x > 0 && y > 0 && grid[x - 1][y - 1] == 0 {
                queue.push_back((x - 1, y - 1));
                grid[x - 1][y - 1] = 1;
            }
            if x > 0 && y < n - 1 && grid[x - 1][y + 1] == 0 {
                queue.push_back((x - 1, y + 1));
                grid[x - 1][y + 1] = 1;
            }
            if x < n - 1 && y > 0 && grid[x + 1][y - 1] == 0 {
                queue.push_back((x + 1, y - 1));
                grid[x + 1][y - 1] = 1;
            }
            if x < n - 1 && y < n - 1 && grid[x + 1][y + 1] == 0 {
                queue.push_back((x + 1, y + 1));
                grid[x + 1][y + 1] = 1;
            }
        }
        len = queue.len();
    }
    -1
}

//1380. 矩阵中的幸运数
pub fn lucky_numbers(matrix: Vec<Vec<i32>>) -> Vec<i32> {
    //同一行中最小，同一列中最大
    let mut res = vec![];
    for i in 0..matrix.len() {
        //找到一行中最小的数
        let mut min = 0;
        for j in 1..matrix[0].len() {
            if matrix[i][j] < matrix[i][min] {
                min = j;
            }
        }
        //判断它是否是一列中最大的数
        let mut max = i;
        for j in 0..matrix.len() {
            if matrix[j][min] > matrix[i][min] {
                max = j;
            }
        }

        if max == i {
            res.push(matrix[i][min]);
        }
    }

    res
}
