fn main() {
    println!("Hello, world!");
}

//1447. 最简分数
pub fn simplified_fractions(n: i32) -> Vec<String> {
    let mut res = vec![];

    //判断两个数是否存在出1之外的公约数
    let exist_common = |a: i32, b: i32| {
        for i in 2..=a.min(b) {
            if a % i == 0 && b % i == 0 {
                return true;
            }
        }
        false
    };
    // a/b
    let create_str = |a: i32, b: i32| {
        let mut item = String::new();
        item.push_str(&a.to_string());
        item.push_str("/");
        item.push_str(&b.to_string());
        item
    };

    //从 1/i开始
    for i in 1..n {
        for j in i + 1..=n {
            if !exist_common(i, j) {
                res.push(create_str(i, j));
            }
        }
    }

    res
}