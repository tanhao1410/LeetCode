fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //926. 将字符串翻转到单调递增
    pub fn min_flips_mono_incr(s: String) -> i32 {
        let bytes = s.as_bytes();
        //前面变成0所需的次数
        let mut pre_zero = vec![0; s.len()];
        //后面变成1所需的次数
        let mut pro_one = vec![0; s.len()];
        for i in 0..s.len() {
            pre_zero[i] = *pre_zero.get(i - 1).unwrap_or(&0);
            pro_one[s.len() - 1 - i] = *pro_one.get(s.len() - i).unwrap_or(&0);
            if bytes[i] == b'1' {
                pre_zero[i] += 1;
            }
            if bytes[s.len() - 1 - i] == b'0' {
                pro_one[s.len() - 1 - i] += 1;
            }
        }
        pre_zero
            .into_iter()
            .zip(pro_one.into_iter())
            .map(|(i, j)| i + j)
            .min()
            .unwrap()
            - 1
    }
    //476. 数字的补数
    pub fn find_complement(mut num: i32) -> i32 {
        let mut bits = vec![];
        while num > 0 {
            bits.push(1 - (num & 1));
            num >>= 1;
        }

        // let mut res = 0;
        // for i in bits.into_iter().rev() {
        //     res <<= 1;
        //     res += i;
        // }
        // res
        bits.into_iter()
            .rev()
            .fold(0, |mut res, bit| {
                res <<= 1;
                res + bit
            })
    }
}