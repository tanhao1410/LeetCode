fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //2239. 找到最接近 0 的数字
    pub fn find_closest_number(nums: Vec<i32>) -> i32 {
        let mut distance = i32::MAX;
        let mut res = i32::MIN;
        for num in nums {
            if num.abs() < distance || (num.abs() == distance && num > res) {
                distance = num.abs();
                res = num;
            }
        }
        res
    }
}

use std::collections::HashMap;

extern crate rand;

//535. TinyURL 的加密与解密
struct Codec {
    urls: HashMap<String, String>,
    tiny_urls: HashMap<String, String>,
}

impl Codec {
    fn new() -> Self {
        Self {
            urls: HashMap::new(),
            tiny_urls: HashMap::new(),
        }
    }

    fn encode(&mut self, longURL: String) -> String {
        if let Some(tiny_url) = self.urls.get(longURL.as_str()) {
            return tiny_url.to_string();
        }
        let mut tiny_url = "http://t.cn/".to_string();
        for _ in 0..5 {
            tiny_url.push((rand::random::<u8>() % 48) as char);
        }
        if let Some(old) = self.tiny_urls.get(tiny_url.as_str()) {
            return self.encode(longURL);
        }

        self.urls.insert(longURL.clone(), tiny_url.clone());
        self.tiny_urls.insert(tiny_url.clone(), longURL.clone());
        tiny_url
    }

    fn decode(&self, shortURL: String) -> String {
        self.tiny_urls.get(shortURL.as_str()).unwrap().to_string()
    }
}