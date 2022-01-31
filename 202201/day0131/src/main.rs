fn main() {
    println!("Hello, world!");
}

//77. 组合
pub fn combine(n: i32, k: i32) -> Vec<Vec<i32>> {
    //思路：递归，每次选择一个最大的值，n, 那么剩下的只能在n-1中选k - 1个，
    if k == 1{
        let mut res = vec![];
        for i in 1..=n{
            res.push(vec![i]);
        }
        return res;
    }
    //选择一个最大的数，选择一个次大的数
    let mut res = vec![];
    for i in (k..=n).rev() {
        let mut inner_res = combine(i - 1, k - 1);
        inner_res.iter_mut().for_each(|v| v.push(i));
        res.append(&mut inner_res);
    }
    res
}
