fn main() {
    println!("Hello, world!");
}
struct  Solution;
impl Solution {
    //240. 搜索二维矩阵 II
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        let mut x = 0i32;
        let mut y = matrix[0].len() as i32 - 1;
        while x < matrix.len() as i32 && y >= 0{
            if matrix[x as usize][y as usize] == target{
                return true;
            }
            if matrix[x as usize][y as usize] > target{
                y -= 1;
            }else{
                x += 1;
            }
        }
        false
    }
}