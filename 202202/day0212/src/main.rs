fn main() {
    println!("Hello, world!");
}

//1020. 飞地的数量
pub fn num_enclaves(mut grid: Vec<Vec<i32>>) -> i32 {
    ///采用队列，将边缘所有的1加入进来，然后移动，不停的加入，退出
    let mut queue = vec![];
    let row_num = grid.len();
    let col_num = grid[0].len();
    for i in 0..row_num{
        if grid[i][0] == 1{
            queue.push((i,0));
            grid[i][0] = 0;
        }
        if grid[i][col_num - 1] == 1{
            queue.push((i,col_num - 1));
            grid[i][col_num - 1] = 0;
        }
    }
    for j in 1..col_num - 1{
        if grid[0][j] == 1{
            queue.push((0,j));
            grid[0][j] = 0;
        }
        if grid[row_num - 1][j] == 1{
            queue.push((row_num - 1,j));
            grid[row_num - 1][j] = 0;
        }
    }
    while let Some((x,y)) = queue.pop(){
        //上下左右
        if x > 0 && grid[x - 1][y] == 1{
            queue.push((x - 1,y));
            grid[x-1][y] = 0;
        }
        if x < grid.len() - 1 && grid[x + 1][y] == 1{
            queue.push((x + 1,y));
            grid[x+1][y] = 0;
        }
        if y > 0 && grid[x][y - 1] == 1{
            queue.push((x,y - 1));
            grid[x][y -1] = 0;
        }
        if y < grid[0].len() - 1 && grid[x][y + 1] == 1{
            queue.push((x,y + 1));
            grid[x][y + 1] = 0;
        }
    }
    //遍历矩阵
    let mut res = 0;
    for i in 1..grid.len() - 1{
        for j in 1..grid[0].len() - 1{
            if grid[i][j] == 1{
                res += 1;
            }
        }
    }
    res
}