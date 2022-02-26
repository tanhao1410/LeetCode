fn main() {
    println!("Hello, world!");
}

impl Solution {
    //456. 132 模式
    pub fn find132pattern(nums: Vec<i32>) -> bool {
        if nums.len() < 3 {
            return false;
        }
        //思路：单调栈-找后面比自己小的最大值，前面的最小值,不包括自己
        let mut pre_min = vec![1000000001; nums.len()];
        for i in 1..nums.len() {
            pre_min[i] = pre_min[i - 1].min(nums[i - 1]);
        }
        //单调栈
        let mut stack = vec![*nums.last().unwrap()];
        for i in (1..nums.len() - 1).rev() {
            let cur = nums[i];
            if stack.is_empty() {
                stack.push(cur);
            } else {
                //判断它后面最大数是多少
                let min = pre_min[i];
                let mut top = *stack.last().unwrap();
                while !stack.is_empty() && *stack.last().unwrap() < cur {
                    top = stack.pop().unwrap();
                }
                if top > min && cur > top {
                    return true;
                }
                stack.push(cur);
            }
        }
        false
    }
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