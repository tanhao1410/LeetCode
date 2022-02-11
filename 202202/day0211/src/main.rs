fn main() {
    println!("Hello, world!");
}



//1984. 学生分数的最小差值
pub fn minimum_difference(nums: Vec<i32>, k: i32) -> i32 {
    let mut nums = nums;
    nums.sort_unstable();
    let mut res = i32::MAX;
    for i in 0..nums.len() - k as usize + 1 as usize {
        res = res.min(nums[i + k as usize - 1 as usize] - nums[i]);
    }
    res
}