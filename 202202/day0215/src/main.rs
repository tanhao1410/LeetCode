fn main() {
    println!("Hello, world!");
}

//15. 三数之和
pub fn three_sum(mut nums: Vec<i32>) -> Vec<Vec<i32>> {
    use std::collections::HashMap;
    let mut map = HashMap::new();
    for &num in &nums {
        let mut entry = map.entry(num).or_insert(0);
        *entry += 1;
    }

    // 选择三个数，只有（0,0,0), 选择有两个相同的数情况。其余的只能选择三个不同的数了。
    let mut res = vec![];
    //三个相同的数情况。
    if let Some(count) = map.get_mut(&0) {
        if *count > 2 {
            res.push(vec![0; 3]);
        }
    }

    //两个数相等的情况。 =
    res.append(&mut map
        .iter()
        .filter(|(k, v)| **v > 1 && **k != 0)
        .filter_map(|(k, _)| {
            match map.contains_key(&(-2 * *k)) {
                true => Some(vec![*k, *k, -2 * *k]),
                false => None
            }
        })
        .collect::<Vec<_>>());

    //三个数均不相同的情况，避免重复，按大小顺序排放
    for (&k, _) in &map {
        for (&j, _) in &map {
            let l = -k - j;
            if l < j && j < k && map.contains_key(&l) {
                res.push(vec![l, j, k]);
            }
        }
    }

    res
}

//417. 太平洋大西洋水流问题
pub fn pacific_atlantic(heights: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
    //思路：求流向太平洋的，再求流向大西洋的， 两者共同的即返回的结果
    //如何求流向太平洋的=> 广度或深度优先遍历即可。
    //思路：求流向太平洋的，再求流向大西洋的， 两者共同的即返回的结果
    //如何求流向太平洋的=> 广度或深度优先遍历即可。
    use std::collections::HashSet;
    let mut pacific = HashSet::new();
    let mut atlantic = HashSet::new();
    let (m, n) = (heights.len(), heights[0].len());
    let mut stack1 = vec![];
    let mut stack2 = vec![];
    for i in 0..m {
        stack1.push((i, 0));
        pacific.insert((i, 0));
        stack2.push((i, n - 1));
        atlantic.insert((i, n - 1));
    }
    for i in 0..n {
        stack1.push((0, i));
        pacific.insert((0, i));
        stack2.push((m - 1, i));
        atlantic.insert((m - 1, i));
    }

    let process = |pacific: &mut HashSet<(usize, usize)>, mut stack1: Vec<(usize, usize)>| {
        while let Some((x, y)) = stack1.pop() {
            let cur_height = heights[x][y];
            if x > 0 && heights[x - 1][y] >= cur_height && !pacific.contains(&(x - 1, y)) {
                stack1.push((x - 1, y));
                pacific.insert((x - 1, y));
            }
            if x < m - 1 && heights[x + 1][y] >= cur_height && !pacific.contains(&(x + 1, y)) {
                stack1.push((x + 1, y));
                pacific.insert((x + 1, y));
            }
            if y > 0 && heights[x][y - 1] >= cur_height && !pacific.contains(&(x, y - 1)) {
                stack1.push((x, y - 1));
                pacific.insert((x, y - 1));
            }
            if y < n - 1 && heights[x][y + 1] >= cur_height && !pacific.contains(&(x, y + 1)) {
                stack1.push((x, y + 1));
                pacific.insert((x, y + 1));
            }
        }
    };
    //计算太平洋Pacific
    process(&mut pacific, stack1);
    process(&mut atlantic, stack2);
    pacific
        .into_iter()
        .filter(|e| atlantic.contains(e))
        .map(|(x, y)| vec![x as i32, y as i32])
        .collect()
}

//130. 被围绕的区域
pub fn solve(board: &mut Vec<Vec<char>>) {
    // 思路：广度或深度优先遍历，从边缘为o的地方遍历，最后统一将内部的o变成X，外部遍历的改回来即可。
    const temp: char = 'T';
    let m = board.len();
    let n = board[0].len();
    let mut stack = vec![];
    for i in 0..m {
        if board[i][0] == 'O' {
            stack.push((i, 0));
            board[i][0] = temp;
        }
        if board[i][n - 1] == 'O' {
            stack.push((i, n - 1));
            board[i][n - 1] = temp;
        }
    }
    for i in 1..n - 1 {
        if board[0][i] == 'O' {
            stack.push((0, i));
            board[0][i] = temp;
        }
        if board[m - 1][i] == 'O' {
            stack.push((m - 1, i));
            board[m - 1][i] = temp;
        }
    }
    while let Some((x, y)) = stack.pop() {
        if x > 0 && board[x - 1][y] == 'O' {
            stack.push((x - 1, y));
            board[x - 1][y] = temp;
        }
        if x < m - 1 && board[x + 1][y] == 'O' {
            stack.push((x + 1, y));
            board[x + 1][y] = temp;
        }
        if y > 0 && board[x][y - 1] == 'O' {
            stack.push((x, y - 1));
            board[x][y - 1] = temp;
        }
        if y < n - 1 && board[x][y + 1] == 'O' {
            stack.push((x, y + 1));
            board[x][y + 1] = temp;
        }
    }
    //处理
    for i in 0..m {
        for j in 0..n {
            board[i][j] = match board[i][j] {
                'O' => 'X',
                'T' => 'O',
                _ => 'X'
            };
        }
    }
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
