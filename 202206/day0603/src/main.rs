fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //829. 连续整数求和
    pub fn consecutive_numbers_sum(n: i32) -> i32 {
        (1..)
            .take_while(|&i| i * (i - 1) / 2 < n)
            .filter(|&i| (n - ((i + 1) % 2 * (i - 1) * i / 2)) % i == 0 && (n - ((i + 1) % 2) * (i - 1) * i / 2) / i > 0)
            .count() as i32
    }
}