fn main() {
    println!("Hello, world!");
}

//1380. 矩阵中的幸运数
pub fn lucky_numbers(matrix: Vec<Vec<i32>>) -> Vec<i32> {
    //同一行中最小，同一列中最大
    let mut res = vec![];
    for i in 0..matrix.len() {
        //找到一行中最小的数
        let mut min = 0;
        for j in 1..matrix[0].len() {
            if matrix[i][j] < matrix[i][min] {
                min = j;
            }
        }
        //判断它是否是一列中最大的数
        let mut max = i;
        for j in 0..matrix.len() {
            if matrix[j][min] > matrix[i][min] {
                max = j;
            }
        }

        if max == i {
            res.push(matrix[i][min]);
        }
    }

    res
}
