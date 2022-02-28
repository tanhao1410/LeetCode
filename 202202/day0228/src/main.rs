fn main() {
    println!("Hello, world!");
}

impl Solution {
    //1601. 最多可达成的换楼请求数目
    pub fn maximum_requests(n: i32, requests: Vec<Vec<i32>>) -> i32 {
        //使用java用递归解决了，用栈来解决呢？深度优先遍历
        //用一个数来表示哪些请求使用了或没使用，因为最多16个，所以，一个i32的数足矣
        //(请求的序号，是否要，已经要了的请求）
        let mut stack = vec![(0, 1, 1), (0, 0, 0)];
        //总共16个深度，每步有两种走法，选择或不选择
        let mut res = 0;
        while let Some((index, used, count)) = stack.pop() {
            //走到头了
            if index == requests.len() - 1 {
                let mut item = 0;
                let mut buildings = vec![0; n as usize];
                //还需要判断是否满足要求
                for i in 0..requests.len() {
                    //该位被使用了
                    if used & (1 << i) > 0 {
                        item += 1;
                        buildings[requests[i][0] as usize] -= 1;
                        buildings[requests[i][1] as usize] += 1;
                    }
                }
                if buildings.into_iter().all(|v| v == 0) {
                    res = res.max(item);
                }
            } else {
                stack.push((index + 1, used | (1 << index + 1), count + 1));
                stack.push((index + 1, used, count));
            }
        }
        res
    }
}

struct Solution;