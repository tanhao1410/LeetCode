fn main() {
    println!("Hello, world!");
}

impl Solution {
    //396. 旋转函数
    pub fn max_rotate_function(a: Vec<i32>) -> i32 {
        let sum = a.iter().sum::<i32>();
        let init = a
            .iter()
            .enumerate()
            .fold(0, |pre, (i, &v)| pre + i as i32 * v);
        a
            .iter()
            .enumerate()
            .skip(1)
            .fold((init, init), |(pre, res), (i, &v)| {
                let cur = pre + sum - a.len() as i32 * a[a.len() - i];
                (cur, res.max(cur))
            })
            .1
    }
    //209. 长度最小的子数组
    pub fn min_sub_array_len(target: i32, nums: Vec<i32>) -> i32 {
        //思路：求累计和，然后根据target，找他的下一个，时间复杂度logN n
        // 滑动窗口：如果窗口和大于target，则左边缩减。否则，右边增加。
        let mut sums = nums.clone();
        for i in 1..nums.len() {
            sums[i] += sums[i - 1];
        }
        if sums[nums.len() - 1] < target {
            return 0;
        }
        let mut res = nums.len() as i32;
        //求下一个累计和
        for i in 0..nums.len() {
            //该数前面的总和
            let cur_sum = sums[i] - nums[i];
            //目标值
            let new_target = cur_sum + target;
            let end = Self::target_sub(new_target, &sums[i..]);
            match end {
                None => return res,
                Some(end) => {
                    res = res.min(end + 1);
                }
            }
        }
        res
    }

    fn target_sub(target: i32, nums: &[i32]) -> Option<i32> {
        let mut l = 0;
        let mut r = nums.len() - 1;
        if nums[r] < target {
            return None;
        }
        let mut m = (l + r) / 2;
        while l < r {
            if nums[m] < target {
                l = m + 1;
            } else {
                r = m;
            }
            m = (l + r) / 2;
        }
        Some(m as i32)
    }
    //744. 寻找比目标字母大的最小字母
    pub fn next_greatest_letter(letters: Vec<char>, target: char) -> char {
        if *letters.last().unwrap() <= target {
            return letters[0];
        }
        let mut l = 0;
        let mut r = letters.len() - 1;
        let mut m = (l + r) / 2;
        while (l < r) {
            if letters[m] <= target {
                l = m + 1;
            } else {
                r = m;
            }
        }
        letters[m]
    }
}

struct Solution;