fn main() {
    println!("Hello, world!");
}

pub struct Solution {}

impl Solution {

    //238. 除自身以外数组的乘积
    pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
        let mut res = vec![1;nums.len()];

        (1..nums.len()).for_each(|i|{
            res[i] = res[i-1] * nums[i];
        });

        let mut temp = res[res.len() - 1];
        (0..nums.len() - 1).rev().for_each(|i|{
            res[i] *= temp;
            temp *= nums[i];
        });

        res
    }

    //238. 除自身以外数组的乘积-使用除法
    pub fn product_except_self2(nums: Vec<i32>) -> Vec<i32> {
        let sum = nums.iter().fold(1,|p,q| p * q);
        nums.iter().map(|i| sum / i).collect()
    }

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