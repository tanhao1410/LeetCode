fn main() {
    println!("Hello, world!");
}

//153. 寻找旋转排序数组中的最小值
pub fn find_min(nums: Vec<i32>) -> i32 {
    //寻找旋转位置
    if nums.len() == 1 || nums[nums.len() - 1] > nums[0] {
        return nums[0];
    }

    let mut start = 0;
    let mut end = nums.len() - 1;
    let mut mid = (end + start) / 2;
    while end > start {
        //如果中间的数比nums[0]大，说明，旋转位置还在后面，如果比它小，说明旋转位置是它或它前面
        if nums[mid] >= nums[0] {
            start = mid + 1;
        } else {
            end = mid;
        }
        mid = (end + start) / 2;
    }
    nums[mid]
}

//2006. 差的绝对值为 K 的数对数目
pub fn count_k_difference(nums: Vec<i32>, k: i32) -> i32 {
    use std::collections::HashMap;
    let mut map = HashMap::new();
    for num in nums {
        let entry = map.entry(num).or_insert(0);
        *entry += 1;
    }
    let mut res = 0;
    for (kay, v) in &map {
        res += *map.get(&(*kay + k)).unwrap_or(&0) * *v;
        res += *map.get(&(*kay - k)).unwrap_or(&0) * *v;
    }
    res / 2
}