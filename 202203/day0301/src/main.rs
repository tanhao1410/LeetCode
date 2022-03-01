fn main() {
    println!("Hello, world!");
}

impl Solution {
    //1042. 不邻接植花
    pub fn garden_no_adj(n: i32, paths: Vec<Vec<i32>>) -> Vec<i32> {
        use std::collections::VecDeque;
        let mut path = vec![vec![];n as usize];
        for p in &paths{
            path[p[0] as usize - 1].push(p[1] as usize - 1);
            path[p[1] as usize - 1].push(p[0] as usize - 1);
        }
        let mut queue = VecDeque::new();
        let mut res = vec![0;n as usize];
        for i in 0..n as usize{
            if res[i] == 0{
                queue.push_back(i);
                while !queue.is_empty(){
                    let len = queue.len();
                    for _ in 0..len{
                        let cur = queue.pop_front().unwrap();
                        let mut temp = vec![0;4];
                        //应该用哪种颜色
                        for &arround in &path[cur]{
                            if res[arround] == 0{
                                queue.push_back(arround);
                            }else{
                                temp[res[arround] as usize - 1] = 1;
                            }
                        }
                        for i in 0..4{
                            if temp[i] == 0{
                                res[cur] = i as i32 + 1;
                                break;
                            }
                        }
                    }
                }
            }
        }
        res
    }
    //58. 最后一个单词的长度
    pub fn length_of_last_word(s: String) -> i32 {
        let bytes = s.as_bytes();
        let mut end = s.len() - 1;
        while end >= 0 && bytes[end] == b' ' {
            end -= 1;
        }
        let mut start = end;
        while start >= 0 && bytes[start] != b' ' {
            start -= 1;
        }
        (end - start) as i32
    }
    //739. 每日温度
    pub fn daily_temperatures(temperatures: Vec<i32>) -> Vec<i32> {
        let mut res = vec![0; temperatures.len()];
        let mut stack: Vec<(i32, usize)> = vec![];
        //单调栈,栈中存放后面比它大的数及坐标
        for i in (0..temperatures.len()).rev() {
            //如果栈顶元素比当前元素小或等于，则弹出元素。
            while stack.len() > 0 && stack[stack.len() - 1].0 <= temperatures[i] {
                stack.pop();
            }
            //如果栈不为空，则说明存在了一个比当前大的数，且是最近的
            if stack.len() > 0 {
                res[i] = stack[stack.len() - 1].1 as i32 - i as i32;
            }
            stack.push((temperatures[i], i));
        }
        res
    }
    //6. Z 字形变换
    pub fn convert(s: String, num_rows: i32) -> String {
        if num_rows == 1 {
            return s;
        }
        let mut m = vec![vec![]; num_rows as usize];
        let mut x = 0;
        let mut y = 0;
        let mut is_down = true;//方向有往下和往右上
        let bytes = s.as_bytes();
        for &b in bytes {
            m[x].push(b);
            if is_down {
                if x < num_rows as usize - 1 {
                    x += 1;
                } else {
                    is_down = false;
                    x -= 1;
                    y += 1;
                }
            } else {
                if x == 0 {
                    is_down = true;
                    x += 1;
                } else {
                    x -= 1;
                    y += 1;
                }
            }
        }
        String::from_utf8(m
            .into_iter()
            .flat_map(|v| v.into_iter())
            .collect::<Vec<_>>()).unwrap()
    }
}

struct Solution;