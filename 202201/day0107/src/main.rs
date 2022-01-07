fn main() {
    println!("{}", Solution::max_depth("((1+2))+(((2)))".to_string()));
}

struct Solution {}

impl Solution {
    pub fn max_depth(s: String) -> i32 {
        s.as_bytes().iter().fold((0, 0), |(deep, max), &b| {
            match b {
                b'(' => (deep + 1, max.max(deep + 1)),
                b')' => (deep - 1, max),
                _ => (deep, max)
            }
        }).1
    }
}