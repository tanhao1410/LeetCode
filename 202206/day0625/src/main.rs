fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //剑指 Offer II 091. 粉刷房子
    pub fn min_cost(costs: Vec<Vec<i32>>) -> i32 {
        let mut dp = costs.clone();

        for x in 1..costs.len() {
            for y in 0..3 {
                //对于x号房，涂成 y号颜色，最实惠的方式是：
                //前面的不能涂成y 即可
                let mut min_cost = i32::MAX;
                for pre_y in 0..3 {
                    if pre_y != y {
                        min_cost = min_cost.min(dp[x - 1][pre_y]);
                    }
                }
                dp[x][y] += min_cost;
            }
        }

        *dp[dp.len() - 1].iter().min().unwrap()
    }
}