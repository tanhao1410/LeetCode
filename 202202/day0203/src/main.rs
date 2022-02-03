fn main() {
    println!("Hello, world!");
}

//136. 只出现一次的数字
pub fn single_number(nums: Vec<i32>) -> i32 {
    nums
        .iter()
        .fold(0,|p,&r|p ^ r)
}

//190. 颠倒二进制位
pub fn reverse_bits(x: u32) -> u32 {
    let mut bytes = vec![0; 32];
    for i in 0..32 {
        bytes[i] = (x & (1 << i)) >> i;
    }
    let mut res = 0;
    for i in 0..32 {
        res += (bytes[i] << (31 - i));
    }
    res
}
