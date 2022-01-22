fn main() {
    println!("Hello, world!");
}

//1217. 玩筹码
pub fn min_cost_to_move_chips(position: Vec<i32>) -> i32 {
    //思路；先无损移动，2之后的数据，都可以移动到位置1或2上
    position
        .iter()
        .fold((0, 0, 0), |(p, q, _), i|
            match *i % 2 {
                0 => (p + 1, q, q.min(p + 1)),
                _ => (p, q + 1, p.min(q + 1))
            })
        .2
}

//189. 轮转数组
pub fn rotate(nums: &mut Vec<i32>, k: i32) {
    //思路:如果k大于nums.len，多于的无意义。因为，需要 k %= nums.len()
    let k = k as usize % nums.len();
    //使用中间的变量记录
    let mut temp = nums.clone();
    for i in 0..nums.len() {
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
    let mut res = vec![0; a.len()];
    //如果a里都是小于0或大于0，则简单。否则，从两边开始逼近，
    let mut i = 0;
    let mut j = a.len() - 1;
    let mut p = res.len() - 1;
    while j >= i {
        if a[j].abs() > a[i].abs() {
            res[p] = a[j] * a[j];
            j -= 1;
        } else {
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
    let mut dp = nums.clone();
    let mut res = i32::MIN;
    for i in 1..nums.len() {
        dp[i] = nums[i] + 0.max(dp[i - 1]);
        res = res.max(dp[i]);
    }

    let mut dp2_cur = nums.clone();
    for i in 1..nums.len(){
        dp2_cur[i] = dp2_cur[i - 1] + nums[i];
    }

    //用一个dp2 记录，以nums[i..]的最大值，以及当前值
    let mut dp_max = nums.clone();
    let mut dp_cur = nums.clone();
    for i in (0..nums.len() - 1).rev() {
        dp_cur[i] = dp_cur[i + 1] + nums[i];
        dp_max[i] = dp_cur[i].max(dp_max[i + 1]);
    }

    //最大值
    for i in 0..nums.len() - 1 {
        //必须要用从0位开始的，不然，和后面的连不起来
        res = res.max(dp2_cur[i] + dp_max[i + 1]);
    }
    res
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
    //判断是否是回文，如果是，那么就删一次，如果不是，删两次即可。因为，就两种字母，a,b，先删a,再删b

    //判断是否是回文？简单办法：倒置后看是否相等
    //


    // let mut i = 0;
    // let mut j = s.len() - 1;
    // let mut chars = s.as_bytes();
    // while j > i{
    //     if chars[i] != chars[j]{
    //         return 2;
    //     }
    //     j -= 1;
    //     i += 1;
    // }
    // 1
    let mut s = s.as_bytes();
    while let [first, rest @ .., last] = s {
        if first != last {
            return 2;
        }
        s = rest;
    }
    1
}