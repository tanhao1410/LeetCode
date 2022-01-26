use std::collections::HashMap;

fn main() {
    println!("Hello, world!");

    let mut square = DetectSquares::new();
    square.add(vec![3,10]);
    square.add(vec![11,2]);
    square.add(vec![3,2]);
    println!("{}", square.count(vec![11, 10]));
    println!("{}", square.count(vec![14, 8]));
    square.add(vec![11,2]);
    println!("{}", square.count(vec![11, 10]));

}

//2013. 检测正方形
struct DetectSquares {
    x_map: HashMap<i32, HashMap<i32, i32>>,
    y_map: HashMap<i32, HashMap<i32, i32>>,
    // key  代表 x,y轴坐标。value 代表 在这条线上，有多少个点，value 中的key代表 另一个轴的坐标。value代表 数量
    map: HashMap<(i32, i32), i32>,
}

impl DetectSquares {
    fn new() -> Self {
        Self {
            x_map: HashMap::new(),
            y_map: HashMap::new(),
            map: HashMap::new(),
        }
    }

    fn add(&mut self, point: Vec<i32>) {
        let x = point[0];
        let y = point[1];

        //加入到x轴map中
        let x_map_value = self.x_map.entry(x).or_insert(HashMap::new());
        let x_map_value_count = x_map_value.entry(y).or_insert(0);
        *x_map_value_count += 1;


        //加入到y轴map中
        let y_map_value = self.y_map.entry(y).or_insert(HashMap::new());
        let y_map_value_count = y_map_value.entry(x).or_insert(0);
        *y_map_value_count += 1;

        let count = self.map.entry((x, y)).or_insert(0);
        *count += 1;
    }

    fn count(&self, point: Vec<i32>) -> i32 {
        let x = point[0];
        let y = point[1];
        //看对应x轴有多少个点，根据他们之间的距离，去对应的 y轴再去寻找，如果都找到了，
        //看斜对角是否存在这个点。
        let mut res = 0;
        let x_map = self.x_map.get(&x);
        let y_map = self.y_map.get(&y);
        if x_map.is_none() || y_map.is_none() {
            return 0;
        }
        let x_map = x_map.unwrap();
        let y_map = y_map.unwrap();
        //找x轴上点
        for (&k, &count) in y_map.iter() {
            //距离
            let dis = (x - k).abs();
            if dis == 0 {
                continue;
            }
            //找对应的y轴上的点。上、下
            let up_y_point = x_map.get(&(y + dis));
            if let Some(y_count) = up_y_point {
                //上的y轴的点也找到了，找斜对角
                res += count * y_count * self.map.get(&(k, y + dis)).unwrap_or(&0);
            }
            let down_y_point = x_map.get(&(y - dis));
            if let Some(y_count) = down_y_point {
                res += count * y_count * self.map.get(&(k, y - dis)).unwrap_or(&0);
            }
        }
        res
    }
}

//567. 字符串的排列
pub fn check_inclusion(s1: String, s2: String) -> bool {
    let mut s1_map = vec![0; 26];
    s1
        .as_bytes()
        .iter()
        .for_each(|&l| s1_map[(l - b'a') as usize] += 1);
    //寻找一个大小相等的窗口
    if s2.len() < s1.len() {
        return false;
    }
    let mut s2_map = vec![0; 26];
    s2
        .as_bytes()
        .iter()
        .take(s1.len())
        .for_each(|&l| s2_map[(l - b'a') as usize] += 1);
    let is_equal = |s2_map: &[i32]| {
        for i in 0..26 {
            if s2_map[i] != s1_map[i] {
                return false;
            }
        }
        true
    };
    //先判断
    if is_equal(&s2_map) {
        return true;
    }
    let s2_bytes = s2.as_bytes();
    for i in s1.len()..s2.len() {
        //加进来一个，减去一个
        let add = s2_bytes[i] - b'a';
        let remove = s2_bytes[i - s1.len()] - b'a';
        s2_map[add as usize] += 1;
        s2_map[remove as usize] -= 1;
        if is_equal(&s2_map) {
            return true;
        }
    }
    false
}

//3. 无重复字符的最长子串
pub fn length_of_longest_substring(s: String) -> i32 {
    //思路：dp[i]以s[i]结尾的最长无重复子串。
    //dp[i] = if map.get[s[i]] isnone i + 1 or i - map.get[s[i]] .min(dp[i - 1] + 1)
    let mut dp = vec![1; s.len()];
    let mut map = HashMap::new();
    let bytes = s.as_bytes();
    for i in 0..s.len() {
        let cur_letter = bytes[i];
        //得到它的前一个位置
        if let Some(l) = map.get(&cur_letter) {
            dp[i] = i - l;
        } else {
            dp[i] = i + 1;
        }
        if i > 0 {
            dp[i] = dp[i].min(dp[i - 1] + 1);
        }
        //更新该字母位置
        map.insert(cur_letter, i);
    }
    dp
        .into_iter()
        .max()
        .unwrap_or(0) as i32
}