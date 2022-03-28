fn main() {
    println!("Hello, world!");
}

impl Solution {
    //693. 交替位二进制数
    pub fn has_alternating_bits(n: i32) -> bool {
        let mut bits = vec![0; 32];
        for i in 0..32 {
            bits[i] = n & (1 << i);
        }
        let mut i = 31;
        while i >= 0 && bits[i] == 0 {
            i -= 1;
        }
        while i >= 0 {
            if bits[i] == bits[i + 1] {
                return false;
            }
            i -= 1;
        }
        true
    }
}

struct Solution;