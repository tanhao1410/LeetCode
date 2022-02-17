fn main() {
    println!("Hello, world!");
}

impl Solution {

    //40. 组合总和 II
    pub fn combination_sum2(mut candidates: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        //采用递归来做
        //怎么表示数字已经被使用了呢？用一个数组来记录
        candidates.sort_unstable();
        Self::combination_sum(&candidates,target,&mut vec![false;candidates.len()],0)
    }

    fn combination_sum(candidates:&[i32],target:i32,already:&mut Vec<bool>,pre:i32) ->Vec<Vec<i32>>{
        let mut res = vec![];
        if target == 0{
            let mut item = vec![];
            for i in 0..already.len(){
                if already[i]{
                    item.push(candidates[i]);
                }
            }
            res.push(item);
        }else{
            let mut pre_num = 0;
            //从剩下的元素中选择一个元素
            for i in 0..candidates.len(){
                //过滤掉已经选择过的元素
                if !already[i]{
                    if candidates[i] != pre_num && candidates[i] <= target && candidates[i] >= pre{
                        already[i] = true;
                        res.append(&mut Self::combination_sum(candidates,target - candidates[i],already,candidates[i]));
                        already[i] = false;
                    }
                    pre_num = candidates[i];
                }
            }
        }
        res
    }

    //1376 通知所有员工所需的时间
    pub fn num_of_minutes(n: i32, head_id: i32, manager: Vec<i32>, inform_time: Vec<i32>) -> i32 {
        //广度优先遍历，而应该用深度优先算法，计算最长路径
        // 先把manager数组变成一个vec[vec] v为该员工的下属
        let mut subs = vec![vec![]; n as usize];
        for i in 0..n as usize {
            //找到它的上级
            if manager[i] != -1 {
                subs[manager[i] as usize].push(i);
            }
        }
        //需要记录通知到某个员工的时间
        let mut reach_time = vec![0; n as usize];
        reach_time[head_id as usize] = inform_time[head_id as usize];
        let mut stack = vec![];
        stack.push(head_id as usize);
        //需要多少时间来通知它的下属
        while let Some(n) = stack.pop() {
            //如果没有下属了，代表这条线路走到了终点。
            //它的下属们进栈
            for sub in &subs[n] {
                stack.push(*sub);
                //上一级的时间 + 本机的通知时间
                reach_time[*sub] = reach_time[n] + inform_time[*sub];
            }
        }
        reach_time
            .into_iter()
            .max()
            .unwrap()
    }

    //47. 全排列 II
    pub fn permute_unique(mut nums: Vec<i32>) -> Vec<Vec<i32>> {
        nums.sort_unstable();
        //选择一个数，然后选择第二个数，如果接下来的数和自己相同，跳过该数。
        Self::select_next_num(nums, vec![])
    }

    fn select_next_num(nums: Vec<i32>, pre: Vec<i32>) -> Vec<Vec<i32>> {
        let mut res = vec![];
        if nums.len() == 0 {
            res.push(pre);
            return res;
        }
        let mut pre_num = 11;
        for i in 0..nums.len() {
            if nums[i] != pre_num {
                //选择当前数字
                let mut new_pre = pre.clone();
                new_pre.push(nums[i]);
                let mut new_nums = vec![0; nums.len() - 1];
                for j in 0..i {
                    new_nums[j] = nums[j];
                }
                for j in i + 1..nums.len() {
                    new_nums[j - 1] = nums[j];
                }
                res.append(&mut Self::select_next_num(new_nums, new_pre));
                pre_num = nums[i];
            }
        }
        res
    }
}