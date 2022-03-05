fn main() {
    println!("Hello, world!");
}

impl Solution {
    //42. 接雨水
    pub fn trap(height: Vec<i32>) -> i32 {
        //它前面的最大值，它后面的最大值
        let mut pre_max = vec![0; height.len()];
        let mut tail_max = vec![0; height.len()];
        for i in 0..height.len() {
            if i > 0 {
                pre_max[i] = pre_max[i - 1].max(height[i]);
                tail_max[height.len() - 1 - i] = tail_max[height.len() - i].max(height[height.len() - 1 - i]);
            } else {
                pre_max[i] = height[i];
                tail_max[height.len() - 1 - i] = height[height.len() - 1 - i];
            }
        }
        //每一个高度柱能放多少水呢，取决于，它左右两边的最高值的最小值
        pre_max.into_iter()
            .zip(tail_max.into_iter())
            .map(|e| e.0.min(e.1))
            .zip(height.into_iter())
            .map(|v| v.0 - v.1)
            .sum()
    }
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