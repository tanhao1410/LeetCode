

fn main() {
    println!("Hello, world!");
    Solution::find_nth_digit(11);
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

    //剑指 Offer 44. 数字序列中某一位的数字
    pub fn find_nth_digit(n: i32) -> i32 {

        //10->1 //1000000000 ->1
        if n < 10 {
            return n;
        }
        if n == 10{
            return 1
        }

        fn get_bit(mut n: i32, mut base: i32, mut index: i32) -> i32 {
            let mut res = n;
            while index >= 0 {
                res = n / base;
                n = n % base;
                base /= 10;
                index -= 1;
            }
            res
        }

        let mut dp:Vec<i64> = vec![0, 10];
        let mut max:i64 = 10;
        let mut i = 1;
        let mut ten_i = 1;
        while max < n as i64 {
            i += 1;
            ten_i *= 10;
            max = *dp.last().unwrap() + 9 * i * ten_i;
            dp.push(max);
        }
        //前面切割掉
        let new_index = n - dp[dp.len() - 2] as i32;
        //此时的数字的位数都是 i
        let skip_count = new_index / i as i32;
        let skip_index = skip_count + ten_i as i32;
        get_bit(skip_index, ten_i as i32, new_index % i as i32)
    }

}