fn main() {
    println!("Hello, world!");
}

struct Solution{}

impl Solution{

    //240. 搜索二维矩阵 II
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        if matrix.len() == 0 || matrix[0].len() == 0 {
            return false;
        }
        let (row, col) = (matrix.len(), matrix[0].len());
        //思路：先根据二分法，在最后一列中查询，得到该数可能存在的行
        //在余下来的行中进行查找
        fn binary_search(row: &Vec<i32>, target: i32) -> bool {
            let (mut start, mut end) = (0, row.len() - 1);
            let mut middle = (start + end) / 2;
            while start as i32 <= end as i32 {
                if row[middle] == target {
                    return true;
                } else if row[middle] > target {
                    end = middle - 1;
                } else {
                    start = middle + 1;
                }
                middle = (end + start) / 2;
            }
            false
        }

        //找到可能的开始行
        let (mut start, mut end) = (0, row - 1);
        let mut middle = (start + end) / 2;
        while start as i32 <= end as i32 {
            if matrix[middle][col - 1] > target {
                end = middle - 1;
            } else if matrix[middle][col - 1] < target {
                start = middle + 1;
            } else {
                return true;
            }
            middle = (start + end) / 2;
        }

        //找到可能的结束行
        let (mut start1, mut end1) = (0, row - 1);
        middle = (start1 + end1) / 2;
        while start1 as i32 <= end1 as i32 {
            if matrix[middle][0] < target {
                start1 = middle + 1;
            } else if matrix[middle][0] > target {
                end1 = middle - 1;
            } else {
                return true;
            }
            middle = (start1 + end1) / 2;
        }

        for i in start..start1 {
            if binary_search(&matrix[i], target) {
                return true;
            }
        }
        false
    }
}