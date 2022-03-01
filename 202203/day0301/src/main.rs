fn main() {
    println!("Hello, world!");
}

impl Solution {
    //6. Z 字形变换
    pub fn convert(s: String, num_rows: i32) -> String {
        if num_rows == 1 {
            return s;
        }
        let mut m = vec![vec![]; num_rows as usize];
        let mut x = 0;
        let mut y = 0;
        let mut is_down = true;//方向有往下和往右上
        let bytes = s.as_bytes();
        for &b in bytes {
            m[x].push(b);
            if is_down {
                if x < num_rows as usize - 1 {
                    x += 1;
                } else {
                    is_down = false;
                    x -= 1;
                    y += 1;
                }
            } else {
                if x == 0 {
                    is_down = true;
                    x += 1;
                } else {
                    x -= 1;
                    y += 1;
                }
            }
        }
        String::from_utf8(m
            .into_iter()
            .flat_map(|v| v.into_iter())
            .collect::<Vec<_>>()).unwrap()
    }
}

struct Solution;