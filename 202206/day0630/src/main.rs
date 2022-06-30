fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //1175. 质数排列
    pub fn num_prime_arrangements(n: i32) -> i32 {
        const MOD: i64 = 1000000007;
        let is_prime = |num: &i32| {
            for i in 2..*num {
                if *num % i == 0 {
                    return false;
                }
            }
            true && *num > 1
        };
        let prime_count = (1..n).filter(is_prime).count() as i64;
        let mut res = 1i64;
        for i in 2..=prime_count {
            res = res * i % MOD;
        }
        for i in 2..=n - prime_count {
            res = res * i % MOD;
        }
        res as i32
    }
}