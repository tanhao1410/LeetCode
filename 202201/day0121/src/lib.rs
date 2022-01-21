#[cfg(test)]
mod tests {
    use crate::can_jump;

    #[test]
    fn it_works() {
        for i in 2..1 {}
        assert_eq!(can_jump(vec![2, 3, 1, 1, 4]), true);
    }
}

//45. 跳跃游戏 II
pub fn jump(nums: Vec<i32>) -> i32 {
    //思路1：用一个dp[i]记录，从i位置到达结尾处最小的步骤，dp[nums.len() - 1] = 0
    let mut dp = vec![0; nums.len()];
    for i in (0..dp.len() - 1).rev() {
        let min = (1..=nums[i])
            .take_while(|&n| n as usize + i < nums.len())//不能超过数组的最大下标
            .map(|n| 1 + dp[n as usize + i])//该位置走n步到达下一个下标，到达结尾处需要的最重步数
            .min();//最小步数
        dp[i] = min.unwrap_or(10000);
    }
    dp[0]
}

//55. 跳跃游戏
pub fn can_jump(nums: Vec<i32>) -> bool {
    //用一个dp_num 记录能到达末尾的最小下标，开始时是len - 1,依次往前走，开能否走到0
    // let mut dp = nums.len() - 1;
    // for i in (0..dp).rev() {
    //     if nums[i] + i >= dp as i32 {
    //         dp = i;
    //     }
    // }
    // dp == 0
    nums
        .iter()
        .enumerate()
        .rev()
        .fold(nums.len() - 1, |dp, (i, &num)|
            match i + num as usize >= dp {
                true => i,
                false => dp
            })
        == 0
}