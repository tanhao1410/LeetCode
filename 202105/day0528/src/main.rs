fn main() {
    println!("Hello, world!");
}

//451. 根据字符出现频率排序
pub fn frequency_sort(s: String) -> String {
    //统计字母出现的次数，排序，然后生成新的字符串
    let mut char_count: Vec<i32> = vec![0; 128];
    s.as_bytes().iter().for_each(|&c| char_count[c as usize] += 1);
    let mut char_count = char_count.iter().enumerate().collect::<Vec<(usize, &i32)>>();
    char_count.sort_by_key(|&(c, v)| -v);
    char_count.iter().fold(String::new(), |res, &(c, &count)| res + &String::from_utf8(vec![c as u8; count as usize]).unwrap())
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