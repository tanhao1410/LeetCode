fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //1577 数的平方等于两数乘积的方法数
    pub fn num_triplets(nums1: Vec<i32>, nums2: Vec<i32>) -> i32 {
        let triplets = |nums1: &[i32], nums2: &[i32]| -> i32{
            use std::collections::HashMap;
            let mut map: HashMap<i32, i32> = HashMap::new();
            for num in nums2 {
                let entry = map.entry(*num).or_insert(0);
                *entry += 1;
            }
            let mut res = 0;
            for num in nums1 {
                //数字过大怎么办？
                let num = *num as i64 * *num as i64;
                for (k, v) in map.iter() {
                    if (num % *k as i64) == 0 && map.contains_key(&((num / *k as i64) as i32)) {
                        if num / *k as i64 == *k as i64 {
                            res += v * (v - 1) / 2;
                        } else if *k as i64 > num / *k as i64 {
                            res += v * map.get(&((num / *k as i64) as i32)).unwrap();
                        }
                    }
                }
            }
            res
        };

        triplets(&nums2, &nums1) + triplets(&nums1, &nums2)
    }

    //1089. 复写零
    pub fn duplicate_zeros(arr: &mut Vec<i32>) {
        let arr2 = arr.clone();
        let mut index = 0;
        let mut index2 = 0;
        while index < arr.len() {
            if arr2[index2] == 0 {
                arr[index] = 0;
                if index + 1 < arr.len() {
                    arr[index + 1] = 0;
                }
                index += 2;
            } else {
                arr[index] = arr2[index2];
                index += 1;
            }
            index2 += 1;
        }
    }
}