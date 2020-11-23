fn main() {
    println!("Hello, world!");
}

//119. 杨辉三角 II
pub fn get_row(row_index: i32) -> Vec<i32> {
    let mut res = vec![1; row_index as usize];
    let mut temp = vec![1; row_index as usize];
    for i in 1..row_index {
        for j in 0..i {
            if j > 0 {
                res[j as usize] += temp[j as usize - 1];
            }
        }
        for j in 0..i as usize {
            temp[j] = res[j]
        }
    }
    res
}