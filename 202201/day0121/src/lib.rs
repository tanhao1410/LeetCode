#[cfg(test)]
mod tests {
    use crate::can_jump;

    #[test]
    fn it_works() {
        assert_eq!(can_jump(vec![2,3,1,1,4]), true);
    }
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