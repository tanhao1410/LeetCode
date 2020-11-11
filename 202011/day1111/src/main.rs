fn main() {
    println!("Hello, world!");
}

struct Solution{}

impl Solution{

    //221. 最大正方形  [["1"]]
    pub fn maximal_square(matrix: Vec<Vec<char>>) -> i32 {

        if matrix.len() == 0 || matrix[0].len()==0{
            return 0;
        }

        fn sqr(num :i32)->usize{
            for i in 1..num{
                if i * i == num{
                    return i as usize;
                }
            }
            1
        }

        let mut res = 0;

        let mut dp = Vec::new();
        for i in matrix {
            let mut row = vec![];
            for j in i {
                if j == '0' {
                    row.push(0);
                } else {
                    row.push(1);
                }
            }
            dp.push(row);
        }

        //思路：先从第一行开始和第一列开始，0 即0,1即1
        for i in 0..dp.len() {
            for j in 0..dp[0].len() {
                if i > 0 && j > 0 && dp[i][j] != 0 && dp[i - 1][j - 1] != 0 {

                    //看它的左上角
                    let pre_width = sqr(dp[i - 1][j - 1]);
                    //往左走，和往上走，
                    let mut temp = 0;
                    while temp < pre_width && dp[i][j - temp - 1] != 0 {
                        temp += 1;
                    }
                    let mut temp2 = 0;
                    while temp2 < pre_width && dp[i - temp2 - 1][j] != 0 {
                        temp2 += 1;
                    }
                    if temp < temp2{
                        dp[i][j] = ((1 + temp) * (1 + temp)) as i32;
                    }else{
                        dp[i][j] = ((1 + temp2) * (1 + temp2)) as i32;
                    }
                }
                if dp[i][j] > res{
                    res = dp[i][j];
                }
            }
        }
        res as i32
    }

    //165. 比较版本号
    pub fn compare_version(version1: String, version2: String) -> i32 {
        let version1 = version1.as_bytes();
        let version2 = version2.as_bytes();

        let zero: u8 = '0' as u8;
        let point: u8 = '.' as u8;

        fn compare(v1: &[u8], v2: &[u8]) -> i32 {
            if v1.len() == v2.len() {
                let mut i = 0;
                while i < v1.len() && v1[i] == v2[i] {
                    i += 1;
                }
                if i == v1.len() || v1[i] == v2[i] {
                    0
                } else if v1[i] > v2[i] {
                    1
                } else {
                    -1
                }
            } else if v1.len() > v2.len() {
                1
            } else {
                -1
            }
        }
        let (mut i, mut j) = (0, 0);
        let (mut end1, mut end2) = (0, 0);
        while i < version1.len() || j < version2.len() {

            //去除前导0
            while i < version1.len() && version1[i] == zero {
                i += 1;
            }
            while j < version2.len() && version2[j] == zero {
                j += 1;
            }

            //向后截取到小数点
            end1 = i;
            end2 = j;
            while end1 < version1.len() && version1[end1] != point {
                end1 += 1;
            }
            while end2 < version2.len() && version2[end2] != point {
                end2 += 1;
            }

            //某个版本号可能走到了最后面了
            if i >= version1.len() {
                i = version1.len();
                end1 = i;
            }
            if j >= version2.len() {
                j = version2.len();
                end2 = j;
            }

            //比较版本当前级别的版本号
            if compare(&version1[i..end1], &version2[j..end2]) != 0 {
                return compare(&version1[i..end1], &version2[j..end2]);
            }
            //开始下一轮
            i = end1 + 1;
            j = end2 + 1;
        }
        0
    }

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