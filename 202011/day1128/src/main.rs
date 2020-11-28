fn main() {
    println!("Hello, world!");
    unsafe {
        println!("{}", Solution::guessNumber(2126753390))
    }
}

struct Solution {}

impl Solution {
    //374. 猜数字大小
    unsafe fn guessNumber(n: i32) -> i32 {
        unsafe fn guessNum(min: i32, max: i32) -> i32 {
            //要考虑越界的问题；
            let middle = (max - min) / 2 + min;
            if guess(middle) == 0 {
                return middle;
            } else if guess(middle) < 0 {
                return guessNum(min, middle - 1);
            }
            return guessNum(middle + 1, max);
        }
        guessNum(1, n)
    }

    //每日一题：493. 翻转对
    pub fn reverse_pairs(nums: Vec<i32>) -> i32 {
        //暴力思想。时间超时无法完成。
        let mut count = 0;
        for i in 0..nums.len() - 1 {
            for j in i + 1..nums.len() {
                if nums[i] >> 1 > nums[j] || (nums[i] >> 1 == nums[j] && nums[i] % 2 == 1) {
                    count += 1
                }
            }
        }
        count
    }

    //389. 找不同
    pub fn find_the_difference(s: String, t: String) -> char {
        //使用数组来记录在s中字母出现了几次
        let mut letter_count = vec![0; 26];
        for i in s.as_bytes() {
            letter_count[(*i - 'a' as u8) as usize] += 1
        }
        for i in t.as_bytes() {
            letter_count[(*i - 'a' as u8) as usize] -= 1
        }

        let mut index = 0;
        while index < 26 && letter_count[index] == 0 {
            index += 1
        }

        ('a' as u8 + index as u8) as char
    }
}

unsafe fn guess(num: i32) -> i32 {
    return 1702766719 - num;
}

//304. 二维区域和检索 - 矩阵不可变
struct NumMatrix {
    matrix: Vec<Vec<i32>>
}

impl NumMatrix {
    fn new(matrix: Vec<Vec<i32>>) -> Self {
        NumMatrix {
            matrix: matrix
        }
    }

    fn sum_region(&self, mut row1: i32, mut col1: i32, row2: i32, col2: i32) -> i32 {
        //没有元素时，直接返回0
        if self.matrix.len() == 0 || self.matrix[0].len() == 0 || row2 < 0 || col2 < 0
            || row2 >= self.matrix.len() as i32 || col2 >= self.matrix[0].len() as i32 {
            return 0;
        }
        let mut res = 0;

        if row1 < 0 {
            row1 == 0
        }
        if col1 < 0 {
            col1 == 0
        }

        for i in row1..row2 + 1 {
            for j in col1..col2 + 1 {
                res += self.matrix[i as usize][j as usize]
            }
        }
        res
    }
}