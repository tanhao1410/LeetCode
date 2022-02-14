fn main() {
    println!("Hello, world!");
}

//540. 有序数组中的单一元素
pub fn single_non_duplicate(nums: Vec<i32>) -> i32 {
    //思路，如果就一个元素，直接返回，否则，肯定是奇数个。从中间开始，看该数是否是单个，如果是单个，返回。
    //否则，如果和前面相同，
    // 用递归的方式，更简单
    if nums.len() == 1 {
        return nums[0];
    }
    let mut start = 0;
    let mut end = nums.len() - 1;
    let mut mid = (end + start) / 2;
    while start < end {
        //看中间的和前面相等还是后面相等
        if nums[mid] == nums[mid - 1] {
            if mid % 2 == 0 {
                //在前面
                end = mid - 2;
            } else {
                //在后面
                start = mid + 1;
            }
        } else if nums[mid] == nums[mid + 1] {
            if mid % 2 == 0 {
                //在后面
                start = mid + 2;
            } else {
                end = mid - 1;
            }
        } else {
            return nums[mid];
        }
        mid = (end + start) / 2;
    }
    nums[mid]
}