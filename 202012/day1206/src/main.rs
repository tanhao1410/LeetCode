fn main() {
    println!("Hello, world!");
    println!("{:?}",Solution::generate(10))
}

struct Solution{}
impl Solution {
    //每日一题：118. 杨辉三角
    pub fn generate(num_rows: i32) -> Vec<Vec<i32>> {
        let mut res:Vec<Vec<i32>> = vec![];
        for i in 0..num_rows{
            //第i行
            if let Some(pre) = res.last(){
                let mut item = vec![1;pre.len()+1];
                for j in 1..pre.len(){
                    item[j] = pre[j - 1] + pre[j]
                }
                res.push(item);
            }else{
                res.push(vec![1]);
            }
        }
        res
    }
}