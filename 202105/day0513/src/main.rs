fn main() {
    println!("Hello, world!");
}

pub struct Solution {}

impl Solution {
    //279. 完全平方数
    pub fn num_squares(n: i32) -> i32 {
        //先得到所有的完全平方数
        let mut dp = vec![0; n as usize + 1];
        (1..n + 1).for_each(|i| {
            let mut min_cur = i32::MAX;
            for j in (1..101) {
                if j * j > i {
                    break;
                }
                min_cur = min_cur.min(dp[(i - j * j) as usize] + 1);
            }
            dp[i as usize] = min_cur;
        });
        dp[n as usize]
    }
}