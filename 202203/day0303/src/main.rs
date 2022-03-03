fn main() {
    println!("Hello, world!");
}

impl Solution {
    //258. 各位相加
    pub fn add_digits(num: i32) -> i32 {
        if num < 9 {
            return num;
        }
        if num % 9 == 0 {
            return 9;
        }
        return num % 9;
    }
}

struct Solution;