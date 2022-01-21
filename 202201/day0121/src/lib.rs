#[cfg(test)]
mod tests {
    use crate::{can_jump, jump2, min_jumps};

    #[test]
    fn it_works() {

        assert_eq!(min_jumps(vec![7]), 0);
        assert_eq!(min_jumps(vec![6,1,9]), 2);
        assert_eq!(min_jumps(vec! [11,22,7,7,7,7,7,7,7,22,13]), 3);
        assert_eq!(min_jumps(vec!  [100,-23,-23,404,100,23,23,23,3,404]), 3);
    }
}

//1345. 跳跃游戏 IV
pub fn min_jumps(arr: Vec<i32>) -> i32 {
    //用一个map记录所有的数字 所在的下标，
    use std::collections::HashMap;
    let mut map = HashMap::new();
    for i in 0..arr.len() {
        let v = map.entry(arr[i]).or_insert(vec![]);
        v.push(i);
    }

    //按照一步一步往前走的走法
    let mut dp = vec![0; arr.len()];
    for i in 0..dp.len() {
        dp[i] = (dp.len() - 1 - i);
    }

    //更新相等的数，到最终节点的步数是相同的。
    for (_, v) in map.iter() {
        //得到里面距离终点步数最少的
        let min = v.iter().map(|&i| dp[i]).min().unwrap();
        //需要min + 1,因为需要多走一步，走到这个位置
        v.iter().for_each(|&i| dp[i] = dp[i].min(min + 1));
    }
    dp[arr.len() - 1] = 0;

    //开始遍历，如果一个数，可以往前走往后走，发现比自己本来的步数小了，则更新步数，它的更新，会顺带着与自己相同的都会进行更新，都没得更新的话，返回结果
    let mut flag = true;
    while flag {
        flag = false;
        for i in 0..arr.len() - 1 {
            // 更新更快捷的方式,可以看它前面或后面
            if dp[i] > dp[i + 1] + 1 || (i > 0 && dp[i - 1] + 1 < dp[i]) {
                flag = true;
                dp[i] = dp[i + 1] + 1;
                if i > 0{
                    dp[i] = dp[i].min(dp[i - 1] + 1);
                }
                //与它相同key的也更新一下
                let v = map.get(&arr[i]).unwrap();
                for &o in v.iter() {
                    if dp[o] > dp[i] + 1 {
                        dp[o] = dp[i] + 1;
                    }
                }
            }
        }
    }

    dp[0] as i32
}

//45. 跳跃游戏 II
pub fn jump2(nums: Vec<i32>) -> i32 {
    //参考解法：每一个位置能走到的最远距离
    let mut step = 0;
    let mut end = 0;
    let mut max_position = 0;
    for i in 0..nums.len() - 1 {
        max_position = max_position.max(nums[i] + i as i32);
        if i == end {
            end = max_position as usize;
            step += 1;
        }
    }
    step
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