fn main() {
    println!("Hello, world!");
}

//118. 杨辉三角
pub fn generate(num_rows: i32) -> Vec<Vec<i32>> {
    let mut res = vec![vec![1]];
    for _ in 1..num_rows as usize {
        let cur_line = &res[res.len() - 1];
        let mut next_line = vec![1; cur_line.len() + 1];
        for i in 1..next_line.len() - 1 {
            next_line[i] = cur_line[i - 1] + cur_line[i];
        }
        res.push(next_line);
    }
    res
}