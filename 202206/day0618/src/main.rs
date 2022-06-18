use std::cmp::Ordering;

fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //1985. 找出数组中的第 K 大整数
    pub fn kth_largest_number(nums: Vec<String>, k: i32) -> String {
        let mut nums = nums;
        nums.sort_unstable_by(|pre, pro| {
            if pre.len() < pro.len() {
                return Ordering::Greater;
            }
            if Self::less_equal(pre, pro) {
                return Ordering::Greater;
            }
            Ordering::Less
        });
        nums[k as usize - 1].clone()
    }

    fn less_equal<T: AsRef<str>>(s1: T, s2: T) -> bool {
        let s1 = s1.as_ref();
        let s2 = s2.as_ref();
        //先比较长度，长的大
        if s1.len() != s2.len() {
            return s1.len() < s2.len();
        }
        let s1_bytes = s1.as_bytes();
        let s2_bytes = s2.as_bytes();
        //长度相同时，依次比较
        for i in 0..s1.len() {
            if s1_bytes[i] < s2_bytes[i] {
                return true;
            } else if s1_bytes[i] > s2_bytes[i] {
                return false;
            }
        }
        true
    }

    //1828. 统计一个圆中点的数目
    pub fn count_points(points: Vec<Vec<i32>>, queries: Vec<Vec<i32>>) -> Vec<i32> {
        queries.into_iter()
            .map(|q| points
                .iter()
                .filter(|p| (p[0] - q[0]).pow(2) + (p[1] - q[1]).pow(2) <= q[2].pow(2))
                .count() as i32)
            .collect()
    }
}