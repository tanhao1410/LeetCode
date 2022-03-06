fn main() {
    println!("Hello, world!");
}

impl Solution {
    //2100. 适合打劫银行的日子
    pub fn good_days_to_rob_bank(security: Vec<i32>, time: i32) -> Vec<i32> {
        //down[i] up[i]前面非递增的天数，后面非递减的数量
        let mut down = vec![0; security.len()];
        let mut up = vec![0; security.len()];
        for i in 0..security.len() {
            if i == 0 {
                down[i] = 0;
                up[security.len() - 1 - i] = 0;
            } else {
                if security[i] <= security[i - 1] {
                    down[i] = down[i - 1] + 1;
                }
                if security[security.len() - 1 - i] <= security[security.len() - i] {
                    up[security.len() - 1 - i] = up[security.len() - i] + 1;
                }
            }
        }
        let mut res = vec![];
        for i in 0..up.len() {
            if down[i] >= time && up[i] >= time {
                res.push(i as i32);
            }
        }
        res
    }
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