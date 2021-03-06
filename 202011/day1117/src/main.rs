fn main() {
    println!("Hello, world!");
    println!("{:?}", Solution::all_cells_dist_order(3, 2, 0, 0));
}

impl Solution {

    //剑指 Offer 31. 栈的压入、弹出序列
    pub fn validate_stack_sequences(pushed: Vec<i32>, popped: Vec<i32>) -> bool {
        let mut stack = vec![];
        let (mut i ,mut j) = (0,0);
        while i < popped.len() && (j < pushed.len() || !stack.is_empty()){
            if stack.is_empty(){
                stack.push(pushed[j]);
                j += 1;
            }
            while *stack.last().unwrap() != popped[i] && j < pushed.len(){
                stack.push(pushed[j]);
                j += 1;
            }
            if *stack.last().unwrap() == popped[i]{
                stack.pop();
                i += 1;
            }else{
                return false;
            }
        }
        i == popped.len()
    }

    //每日一题：1030. 距离顺序排列矩阵单元格
    pub fn all_cells_dist_order(r: i32, c: i32, r0: i32, c0: i32) -> Vec<Vec<i32>> {
        let mut res = vec![vec![r0, c0]];
        //距离最近的，第一个肯定是自己本身，距离为0.然后距离为1的，距离为2的，。。。。直到所有的都加入进来即可。
        let mut dis = 1;
        while (r * c) as usize > res.len() {
            let mut dis_r = 0;
            while dis_r <= dis {
                let dis_c = dis - dis_r;
                if r0 + dis_r < r {
                    if c0 + dis_c < c {
                        res.push(vec![r0 + dis_r, c0 + dis_c]);
                    }
                    if c0 - dis_c >= 0 && c0 + dis_c != c0 - dis_c {
                        res.push(vec![r0 + dis_r, c0 - dis_c]);
                    }
                }
                if r0 - dis_r >= 0 && r0 - dis_r != r0 + dis_r {
                    if c0 + dis_c < c {
                        res.push(vec![r0 - dis_r, c0 + dis_c]);
                    }
                    if c0 - dis_c >= 0 && c0 + dis_c != c0 - dis_c {
                        res.push(vec![r0 - dis_r, c0 - dis_c]);
                    }
                }
                dis_r += 1;
            }
            dis += 1;
        }
        res
    }
}

struct Solution {}