fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //1051. 高度检查器
    pub fn height_checker(heights: Vec<i32>) -> i32 {
        let mut h2 = heights.clone();
        h2.sort_unstable();
        h2.into_iter()
            .zip(heights.into_iter())
            .filter(|&e| e.1 != e.0)
            .count() as i32
    }
}