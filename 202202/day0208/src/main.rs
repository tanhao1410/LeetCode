fn main() {
    println!("Hello, world!");
    println!("{:?}", search_range(vec![0, 1, 1, 1, 1], 1));
    println!("{}", search(vec![4, 5, 6, 7, 0, 1, 2], 0));
}

//74. 搜索二维矩阵
pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
    //思路：从右上角开始，遇到相等的返回true，遇到比target大的，往左走，遇到比target小的，往下走，走不动了，返回false
    let (mut x, mut y) = (0, matrix[0].len() - 1);
    loop {
        if matrix[x][y] == target {
            return true;
        } else if matrix[x][y] > target {
            //往左边走
            if y == 0 {
                return false;
            }
            y -= 1;
        } else {
            //往下走
            if x == matrix.len() - 1 {
                return false;
            }
            x += 1;
        }
    }
}

//33. 搜索旋转排序数组
pub fn search(nums: Vec<i32>, target: i32) -> i32 {
    // 如果最后一个元素大于第一个元素，说明是有序的。
    //如果不大于，说明旋转了。找中间位置。如果中间的数大于 第一个数，说明，旋转位置还在后面。如果 小于第一个数，说明旋转的数在前面
    if nums.len() == 1 {
        if nums[0] == target {
            return 0;
        } else {
            return -1;
        }
    }
    let binary_search = |nums: &[i32], target: i32| -> i32{
        if nums.len() == 0 {
            return -1;
        }
        let mut start = 0;
        let mut end = nums.len() - 1;
        let mut middle = (end - start) / 2 + start;
        while end >= start {
            if nums[middle] == target {
                return middle as i32;
            } else if nums[middle] < target {
                start = middle + 1;
            } else {
                if middle == 0 {
                    break;
                }
                end = middle - 1;
            }
            middle = (end + start) / 2;
        }
        -1
    };

    //有序的
    if nums[nums.len() - 1] > nums[0] {
        return binary_search(&nums, target);
    }

    let mut start = 0;
    let mut end = nums.len() - 1;
    let mut middle = (end - start) / 2 + start;
    // 如果target 处于 两者之间，则只需要在区间进行查找即可。
    while start <= end {
        if nums[middle] >= nums[0] {
            start = middle + 1;
        } else {
            if middle == 0 {
                break;
            }
            end = middle - 1;
        }
        middle = (start + end) / 2;
    }
    //start指向的位置就是旋转的点
    //println!("{}-{}",start,end);

    if target > nums[nums.len() - 1] {
        return binary_search(&nums[..start], target);
    } else {
        let res = binary_search(&nums[start..], target);
        if res == -1 {
            return -1;
        }
        return res + start as i32;
    }
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