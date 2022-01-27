fn main() {
    println!("Hello, world!");
}

//91. 解码方法
pub fn num_decodings(mut s: String) -> i32 {
    let mut dp = vec![1; s.len()];
    //1234567987654322
    //以s[i]结尾的组合个数，不存在0了。dp[i] = s[i] 单独存在，
    //s[i]与前面结合。所以，dp[i] = dp[i-1]+dp[i-2]
    //单独求dp[1]
    let bytes = s.as_bytes();
    //以0开头的返回0
    if bytes[0] == b'0' {
        return 0;
    }
    //长度小于2返回1
    if bytes.len() < 2 {
        return 1;
    }
    for i in 1..bytes.len() {
        //看前面的数
        let pre = bytes[i - 1] - b'0';
        let cur = bytes[i] - b'0';
        match (pre, cur) {
            (1..=2, 0) => dp[i] = if i >= 2 { dp[i - 2] } else { 1 },
            (_, 0) => return 0,
            (2, cur) if cur < 7 => dp[i] = dp[i - 1] + if i > 1 { dp[i - 2] } else { 1 },
            (1, _) => dp[i] = dp[i - 1] + if i > 1 { dp[i - 2] } else { 1 },
            _ => dp[i] = dp[i - 1]
        }
    }
    dp[dp.len() - 1]
}

//695. 岛屿的最大面积
pub fn max_area_of_island(mut grid: Vec<Vec<i32>>) -> i32 {
    //思路：遍历，遇到1，遍历能碰到的1，得到面积，同时，将1置0
    let row = grid.len();
    let col = grid[0].len();
    let mut stack = vec![];

    let mut res = 0;
    for x in 0..row {
        for y in 0..col {
            if grid[x][y] == 1 {
                let mut square = 0;
                //碰到陆地了
                //广度优先遍历
                stack.push((x, y));
                while let Some((x, y)) = stack.pop() {
                    square += grid[x][y];
                    grid[x][y] = 0;
                    //看他的上下左右
                    if y > 0 && grid[x][y - 1] == 1 {
                        stack.push((x, y - 1));
                    }
                    if y < col - 1 && grid[x][y + 1] == 1 {
                        stack.push((x, y + 1));
                    }
                    if x > 0 && grid[x - 1][y] == 1 {
                        stack.push((x - 1, y));
                    }
                    if x < row - 1 && grid[x + 1][y] == 1 {
                        stack.push((x + 1, y));
                    }
                }
                res = res.max(square);
            }
        }
    }
    res
}

//733. 图像渲染
pub fn flood_fill(mut image: Vec<Vec<i32>>, sr: i32, sc: i32, new_color: i32) -> Vec<Vec<i32>> {
    let old_color = image[sr as usize][sc as usize];
    if old_color == new_color {
        return image;
    }
    //深度遍历
    let mut stack = vec![(sr as usize, sc as usize)];
    while let Some((x, y)) = stack.pop() {
        //把该位置换成其他颜色
        image[x][y] = new_color;
        //上下左右加入
        if y > 0 && image[x][y - 1] == old_color {
            stack.push((x, y - 1));
        }
        if y < image[0].len() - 1 && image[x][y + 1] == old_color {
            stack.push((x, y + 1));
        }
        if x > 0 && image[x - 1][y] == old_color {
            stack.push((x - 1, y));
        }
        if x < image.len() - 1 && image[x + 1][y] == old_color {
            stack.push((x + 1, y));
        }
    }
    image
}

//2047. 句子中的有效单词数
pub fn count_valid_words(sentence: String) -> i32 {
    sentence
        .split(" ")
        .filter(|&s| s.len() > 0)
        .filter(|&s| {
            //不含数字，最多含一个-，且两边都有字母,最后一个可以为标点
            let bytes = s.as_bytes();
            //判断最后一个是否是标点
            let last = bytes[bytes.len() - 1];
            let end;
            if last == b',' || last == b'.' || last == b'!' {
                end = bytes.len() - 1;
            } else {
                end = bytes.len();
            }

            let mut substrct_count = 0;
            for i in 0..end {
                //如果不是字母，也不是-，返回false
                if bytes[i] == b'-' {
                    substrct_count += 1;
                    if substrct_count > 1 || i == 0 || i == end - 1 {
                        return false;
                    }
                } else if bytes[i] < b'a' || bytes[i] > b'z' {
                    return false;
                }
            }
            true
        })
        .count() as i32
}