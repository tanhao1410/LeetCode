#[cfg(test)]
mod tests {
    use crate::{can_jump, jump2, min_jumps};

    #[test]
    fn it_works() {
        assert_eq!(min_jumps(vec![7]), 0);
        assert_eq!(min_jumps(vec![6, 1, 9]), 2);
        assert_eq!(min_jumps(vec![11, 22, 7, 7, 7, 7, 7, 7, 7, 22, 13]), 3);
        assert_eq!(min_jumps(vec![100, -23, -23, 404, 100, 23, 23, 23, 3, 404]), 3);

        assert_eq!(min_jumps(vec![100, -23, -23, 404, 100, 23, 23, 23, 3, 404]), 3);
        assert_eq!(min_jumps(vec![100, -23, -23, 404, 100, 23, 23, 23, 3, 404]), 3);
        assert_eq!(min_jumps(vec![100, -23, -23, 404, 100, 23, 23, 23, 3, 404]), 3);
    }
}

//771. 宝石与石头
pub fn num_jewels_in_stones(jewels: String, stones: String) -> i32 {
    stones
        .chars()
        .filter(|c|jewels.contains(*c))
        .count() as i32
}

//1424. 对角线遍历 II
pub fn find_diagonal_order2(nums: Vec<Vec<i32>>) -> Vec<i32> {
    //借鉴思路：同一斜线元素满足下标(i+j)相等
    let mut matrix = vec![];
    for row in (0..nums.len()).rev() {
        for col in 0..nums[row].len() {
            matrix.push((row + col, nums[row][col]));
        }
    }
    matrix.sort_by_key(|&e| e.0);
    matrix.iter().map(|e| e.1).collect()
}

//1424. 对角线遍历 II
pub fn find_diagonal_order(nums: Vec<Vec<i32>>) -> Vec<i32> {
    //用一个vec记录上面的没有完毕的row
    //访问的一趟数字，当还没达到底层的时候，一趟应该有 vec.len + 1个，多一个1是因为当前这一行的存在
    //应该加入的数字，从下往上依次是 nums[row][cur_row - row]
    let mut res = vec![];
    let mut not_comple = vec![];
    for i in 0..nums.len() {
        //先加入自己的
        not_comple.push(i);

        let mut need_rem = vec![];
        //在依次加入，注意顺序
        for not_comple_i in (0..not_comple.len()).rev() {
            res.push(nums[not_comple[not_comple_i]][i - not_comple[not_comple_i]]);
            if nums[not_comple[not_comple_i]].len() - 1 == i - not_comple[not_comple_i] {
                need_rem.push(not_comple_i);
            }
        }
        //如果一个row所有的数字已经加入完毕，则，删除它
        for x in need_rem {
            not_comple.remove(x);
        }
    }

    let mut col_num = 1;
    //到达尾部之后，
    while not_comple.len() > 0 {
        let mut need_rem = vec![];
        //在依次加入，注意顺序
        for not_comple_i in (0..not_comple.len()).rev() {
            res.push(nums[not_comple[not_comple_i]][nums.len() - 1 - not_comple[not_comple_i] + col_num]);
            if nums[not_comple[not_comple_i]].len() - 1 == nums.len() - 1 - not_comple[not_comple_i] + col_num {
                need_rem.push(not_comple_i);
            }
        }
        //如果一个row所有的数字已经加入完毕，则，删除它
        for x in need_rem {
            not_comple.remove(x);
        }
        col_num += 1;
    }

    res
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
                if i > 0 {
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