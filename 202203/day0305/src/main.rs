fn main() {
    println!("Hello, world!");
}

impl Solution {
    //503. 下一个更大元素 II
    pub fn next_greater_elements(nums: Vec<i32>) -> Vec<i32> {
        let mut stack = vec![];
        let mut res = vec![0; nums.len()];
        //单调栈,先进行一圈，好让后面的知道有哪些比它大
        for i in (0..nums.len()).rev() {
            while !stack.is_empty() && stack[stack.len() - 1] <= nums[i] {
                stack.pop();
            }
            stack.push(nums[i]);
        }
        for i in (0..nums.len()).rev() {
            while !stack.is_empty() && stack[stack.len() - 1] <= nums[i] {
                stack.pop();
            }
            if stack.is_empty() {
                res[i] = -1;
            } else {
                res[i] = stack[stack.len() - 1];
            }
            stack.push(nums[i]);
        }
        res
    }
    //521. 最长特殊序列 Ⅰ
    pub fn find_lu_slength(a: String, b: String) -> i32 {
        if a == b {
            -1
        } else {
            a.len().max(b.len()) as i32
        }
    }
}

struct Solution;