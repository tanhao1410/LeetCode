fn main() {
    println!("Hello, world!");
}

//525. 连续数组
pub fn find_max_length(nums: Vec<i32>) -> i32 {
    //记录1与0数量差 最靠前的位置
    let mut m = std::collections::HashMap::new();
    //1的总和，包括自身
    let mut nums = nums;
    let mut res = 0;
    for i in 0..nums.len() {
        let pre = match i {
            0 => 0,
            _ => nums[i - 1]
        };

        if nums[i] == 1 {
            nums[i] = pre + 1
        } else {
            nums[i] = pre - 1
        }
        //从第一个开始算起
        if nums[i] == 0 {
            res = res.max(i + 1);
        } else {
            match m.get(&nums[i]) {
                Some(&preIndex) => {
                    res = res.max(i - preIndex);
                }
                _ => {
                    m.insert(nums[i], i);
                }
            }
        }
    }
    res as i32
}