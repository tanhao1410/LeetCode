fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //885. 螺旋矩阵 III
    pub fn spiral_matrix_iii(rows: i32, cols: i32, mut r_start: i32, mut c_start: i32) -> Vec<Vec<i32>> {
        let mut res = vec![vec![r_start, c_start]];
        let push = |res: &mut Vec<Vec<i32>>, r: i32, c: i32| {
            if r >= 0 && r < rows && c >= 0 && c < cols {
                res.push(vec![r, c]);
            }
        };
        let mut step = 1;

        while res.len() < (rows * cols) as usize {
            for _ in 0..step {
                c_start += 1;
                push(&mut res, r_start, c_start);
            }
            for _ in 0..step {
                r_start += 1;
                push(&mut res, r_start, c_start);
            }
            step += 1;
            for _ in 0..step {
                c_start -= 1;
                push(&mut res, r_start, c_start);
            }
            for _ in 0..step {
                r_start -= 1;
                push(&mut res, r_start, c_start);
            }
            step += 1;
        }

        res
    }

    //829. 连续整数求和
    pub fn consecutive_numbers_sum(n: i32) -> i32 {
        (1..)
            .take_while(|&i| i * (i - 1) / 2 < n)
            .filter(|&i| (n - ((i + 1) % 2 * (i - 1) * i / 2)) % i == 0 && (n - ((i + 1) % 2) * (i - 1) * i / 2) / i > 0)
            .count() as i32
    }
}