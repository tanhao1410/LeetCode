fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    ///747. 至少是其他数字两倍的最大数
    pub fn dominant_index(nums: Vec<i32>) -> i32 {
        match nums
            .iter()
            .enumerate()
            .fold((0, 0, 0), |(index, max, big), (i, &n)| //index,max,bigger
                match (n > max, n > big) {
                    (true, _) => (i, n, max),
                    (_, true) => (index, max, n),
                    _ => (index, max, big)
                },
            ) {
            (index, max, big) if max >= 2 * big => index as i32,
            _ => -1
        }
    }
}