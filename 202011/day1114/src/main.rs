fn main() {
    println!("Hello, world!");
}

struct Solution {}

impl Solution {
    //每日一题：1122. 数组的相对排序
    pub fn relative_sort_array(arr1: Vec<i32>, arr2: Vec<i32>) -> Vec<i32> {
        let mut arr1 = arr1;
        //用一个1000的数组记录arr2中数字的位置信息
        let mut res = vec![1001; 1001];
        for i in 0..arr2.len() {
            res[arr2[i] as usize] = i;
        }
        //标准的插入排序方式
        for i in 0..arr1.len() - 1 {
            for j in i + 1..arr1.len() {
                //判断该数是否在arr2中存在
                if res[arr1[j] as usize] < res[arr1[i] as usize] || (res[arr1[j] as usize] == res[arr1[i] as usize] && arr1[j] < arr1[i]) {
                    let temp = arr1[i];
                    arr1[i] = arr1[j];
                    arr1[j] = temp;
                }
            }
        }
        arr1
    }

    //面试题 01.08. 零矩阵
    pub fn set_zeroes(matrix: &mut Vec<Vec<i32>>) {
        if matrix.len() == 0 || matrix[0].len() == 0 {
            return;
        }
        let m = matrix.len();
        let n = matrix[0].len();
        let mut row = vec![1; m];
        let mut col = vec![1; n];
        for i in 0..m {
            for j in 0..n {
                if matrix[i][j] == 0 {
                    row[i] = 0;
                    col[j] = 0;
                }
            }
        }

        for i in 0..m {
            if row[i] == 0 {
                for j in 0..n {
                    matrix[i][j] = 0;
                }
            }
        }

        for i in 0..n {
            if col[i] == 0 {
                for j in 0..m {
                    matrix[j][i] = 0;
                }
            }
        }
    }
}