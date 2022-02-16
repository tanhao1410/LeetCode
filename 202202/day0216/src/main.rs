use std::collections::BTreeMap;

fn main() {
    println!("Hello, world!");
    println!("{:?}", Solution::subsets_with_dup(vec![1, 2, 2]));
    println!("{}", Solution::shortest_bridge(vec![vec![0, 1], vec![1, 0]]));
}

struct Solution;

impl Solution {
    //75. 颜色分类
    pub fn sort_colors(nums: &mut Vec<i32>) {
        let (mut r, mut w, mut b) = (0, 0, 0);
        for &num in nums.iter() {
            match num {
                0 => r += 1,
                1 => w += 1,
                _ => b += 1
            }
        }
        for i in 0..nums.len() {
            if i < r {
                nums[i] = 0;
            } else if i < r + w {
                nums[i] = 1;
            } else {
                nums[i] = 2;
            }
        }
    }

    //1926. 迷宫中离入口最近的出口
    pub fn nearest_exit(mut maze: Vec<Vec<char>>, entrance: Vec<i32>) -> i32 {
        //先遍历，找到出口。
        use std::collections::{HashSet, VecDeque};
        let mut outer = HashSet::new();
        for i in 0..maze.len() {
            if maze[i][0] == '.' {
                outer.insert((i, 0));
            }
            if maze[i][maze[0].len() - 1] == '.' {
                outer.insert((i, maze[0].len() - 1));
            }
        }
        for i in 1..maze[0].len() - 1 {
            if maze[0][i] == '.' {
                outer.insert((0, i));
            }
            if maze[maze.len() - 1][i] == '.' {
                outer.insert((maze.len() - 1, i));
            }
        }
        outer.remove(&(entrance[0] as usize, entrance[1] as usize));
        if outer.len() == 0 {
            return -1;
        }
        //开始广度优先遍历
        let mut queue = VecDeque::new();
        queue.push_back((entrance[0] as usize, entrance[1] as usize));
        maze[entrance[0] as usize][entrance[1] as usize] = '+';
        let mut len = queue.len();
        let mut distance = 0;
        while len > 0 {
            for _ in 0..len() {
                let (x, y) = queue.pop_front().unwrap();
                //上下左右移动
                if outer.contains(&(x, y)) {
                    return distance;
                }
                if x > 0 && maze[x - 1][y] == '.' {
                    queue.push_back((x - 1, y));
                    maze[x - 1][y] = '+';
                }
                if x < maze.len() - 1 && maze[x + 1][y] == '.' {
                    queue.push_back((x + 1, y));
                    maze[x + 1][y] = '+';
                }
                if y > 0 && maze[x][y - 1] == '.' {
                    queue.push_back((x, y - 1));
                    maze[x][y - 1] = '+';
                }
                if y < maze[0].len() - 1 && maze[x][y + 1] == '.' {
                    queue.push_back((x, y + 1));
                    maze[x][y + 1] = '+';
                }
            }
            distance += 1;
            len = queue.len();
        }
        -1
    }

    //934. 最短的桥
    pub fn shortest_bridge(mut grid: Vec<Vec<i32>>) -> i32 {
        use std::collections::{HashSet, VecDeque};
        let mut queue = VecDeque::new();
        let mut flag = false;
        for i in 0..grid.len() {
            for j in 0..grid[0].len() {
                if grid[i][j] == 1 {
                    grid[i][j] = 2;
                    queue.push_back((i, j));
                    flag = true;
                    break;
                }
            }
            if flag {
                break;
            }
        }

        let push_queue = |grid: &mut Vec<Vec<i32>>, x: i32, y: i32, queue: &mut VecDeque<(usize, usize)>
                          , target: i32, target2: i32, change: i32| {
            if x >= 0 && x < grid.len() as i32 && y >= 0 && y < grid[0].len() as i32 {
                let x = x as usize;
                let y = y as usize;
                if grid[x][y] == target {
                    grid[x][y] = change;
                    queue.push_back((x, y));
                }
                if grid[x][y] == target2 {
                    queue.push_back((x, y));
                }
            }
        };

        //广度优先
        let mut len = queue.len();
        while len > 0 {
            for _ in 0..len {
                let (x, y) = queue.pop_front().unwrap();
                let x = x as i32;
                let y = y as i32;
                push_queue(&mut grid, x - 1, y, &mut queue, 1, 1, 2);
                push_queue(&mut grid, x + 1, y, &mut queue, 1, 1, 2);
                push_queue(&mut grid, x, y - 1, &mut queue, 1, 1, 2);
                push_queue(&mut grid, x, y + 1, &mut queue, 1, 1, 2);
            }
            len = queue.len();
        }
        for i in 0..grid.len() {
            for j in 0..grid[0].len() {
                if grid[i][j] == 2 {
                    grid[i][j] = 3;
                    queue.push_back((i, j));
                }
            }
        }
        //广度优先遍历
        let mut distance = 0;
        len = queue.len();
        loop {
            for _ in 0..len {
                let (x, y) = queue.pop_front().unwrap();
                //判断是否应该结束
                if grid[x][y] == 1 {
                    return distance - 1;
                }
                //如果它的上下周围有1，则返回,没有，则加入周边的0，代表翻转了
                push_queue(&mut grid, x as i32 + 1, y as i32, &mut queue, 0, 1, 3);
                push_queue(&mut grid, x as i32 - 1, y as i32, &mut queue, 0, 1, 3);
                push_queue(&mut grid, x as i32, y as i32 + 1, &mut queue, 0, 1, 3);
                push_queue(&mut grid, x as i32, y as i32 - 1, &mut queue, 0, 1, 3);
            }
            distance += 1;
            len = queue.len();
        }
    }

    //90. 子集 II
    pub fn subsets_with_dup(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut map = BTreeMap::new();
        for &num in &nums {
            let entry = map.entry(num).or_insert(0);
            *entry += 1;
        }
        let nums = map.into_iter().map(|(k, v)| (k, v)).collect::<Vec<(i32, i32)>>();
        Self::subsets_with_dup2(vec![], &nums)
    }

    fn subsets_with_dup2(mut pre: Vec<Vec<i32>>, nums: &[(i32, i32)]) -> Vec<Vec<i32>> {
        if nums.len() == 0 {
            return pre;
        }
        let mut new_pre = vec![];
        let (k, v) = nums[0];
        if pre.len() == 0 {
            for i in 0..=v {
                new_pre.push(vec![k; i as usize]);
            }
        } else {
            for item in &pre {
                for i in 1..=v {
                    let mut new_item = item.clone();
                    new_item.append(&mut vec![k; i as usize]);
                    new_pre.push(new_item);
                }
            }
            new_pre.append(&mut pre);
        }
        Self::subsets_with_dup2(new_pre, &nums[1..])
    }

    //78. 子集
    pub fn subsets(nums: Vec<i32>) -> Vec<Vec<i32>> {
        Self::subsets2(vec![], &nums)
    }

    fn subsets2(mut pre: Vec<Vec<i32>>, nums: &[i32]) -> Vec<Vec<i32>> {
        if nums.len() == 0 {
            return pre;
        }
        let mut new_pre = vec![];
        if pre.len() == 0 {
            new_pre.push(vec![]);
            new_pre.push(vec![nums[0]]);
        } else {
            for item in &pre {
                let mut new_item = item.clone();
                new_item.push(nums[0]);
                new_pre.push(new_item);
            }
            new_pre.append(&mut pre);
        }
        Self::subsets2(new_pre, &nums[1..])
    }
}