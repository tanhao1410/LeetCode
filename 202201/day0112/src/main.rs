use std::cmp::Ordering;

fn main() {}

struct Solution;

impl Solution {
    pub fn increasing_triplet(nums: Vec<i32>) -> bool {
        //每一次迭代只需要一个或两个更小的数即可。开始时是三个数，
        nums
            .iter()
            .fold((None, None, None, false), |pre, &num|
                match pre {
                    (.., true) => pre,
                    (None, .., false) => (Some(num), pre.1, pre.2, pre.3),
                    (Some(a), None, ..) => match a.cmp(&num) {
                        Ordering::Less => (Some(a), Some(num), None, false),
                        Ordering::Greater => (Some(num), None, None, false),
                        _ => pre
                    }
                    (Some(a), Some(b), Some(c), ..) => {
                        //三个数都有值了
                        if num > b {
                            (None, None, None, true)
                        } else if num > a {
                            (Some(a), Some(num), Some(c), false)
                        } else if num > c {
                            (Some(c), Some(num), None, false)
                        } else {
                            (Some(a), Some(b), Some(c), false)
                        }
                    }
                    (Some(a), Some(b), ..) => {
                        //第三个数没值
                        if num < a {
                            (Some(a), Some(b), Some(num), false)
                        } else if num < b && num > a {
                            (Some(a), Some(num), None, false)
                        } else if num > b {
                            return (None, None, None, true);
                        } else {
                            (Some(a), Some(b), None, false)
                        }
                    }
                },
            )
            .3
    }
}