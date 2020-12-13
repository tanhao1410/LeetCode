use std::collections::HashSet;

fn main() {
    println!("Hello, world!");
}

struct Solution {}

impl Solution {
    //168. Excel表列名称
    pub fn convert_to_title(n: i32) -> String {
        //1->a
        let mut res = String::new();
        //相当于26进制
        let mut n = n;
        while n > 26 {
            let t = (n % 26) as u8;
            if t == 0 {
                res.insert(0, 'Z');
                n = n / 26 - 1;
            } else {
                res.insert(0, ('A' as u8 + t - 1) as char);
                n = n / 26;
            }
        }
        res.insert(0, ('A' as u8 + n as u8 - 1) as char);
        res
    }

    //167. 两数之和 II - 输入有序数组
    pub fn two_sum(numbers: Vec<i32>, target: i32) -> Vec<i32> {
        //思路：双指针法，分别从前和尾开始向中间走
        let (mut i, mut j) = (0, numbers.len() - 1);
        while numbers[i] + numbers[j] != target {
            if numbers[i] + numbers[j] > target {
                j -= 1;
            } else {
                i += 1;
            }
        }
        vec![i as i32 + 1, j as i32 + 1]
    }

    //169. 多数元素
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        //思路：摩尔投票法
        let (mut res, mut v) = (nums[0], 1);
        for i in 1..nums.len() {
            if v == 0 {
                res = nums[i];
            }
            if nums[i] == res {
                v += 1;
            } else {
                v -= 1;
            }
        }
        res
    }

    //每日一题：217. 存在重复元素
    pub fn contains_duplicate(nums: Vec<i32>) -> bool {

        //hashSet方法：
        // let mut nums_set = HashSet::new();
        // for i in nums{
        //     if nums_set.contains(&i){
        //         return true
        //     }
        //     nums_set.insert(i);
        // }
        // false

        if nums.is_empty() {
            return false;
        }
        let mut nums = nums;
        nums.sort_unstable();
        for i in 0..nums.len() - 1 {
            if nums[i] == nums[i + 1] {
                return true;
            }
        }
        false
    }
}