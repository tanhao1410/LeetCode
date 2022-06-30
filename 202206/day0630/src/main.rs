fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //2215. 找出两数组的不同
    pub fn find_difference(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<Vec<i32>> {
        use std::collections::HashSet;
        let nums1 = nums1.into_iter().collect::<HashSet<_>>();
        let nums2 = nums2.into_iter().collect::<HashSet<_>>();
        let v1 = nums1.iter()
            .filter(|e| !nums2.contains(e))
            .map(|e| *e)
            .collect();
        let v2 = nums2.iter()
            .filter(|e| !nums1.contains(e))
            .map(|e| *e)
            .collect();
        vec![v1, v2]
    }

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