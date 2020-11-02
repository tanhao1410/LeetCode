fn main() {
    println!("Hello, world!");
    println!("{}", Solution::valid_mountain_array(vec![0, 3, 2, 1]))
}

struct Solution {}

impl Solution {
    //每日一题：941. 有效的山脉数组
    pub fn valid_mountain_array(a: Vec<i32>) -> bool {
        if a.len() < 3 || a[1] <= a[0] {
            return false;
        }
        let mut pre = a[0];
        //先递增，后递减
        let mut is_increase = true;
        for i in 1..a.len() {
            if is_increase && a[i] > pre {} else if a[i] < pre {
                is_increase = false;
            } else {
                return false;
            }
            pre = a[i];
        }
        !is_increase
    }
}