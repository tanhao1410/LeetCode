impl Solution {
    ///2022. 将一维数组转变成二维数组
    pub fn construct2_d_array(original: Vec<i32>, m: i32, n: i32) -> Vec<Vec<i32>> {
        if m * n != original.len() as i32 {
            return vec![];
        }
        //original.chunks(n as usize).map(|v|)
        //original.chunks(n as usize).map(|x| x.to_vec()).collect()
        original.into_iter().fold(vec![], |mut res, v| {
            if res.last().is_none() || res.last().unwrap().len() == n as usize {
                res.push(vec![]);
            }
            res.last_mut().unwrap().push(v);
            res
        })
    }
}

struct Solution {}


fn main() {
    println!("Hello, world!");
}
