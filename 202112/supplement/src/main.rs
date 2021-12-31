fn main() {
    println!("Hello, world!");
}

struct Solution{}

impl Solution {
    //12-31-每日一题：507. 完美数
    pub fn check_perfect_number(num: i32) -> bool {
        (1..)
            .take_while(|&n| n * n <= num)
            .filter(|&n|num % n == 0 && num != n)
            .map(|n| n + num / n)
            .sum::<i32>()
            == 2 * num
    }
}
