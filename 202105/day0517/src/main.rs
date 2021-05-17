fn main() {
    println!("Hello, world!");
}

pub struct Solution {}

impl Solution {

    //405. 数字转换为十六进制数
    pub fn to_hex(num: i32) -> String {
        if num == 0 {
            return "0".to_string();
        }

        let mut bi_vec = vec![false; 32];
        let mut maske = 1;
        for i in 0..32 {
            bi_vec[31 - i] = num & maske != 0;
            maske <<= 1;
        }

        let mut ox_vecc = vec![];
        for i in 0..8 {
            let mut n = 0;
            for j in 0..4 {
                if bi_vec[4 * i + j] {
                    n += 1 << (3 - j)
                }
            }
            if n < 10 {
                ox_vecc.push('0' as u8 + n as u8);
            } else {
                ox_vecc.push('a' as u8 + n as u8 - 10);
            }
        }
        String::from_utf8(ox_vecc).unwrap().trim_start_matches('0').to_string()
    }

    //507. 完美数
    pub fn check_perfect_number(num: i32) -> bool {
        (1..10001).filter(|&i| i * i <= num && num % i == 0 && num != i)
            .fold(0, |p, q| p + q + num / q) == 2 * num
    }
}