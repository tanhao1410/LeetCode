fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //532. 数组中的 k-diff 数对
    pub fn find_pairs(nums: Vec<i32>, k: i32) -> i32 {
        let mut nums = nums;
        nums.sort_unstable();
        //只应该考虑数字大小，与数量没关系
        let mut uni_nums = vec![];
        let mut res = 0;
        let mut pre_num = i32::MAX;
        for i in nums {
            if !uni_nums.last().eq(&Some(&i)) {
                uni_nums.push(i);
            } else if pre_num != i {
                res += 1;
                pre_num = i;
            }
        }
        if k == 0 {
            return res;
        }
        res = 0;
        //找每一个数相对的，看是否在里面，双指针？
        let (mut left, mut right) = (0, 0);
        while right < uni_nums.len() && left <= right {
            //right 向右走，直到相差的距离大于等于k，
            while right < uni_nums.len() && uni_nums[right] - uni_nums[left] < k {
                right += 1;
            }
            if right == uni_nums.len() {
                break;
            }
            if uni_nums[right] - uni_nums[left] == k {
                res += 1;
                right += 1;
            }
            left += 1;
        }
        res
    }
}