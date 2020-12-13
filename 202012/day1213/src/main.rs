use std::collections::HashSet;

fn main() {
    println!("Hello, world!");
}

struct Solution{}

impl Solution {

    //169. 多数元素
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        //思路：摩尔投票法
        let (mut res,mut v ) = (nums[0],1);
        for i in 1..nums.len(){
            if v == 0{
                res = nums[i];
            }
            if nums[i] == res{
                v += 1;
            }else{
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

        if nums.is_empty(){
            return false
        }
        let mut nums = nums;
        nums.sort_unstable();
        for i in 0..nums.len()-1{
            if nums[i] == nums[i + 1]{
                return true
            }
        }
        false
    }
}