fn main() {
    println!("Hello, world!");
}

pub struct Solution {}

impl Solution {
    //507. 完美数
    pub fn check_perfect_number(num: i32) -> bool {
        (1..10001).filter(|&i| i * i <= num && num % i == 0 && num != i)
            .fold(0, |p, q| p + q + num / q) == 2 * num
    }
}