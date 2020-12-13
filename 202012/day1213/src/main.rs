use std::collections::HashSet;

fn main() {
    println!("Hello, world!");
}

struct Solution{}

impl Solution {
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