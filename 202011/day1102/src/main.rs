

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

    //面试题 01.02. 判定是否互为字符重排
    pub fn check_permutation(s1: String, s2: String) -> bool {
        if s1.len() != s2.len() {
            return false;
        }
        //采用map的形式
        let mut m1 = std::collections::HashMap::new();
        let mut m2 = std::collections::HashMap::new();
        for i in 0..s1.len() {
            if m1.contains_key(&s1.chars().nth(i).unwrap()){
                m1.insert(s1.chars().nth(i).unwrap(),m1.get(&s1.chars().nth(i).unwrap()).unwrap() + 1);
            }else{
                m1.insert(s1.chars().nth(i).unwrap(),1);
            }

            if m2.contains_key(&s2.chars().nth(i).unwrap()){
                m2.insert(s2.chars().nth(i).unwrap(),m2.get(&s1.chars().nth(i).unwrap()).unwrap() + 1);
            }else{
                m2.insert(s2.chars().nth(i).unwrap(),1);
            }

        }
        if m1.len() != m2.len() {
            return false;
        }
        for i in m2.into_iter() {
            match  m1.get(&i.0){
                None => return false,
                Some(&j) =>{
                    if i.1 != j{
                        return false
                    }
                }
            }
        }
        true
    }

}