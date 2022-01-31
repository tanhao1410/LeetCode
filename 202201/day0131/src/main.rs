fn main() {
    println!("Hello, world!");
}

//304. 二维区域和检索 - 矩阵不可变
struct NumMatrix {
    dp: Vec<Vec<i32>>
}

impl NumMatrix {
    fn new(matrix: Vec<Vec<i32>>) -> Self {
        let mut dp = matrix.clone();
        dp[0][0] = matrix[0][0];
        for j in 1..matrix[0].len() {
            dp[0][j] = dp[0][j - 1] + matrix[0][j];
        }
        for i in 1..matrix.len() {
            dp[i][0] = dp[i - 1][0] + matrix[i][0];
        }
        //思路：记录matrix[i][j]为右下角的区域之和
        for i in 1..matrix.len() {
            for j in 1..matrix[0].len() {
                dp[i][j] = dp[i - 1][j] + dp[i][j - 1] - dp[i - 1][j - 1] + matrix[i][j];
            }
        }
        NumMatrix { dp }
    }

    fn sum_region(&self, row1: i32, col1: i32, row2: i32, col2: i32) -> i32 {
        let row1 = row1 as usize;
        let row2 = row2 as usize;
        let col1 = col1 as usize;
        let col2 = col2 as usize;
        if row1 == 0 && col1 == 0 {
            return self.dp[row2][col2];
        }
        if row1 == 0 {
            return self.dp[row2][col2] - self.dp[row2][col1 - 1];
        }
        if col1 == 0 {
            return self.dp[row2][col2] - self.dp[row1 - 1][col2];
        }
        self.dp[row2 as usize][col2 as usize] - self.dp[row2][col1 - 1] - self.dp[row1 - 1][col2] + self.dp[row1 - 1][col1 - 1]
    }
}


//784. 字母大小写全排列
pub fn letter_case_permutation(s: String) -> Vec<String> {
    //字母有小写，有大写，递归思路，每一个都有两个选择，大写或小写
    let mut res = vec![];
    let bytes = s.as_bytes();

    let mut res_bytes = vec![];
    //找到第一个为字母的元素
    for i in 0..s.len() {
        if bytes[i] >= b'a' && bytes[i] <= b'z' {
            let mut res_bytes2 = res_bytes.clone();
            res_bytes.push(bytes[i]);

            let mut outer_string = String::from_utf8(res_bytes).unwrap();

            res_bytes2.push(bytes[i] + b'A' - b'a');
            let mut outer_string2 = String::from_utf8(res_bytes2).unwrap();
            //需要知道剩下的结果
            let inner_res = letter_case_permutation(String::from_utf8_lossy(&bytes[i + 1..]).to_string());

            for s in &inner_res {
                let mut res_item = outer_string.clone();
                let mut res_item2 = outer_string2.clone();
                res_item.push_str(s);
                res_item2.push_str(s);
                res.push(res_item);
                res.push(res_item2);
            }
            return res;
        } else if bytes[i] >= b'A' && bytes[i] <= b'Z' {
            //有两种可能性
            let mut res_bytes2 = res_bytes.clone();
            res_bytes.push(bytes[i]);

            let mut outer_string = String::from_utf8(res_bytes).unwrap();

            res_bytes2.push(bytes[i] + b'a' - b'A');
            let mut outer_string2 = String::from_utf8(res_bytes2).unwrap();
            //需要知道剩下的结果
            let inner_res = letter_case_permutation(String::from_utf8_lossy(&bytes[i + 1..]).to_string());

            for s in &inner_res {
                let mut res_item = outer_string.clone();
                let mut res_item2 = outer_string2.clone();
                res_item.push_str(s);
                res_item2.push_str(s);
                res.push(res_item);
                res.push(res_item2);
            }
            return res;

            return res;
        } else {
            res_bytes.push(bytes[i]);
        }
    }

    vec![s]
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
