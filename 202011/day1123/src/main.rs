fn main() {
    println!("Hello, world!");
}

//119. 杨辉三角 II
pub fn get_row(row_index: i32) -> Vec<i32> {
    let mut res = vec![1; row_index as usize];
    let mut temp = vec![1; row_index as usize];
    for i in 1..row_index {
        for j in 0..i {
            if j > 0 {
                res[j as usize] += temp[j as usize - 1];
            }
        }
        for j in 0..i as usize {
            temp[j] = res[j]
        }
    }
    res
}

//121. 买卖股票的最佳时机
pub fn max_profit(mut prices: Vec<i32>) -> i32 {
    //思路,修改数组，从后往前修改，遍历后，把它的值改为从本位置算起的最大的值
    if prices.len() < 2 {
        return 0;
    }
    let mut res = 0;
    for i in (0..prices.len() - 1).rev() {
        if res < prices[i + 1] - prices[i] {
            res = prices[i + 1] - prices[i];
        }
        if prices[i] < prices[i + 1] {
            prices[i] = prices[i + 1]
        }
    }
    res
}

//115. 不同的子序列
pub fn num_distinct(s: String, t: String) -> i32 {
    //思路2：先找到第一个字母可以的方式，然后其余的进行递归求即可。

    let mut res = 0;
    let mut locations = vec![];
    let s_bytes = s.as_bytes();
    let t_bytes = t.as_bytes();
    for i in 0..s.len(){
        if s_bytes[i] == t_bytes[0]{
            locations.push(i);
        }
    }
    if t.len() == 1{
        return locations.len() as i32;
    }

    while !locations.is_empty(){
        let mut start = 1;
        let (_, new_t) = t.split_at(start);
        let (_, new_s) = s.split_at(locations.pop().unwrap()+1);
        println!("{}-{}", new_s, new_t);
        res += Solution::num_distinct(new_s.to_string(), new_t.to_string());
    }
    res
}
