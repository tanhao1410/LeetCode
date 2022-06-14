fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //498. 对角线遍历
    pub fn find_diagonal_order(mat: Vec<Vec<i32>>) -> Vec<i32> {
        let mut res = vec![];
        let (m, n) = (mat.len(), mat[0].len());
        let mut cur = (0, 0);
        let mut is_up = true;
        while res.len() < m * n {
            res.push(mat[cur.0][cur.1]);
            if is_up {
                if cur.0 as i32 - 1 >= 0 && cur.1 + 1 < n {
                    cur.0 -= 1;
                    cur.1 += 1;
                } else if cur.1 + 1 < n {
                    is_up = !is_up;
                    cur.1 += 1;
                } else {
                    is_up = !is_up;
                    cur.0 += 1;
                }
            } else {
                if cur.0 + 1 < m && cur.1 as i32 - 1 >= 0 {
                    cur.0 += 1;
                    cur.1 -= 1;
                } else if cur.0 + 1 < m {
                    cur.0 += 1;
                    is_up = !is_up;
                } else {
                    is_up = !is_up;
                    cur.1 += 1;
                }
            }
        }
        res
    }
}