fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //1903. 字符串中的最大奇数
    pub fn largest_odd_number(num: String) -> String {
        let bytes = num.as_bytes();
        let mut last_odd = bytes.len() as i32 - 1;
        while last_odd >= 0 {
            if (bytes[last_odd as usize] - b'0') % 2 == 1 {
                break;
            }
            last_odd -= 1;
        }
        String::from_utf8_lossy(&bytes[..(last_odd + 1) as usize]).to_string()
    }

    //1561. 你可以获得的最大硬币数目
    pub fn max_coins(piles: Vec<i32>) -> i32 {
        let mut piles = piles;
        piles.sort_unstable();
        let n = piles.len();
        piles.into_iter()
            .skip(n / 3)
            .enumerate()
            .filter(|(i, _)| i % 2 == 0)
            .map(|(_, n)| n)
            .sum()
    }

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