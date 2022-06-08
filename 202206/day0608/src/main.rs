fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //1037. 有效的回旋镖
    pub fn is_boomerang(points: Vec<Vec<i32>>) -> bool {
        (points[2].0 - points[0].0) * (points[2].1 - points[0].1) != (points[1].0 - points[0].0) * (points[1].1 - points[0].1)
    }
}