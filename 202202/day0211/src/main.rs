fn main() {
    println!("Hello, world!");
}

//11. 盛最多水的容器
pub fn max_area(height: Vec<i32>) -> i32 {
    //思路：剪枝。什么样的数字可以删除。当后面有比自己还大或等于的，自己不可能作为结束
    //当前面有比自己大或相等的，自己不可能作为开始。
    //不能作为开始和结束的，直接删除。
    let mut starts = vec![0];
    let mut max = height[0];
    for i in 1..height.len() - 1 {
        if height[i] > max {
            starts.push(i);
            max = height[i];
        }
    }
    let mut ends = vec![height.len() - 1];
    let mut max = height[height.len() - 1];
    for i in (1..height.len() - 1) {
        if height[i] > max {
            ends.push(i);
            max = height[i];
        }
    }

    //从starts中找开始，从ends中找结束，且需要ends > starts
    let mut res = 0;
    for i in 0..starts.len() {
        for j in 0..ends.len() {
            if j > i {
                res = res.max((j - i) as i32 * height[i].min(height[j]));
            }
        }
    }
    res
}

//986. 区间列表的交集
pub fn interval_intersection(first_list: Vec<Vec<i32>>, second_list: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
    //怎么样才算有交集？第一，下面的区间的开始应大于等于上面的开始。或下面的区间开始处大于上面的，且结束处大于
    let mut i = 0;
    let mut j = 0;
    let mut res = vec![];
    while i < first_list.len() && j < second_list.len() {
        //看上下两个区间是否有交集
        //第一个区间的结束小于第二个区间的开始
        if first_list[i][1] < second_list[j][0] {
            i += 1;
        } else if second_list[j][1] < first_list[i][0] {
            //下面区间的结束小于上面区间的开始
            j += 1;
        } else {
            //有交集
            //开始区间取大的，结束区间用小的
            let start = first_list[i][0].max(second_list[j][0]);
            let end = first_list[i][1].min(second_list[j][1]);
            //谁的区间结束了谁往前走
            if first_list[i][1] == end {
                i += 1;
            }
            if second_list[j][i] == end {
                j += 1;
            }
            res.push(vec![start, end]);
        }
    }
    res
}

//844. 比较含退格的字符串
pub fn backspace_compare(s: String, t: String) -> bool {
    let process_str = |s: String| -> String{
        let bytes = s.as_bytes();
        let mut res_bytes = vec![];
        for i in 0..bytes.len() {
            if bytes[i] == b'#' && res_bytes.len() > 0 {
                if res_bytes.len() > 0 {
                    res_bytes.remove(res_bytes.len() - 1);
                }
            } else {
                res_bytes.push(bytes[i]);
            }
        }
        String::from_utf8(res_bytes).unwrap()
    };
    process_str(s) == process_str(t)
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