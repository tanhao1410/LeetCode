fn main() {
    println!("Hello, world!");
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