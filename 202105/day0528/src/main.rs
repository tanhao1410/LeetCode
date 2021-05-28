fn main() {
    println!("Hello, world!");
}

//412. Fizz Buzz
pub fn fizz_buzz(n: i32) -> Vec<String> {
    //3的倍数 fiz,5
    (1..n + 1).map(|num| match num % 15 {
        0 => "FizzBuzz".to_string(),
        3 | 6 | 9 | 12 => "Fizz".to_string(),
        5 | 10 => "Buzz".to_string(),
        _ => num.to_string()
    }).collect()
}

//477. 汉明距离总和
pub fn total_hamming_distance(nums: Vec<i32>) -> i32 {
    //思路：记录每一位0,1的总个数，该位总和为其相乘，总的相加即得答案
    nums.iter().fold(vec![(0, 0); 32], |mut count, &num| {
        for i in 0..32 {
            match num & (1 << i) {
                0 => count[i].0 += 1,
                _ => count[i].1 += 1
            }
        }
        count
    }).iter()
        .inspect(|&&item| println!("{:?}", item))
        .fold(0, |mut res, &(zero, one)| res + zero * one)
}