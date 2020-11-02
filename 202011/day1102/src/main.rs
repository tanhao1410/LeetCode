

fn main() {
    println!("Hello, world!");
}

struct Solution{}

impl Solution {
    //每日一题：349. 两个数组的交集
    pub fn intersection(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
        //思路：用一个hashset来存储1中的数据，然后，遍历nums2，每次拿数据的时候，按顺序放入结果 集即可避免重复的出现。或者放入到set中，然后，再输出为vec
        use std::collections::HashSet;
        let (mut res,mut set1,mut res_set) = (vec![],HashSet::new(),HashSet::new());
        for i  in nums1{
            set1.insert(i);
        }
        for i in nums2{
            if set1.contains(&i){
                res_set.insert(i);
            }
        }
        for i in res_set{
            res.push(i);
        }
        res
    }
}