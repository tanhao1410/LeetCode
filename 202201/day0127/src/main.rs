fn main() {
    println!("Hello, world!");
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