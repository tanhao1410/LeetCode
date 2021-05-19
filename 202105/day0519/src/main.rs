fn main() {
    println!("Hello, world!");
}

//1738. 找出第 K 大的异或坐标值
pub fn kth_largest_value(matrix: Vec<Vec<i32>>, k: i32) -> i32 {
    //思路：v[m][n] = v[m-1][n] ^ v[m][n-1]^ v[m-1][n-1] ^ matrix[m][n]
    let mut v: Vec<Vec<i32>> = matrix;
    let mut nums = vec![];
    //一行一行求
    for i in 0..v.len() {
        for j in 0..v[0].len() {
            if i >= 1 {
                v[i][j] ^= v[i - 1][j]
            }
            if j >= 1 {
                v[i][j] ^= v[i][j - 1]
            }
            if i >= 1 && j >= 1 {
                v[i][j] ^= v[i - 1][j - 1]
            }
            nums.push(v[i][j]);
        }
    }


    nums.sort();
    nums[nums.len() - k as usize]
}