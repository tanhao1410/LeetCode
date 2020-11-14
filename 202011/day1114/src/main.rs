fn main() {
    println!("Hello, world!");
    Solution::smallest_k(vec![1, 3, 5, 7, 2, 4, 6, 8], 4);
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

    //剑指 Offer 50. 第一个只出现一次的字符
    pub fn first_uniq_char(s: String) -> char {
        //思路：用一个数组记录所有的字母出现的次数。
        let mut res = ' ';
        let mut chars = vec![0;27];
        for i in  s.chars(){
            //看i之前是否出现过
            let x = i as usize- 'a' as usize;
            chars[x] += 1;
        }
        for i in s.chars(){
            let x = i as usize- 'a' as usize;
            if chars[x] == 1{
                return i;
            }
        }
        res
    }

    //面试题 17.14. 最小K个数
    pub fn smallest_k(mut arr: Vec<i32>, k: i32) -> Vec<i32> {
        fn f(arr: &mut Vec<i32>, mut start: usize, mut end: usize, k: usize) {

            let first = arr[start];
            let start1 = start;
            start += 1;
            while end > start {
                while start < arr.len() && arr[start] <= first {
                    start += 1;
                }
                while arr[end] > first {
                    end -= 1;
                }
                if end > start {
                    let temp = arr[end];
                    arr[end] = arr[start];
                    arr[start] = temp;
                }
            }
            arr[start1] = arr[end];
            arr[end] = first;

            if end - start1 == k {
                return;
            } else if end - start1 > k {
                f(arr, 0, end - 1, k);
            } else {
                f(arr, end + 1, arr.len() - 1, k - end - 1);
            }
        }

        let max = arr.len() - 1;
        f(&mut arr, 0, max, k as usize);
        return arr[..k as usize].to_vec();
    }
}