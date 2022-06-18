fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
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