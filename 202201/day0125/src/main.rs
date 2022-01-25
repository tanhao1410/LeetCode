fn main() {
    println!("Hello, world!");
}

//1688. 比赛中的配对次数
pub fn number_of_matches(mut n: i32) -> i32 {
    let mut res = 0;
    while n > 1{
        res += n / 2;
        n = n - n / 2;
    }
    res
}
