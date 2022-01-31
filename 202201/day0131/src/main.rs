fn main() {
    println!("Hello, world!");
}

//46. 全排列
pub fn permute(nums: Vec<i32>) -> Vec<Vec<i32>> {
    //采用递归的方式，先选择一个，放在第一个位置，然后，nums中元素少一个，递归实现
    if nums.len() == 1 {
        return vec![nums];
    }
    let mut res = vec![];
    for i in 0..nums.len() {
        //选择一个元素放在最后位置
        //剩下的元素
        let mut remain = vec![0; nums.len() - 1];
        for j in 0..i {
            remain[j] = nums[j];
        }
        for j in i + 1..nums.len() {
            remain[j - 1] = nums[j]
        }

        let mut inner_res = permute(remain);
        inner_res.iter_mut().for_each(|v| v.push(nums[i]));
        res.append(&mut inner_res);
    }

    res
}

//77. 组合
pub fn combine(n: i32, k: i32) -> Vec<Vec<i32>> {
    //思路：递归，每次选择一个最大的值，n, 那么剩下的只能在n-1中选k - 1个，
    if k == 1 {
        let mut res = vec![];
        for i in 1..=n {
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
