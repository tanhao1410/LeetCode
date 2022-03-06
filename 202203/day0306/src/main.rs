fn main() {
    println!("Hello, world!");
}

impl Solution {
    //215. 数组中的第K个最大元素
    pub fn find_kth_largest(mut nums: Vec<i32>, k: i32) -> i32 {
        //以第一个元素为开始，将大于它的放后面，小于它的放前面。然后递归的形式找出元素。
        if nums.len() == 1 {
            return nums[0];
        }
        let first = nums[0];
        let mut l = 0;
        let mut r = nums.len() - 1;
        while r >= l {
            while l < nums.len() && nums[l] >= first {
                l += 1;
            }
            while r > 0 && nums[r] <= first {
                r -= 1;
            }
            if r > l {
                let temp = nums[r];
                nums[r] = nums[l];
                nums[l] = temp;
            }
        }
        nums[0] = nums[r];
        nums[r] = first;
        //0-r之间的数都是大于等于first的
        if k == r as i32 + 1 {
            return first;
        } else if k < r as i32 + 1 {
            //说明在前面
            return Self::find_kth_largest((&nums[..r]).to_vec(), k);
        } else {
            return Self::find_kth_largest((&nums[r + 1..]).to_vec(), k - r as i32 - 1);
        }
    }
    //347. 前 K 个高频元素
    pub fn top_k_frequent(nums: Vec<i32>, k: i32) -> Vec<i32> {
        use std::collections::HashMap;
        let mut map = HashMap::new();
        for &num in &nums {
            let entry = map.entry(num).or_insert(0);
            *entry += 1;
        }
        let mut nums = map.into_iter().collect::<Vec<_>>();
        nums.sort_unstable_by_key(|&e| -e.1);
        nums.into_iter().take(k as usize).map(|e| e.0).collect()
    }
}

struct Solution;