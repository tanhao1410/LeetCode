fn main() {
    println!("Hello, world!");
    println!("{:?}", search_range(vec![0, 1, 1, 1, 1], 1));
}

//34. 在排序数组中查找元素的第一个和最后一个位置
pub fn search_range(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut start = 0;
    let mut end = nums.len() - 1;
    let mut middle = (end - start) / 2 + start;
    while end > start {
        if nums[middle] == target {
            end = middle;
        } else if nums[middle] > target {
            if middle == 0 {
                break;
            }
            end = middle - 1;
        } else {
            start = middle + 1;
        }
        middle = (end - start) / 2 + start;
    }
    //退出循环后，谁是最左边呢？
    if nums[middle] != target {
        return vec![-1, -1];
    }
    let mut res = vec![middle as i32];
    start = 0;
    end = nums.len() - 1;
    middle = (end - start) / 2 + start;
    while end > start {
        if nums[middle] == target {
            //看它后面是不是
            if nums[middle + 1] == target {
                start = middle + 1;
            } else {
                break;
            }
        } else if nums[middle] > target {
            if middle == 0 {
                break;
            }
            end = middle - 1;
        } else {
            start = middle + 1;
        }
        middle = (end - start) / 2 + start;
    }
    res.push(middle as i32);
    res
}