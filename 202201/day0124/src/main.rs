use std::collections::{VecDeque, HashSet};

fn main() {
    println!("Hello, world!");
    println!("{}", second_minimum(5, vec![vec![1, 2], vec![1, 3], vec![1, 4], vec![3, 4], vec![4, 5]], 3, 5));
}

//2045. 到达目的地的第二短时间
pub fn second_minimum(n: i32, edges: Vec<Vec<i32>>, time: i32, change: i32) -> i32 {
    //思路：结构，点->能到达的点集合。
    let mut map = vec![HashSet::new(); n as usize];
    for i in 0..edges.len() {
        let source = edges[i][0] as usize - 1;
        let dest = edges[i][1] as usize - 1;
        //起点能到达的位置 添加进来
        map[source].insert(dest);
        map[dest].insert(source);
    }
    //广度优先，从1到n
    let mut cur_time = 0;
    let mut queue = VecDeque::new();
    let mut fast_reach = false;
    queue.push_back(0usize);
    let mut reach_count = vec![0; n as usize];
    loop {
        //从队列中弹出能达到的位置，每次新到达的位置
        //步数+1
        // res += 1;

        //判断当前是否是红灯，是否需要等待，怎么判断呢，0123/4567
        if cur_time / change % 2 == 0 {//是绿灯
            //走一步
            cur_time += time;
        } else {
            //需要等待的时长
            cur_time += (change - (cur_time % change));
            cur_time += time;
        }

        let mut can_reach = HashSet::new();
        while let Some(p) = queue.pop_front() {
            //本次新达到的位置
            can_reach.extend(map[p].iter());
            //map[p].iter().for_each(|e|{can_reach.insert(*e);});
            //看能到达的里面是否包含终点，如果包含终点，第一次的时候，是最短的，下一次就是最长的了
        }
        //如果第一次已经访问已经出现了，并且本次也能达到，返回第二近的结果
        if fast_reach && can_reach.contains(&(n as usize - 1)) {
            return cur_time;
        }
        if can_reach.contains(&(n as usize - 1)) {
            fast_reach = true;
        }

        can_reach.iter().for_each(|&p| reach_count[p] += 1);
        //每一次都直接加入所有的节点。最终会导致时间超时的发生，需要进行优化。
        // 一个节点，在本次到达后，后面可能还会再来一次，但不会接连进入第三次
        queue.extend(can_reach.iter().filter(|&&p| reach_count[p] < 3));
        //can_reach.iter().for_each(|&e|{queue.push_back(e);});
    }
}

//122. 买卖股票的最佳时机 II
pub fn max_profit2(prices: Vec<i32>) -> i32 {
    //思路:什么时候买，买一个后面有比它大的，什么时候卖：已经买了，比买入贵，后面又降价了。
    let mut res = 0;
    let mut buy = false;
    let mut buy_num = 0;
    for i in 0..prices.len() - 1 {
        if !buy && prices[i + 1] > prices[i] {
            //如果还没买，可以买
            buy = true;
            buy_num = prices[i];
        }
        //什么时候可以卖呢，只有买了才能卖，只有，第二天降价了可以卖
        if buy && prices[i + 1] < prices[i] {
            buy = false;
            res += prices[i] - buy_num;
        }
    }
    if buy {
        res += prices[prices.len() - 1] - buy_num;
    }
    res
}

//121. 买卖股票的最佳时机
pub fn max_profit(prices: Vec<i32>) -> i32 {
    prices
        .iter()
        .fold((i32::MAX, 0), |(min, res), &cur| (min.min(cur), res.max(cur - min)))
        .1

    //思路：dp记录前面最便宜的价格
    // let mut dp = prices[0];
    // let mut res = 0;
    // for i in 1..prices.len(){
    //     res = res.max(prices[i] - dp);
    //     dp = dp.min(prices[i]);
    // }
    // res
}

//1014. 最佳观光组合
pub fn max_score_sightseeing_pair(values: Vec<i32>) -> i32 {
    //思路:dp[i]以values[i]结尾的组合，最大的值
    //dp[i] = dp[i-1] - 1 - values[i - 1] + values[i] 或 values[i] + values[i -1] - 1
    let mut dp = values[0] + values[1] - 1;
    let mut res = dp;
    for i in 2..values.len() {
        if dp > 2 * values[i - 1] {
            dp = dp - 1 - values[i - 1] + values[i];
        } else {
            dp = values[i - 1] - 1 + values[i];
        }
        res = res.max(dp);
    }
    res
}

//557. 反转字符串中的单词 III
pub fn reverse_words(s: String) -> String {
    s
        .split(" ")
        .map(|s| s.chars().rev().collect::<String>())
        .collect::<Vec<_>>()
        .join(" ")
}

//519. 随机翻转矩阵
pub fn reverse_string(s: &mut Vec<char>) {
    let mut i = 0;
    let mut j = s.len() - 1;
    while j > i {
        let temp = s[i];
        s[i] = s[j];
        s[j] = temp;
        i += 1;
        j -= 1;
    }
}