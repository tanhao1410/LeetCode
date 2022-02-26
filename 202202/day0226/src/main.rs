fn main() {
    println!("Hello, world!");
}

impl Solution {
    //2016. 增量元素之间的最大差值
    pub fn maximum_difference(nums: Vec<i32>) -> i32 {
        let mut max = 0;
        let mut res = -1;
        for i in nums.into_iter().rev() {
            if max > i {
                res = res.max(max - i);
            } else {
                max = i;
            }
        }
        res
    }
}

struct Solution;