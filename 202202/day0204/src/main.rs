fn main() {
    println!("Hello, world!");
}

//1200. 最小绝对差
pub fn minimum_abs_difference(mut arr: Vec<i32>) -> Vec<Vec<i32>> {
    //思路：排序
    arr.sort_unstable();
    //先求最小绝对值差。
    let mut min = i32::MAX;
    for i in 1..arr.len() {
        min = min.min(arr[i] - arr[i - 1]);
    }
    let mut res = vec![];
    for i in 1..arr.len() {
        if arr[i] - arr[i - 1] == min {
            res.push(vec![arr[i - 1], arr[i]]);
        }
    }
    res
}

//1725. 可以形成最大正方形的矩形数目
pub fn count_good_rectangles(rectangles: Vec<Vec<i32>>) -> i32 {
    let map = rectangles
        .iter()
        .map(|v| v[0].min(v[1]));
    let max_len = map.max().unwrap();
    map.filter(|v| *v == max_len).count() as i32
}

//376. 摆动序列
pub fn wiggle_max_length(nums: Vec<i32>) -> i32 {
    //思路：两个dp，一个是nums[i]结尾的 增， 一个以nums[i]结尾的 减
    let mut dp_inc = nums.clone();
    let mut dp_dec = nums.clone();

    dp_dec[0] = 1;
    dp_inc[0] = 1;

    for i in 1..nums.len() {
        //对于递增的dp来说， nums[i] 如果大于 num[..i]中的一个，那么， dp_inc = dp_dec[i - x] + 1;
        let mut max_dec = 0;
        let mut max_inc = 0;
        for j in 0..i {
            if nums[j] < nums[i] {
                max_inc = max_inc.max(dp_dec[j]);
            } else if nums[j] > nums[i] {
                max_dec = max_dec.max(dp_inc[j]);
            }
        }
        dp_inc[i] = max_inc + 1;
        dp_dec[i] = max_dec + 1;
    }

    dp_dec.into_iter().max().unwrap().max(
        dp_inc.into_iter().max().unwrap()
    )
}

//300. 最长递增子序列
pub fn length_of_lis(nums: Vec<i32>) -> i32 {
    // o (n ^2)
    //以 nums[i]结尾的最长递增子序列
    let mut dp = nums.clone();
    dp[0] = 1;
    for i in 1..nums.len() {
        //找到它前面的 比nums[i]小的 最大 dp[i]
        let mut max = 0;
        for j in 0..i {
            if nums[j] < nums[i] {
                max = max.max(dp[j]);
            }
        }
        dp[i] = max + 1;
    }
    dp.into_iter().max().unwrap()
}