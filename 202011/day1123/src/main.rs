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
