fn main() {
    println!("Hello, world!");
}

impl Solution {
    //66. 加一
    pub fn plus_one(mut digits: Vec<i32>) -> Vec<i32> {
        //主要是判断是有无进位
        let mut flag = 1;
        for digit in digits.iter_mut().rev() {
            if flag == 0 {
                break;
            }
            *digit += flag;
            flag = *digit / 10;
            *digit %= 10;
        }
        if flag == 1 {
            let mut res = vec![1];
            res.append(&mut digits);
            return res;
        }
        digits
    }
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