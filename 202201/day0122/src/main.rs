fn main() {
    println!("Hello, world!");
}

//189. 轮转数组
pub fn rotate(nums: &mut Vec<i32>, k: i32) {
    //思路:如果k大于nums.len，多于的无意义。因为，需要 k %= nums.len()
    let k = k as usize % nums.len();
    //使用中间的变量记录
    let mut temp = nums.clone();
    for i in 0..nums.len(){
         nums[(i + k) % temp.len()] = temp[i];
    }
}

//977. 有序数组的平方
pub fn sorted_squares(a: Vec<i32>) -> Vec<i32> {
    if a[0] >= 0 {
        return a.into_iter().map(|a| a * a).collect();
    } else if a[a.len() - 1] <= 0 {
        return a.into_iter().rev().map(|a| a * a).collect();
    }
    let mut res = vec![0;a.len()];
    //如果a里都是小于0或大于0，则简单。否则，从两边开始逼近，
    let mut i = 0;
    let mut j = a.len() - 1;
    let mut p = res.len() - 1;
    while j >= i{
        if a[j].abs() > a[i].abs(){
            res[p] = a[j] * a[j];
            j -= 1;
        }else{
            res[p] = a[i] * a[i];
            i += 1;
        }
        p -= 1;
    }
    res
}

//918. 环形子数组的最大和
pub fn max_subarray_sum_circular(nums: Vec<i32>) -> i32 {
    //用上次的手法，同时记录 元素的长度，使得总的长度不超过nums.len
    let mut max = i32::MIN;
    let mut pre = -1;
    let mut count = 0;
    for i in 0..nums.len() * 2 {
        if pre > 0 {
            if count < nums.len() {
                pre = pre + nums[i % nums.len()];
                count += 1;
            } else {
                //说明前面的已经包括了足够的数了
                //问题：？前面应该扣除几个？扣除掉最前面的1个
                //得到最大的值呢？
            }
        } else {
            pre = nums[i % nums.len()];
            count = 1;
        }
        max = max.max(pre);
    }
    max
}

//53. 最大子数组和
pub fn max_sub_array(nums: Vec<i32>) -> i32 {
    //dp[i]以nums[i]结尾的最大子数组和，if dp[i] > 0 则 dp[i-1] = dp[i] + nums[i]
    // nums
    //     .into_iter()
    //     .fold((i32::MIN,-1),|(res,pre),num|{
    //         let pre = 0.max(pre) + num;
    //         (res.max(pre),pre)
    //     }) // res,dp[i-1],
    //     .0
    let mut res = nums[0];
    let mut pre = nums[0];
    for i in 1..nums.len() {
        if pre > 0 {
            pre += nums[i];
        } else {
            pre = nums[i];
        }
        res = res.max(pre);
    }
    res
}

//1332. 删除回文子序列
pub fn remove_palindrome_sub(s: String) -> i32 {
    //思路：先删除最大的回文序列,dp[i]代表以i为中心最大的回文序列
    let mut dp = vec![1; s.len()];

    0
}