use std::ops::Add;

fn main() {
    println!("{}", next_greater_element(1999999999));
}

//1716. 计算力扣银行的钱
pub fn total_money(n: i32) -> i32 {
    //每天应拿的钱 n / 7 + n % 7
    (1..=n).map(|n| (n - 1) / 7 + match n % 7 {
        0 => 7,
        mod_ => mod_,
    }).sum()
}

//556. 下一个更大元素 III
pub fn next_greater_element(n: i32) -> i32 {
    let mut bits = vec![];
    let mut n = n;
    while n > 0 {
        bits.push(n % 10);
        n /= 10;
    }

    let mut i = 0;
    //看它前面的是否大于等于自己
    while i + 1 < bits.len() && bits[i + 1] >= bits[i] {
        i += 1;
    }
    if i + 1 == bits.len(){
        return -1;
    }
    println!("{:?}", bits);
    //否则，就找到了递减的地方了
    //找第一个比它大的地方，交换一下
    for big in 0..=i {
        if bits[big] > bits[i + 1] {
            //交换
            let temp = bits[i + 1];
            bits[i + 1] = bits[big];
            bits[big] = temp;
            break;
        }
    }

    let mut res = 0_i64;
    println!("{:?}", bits);

    //需要考虑溢出情况
    for j in 0..=i {
        //倒序累加
        res += 10_i64.pow((i - j) as u32) * bits[j] as i64;
    }

    for j in i + 1..bits.len() {
        res += 10_i64.pow(j as u32) * bits[j] as i64;
    }

    match res > i32::MAX as i64{
        true=>-1,
        _=>res as i32
    }
}

//556. 下一个更大元素 III
//题目理解错误，理解为 二进制中的每一位了
pub fn next_greater_element2(n: i32) -> i32 {
    let mut bits = vec![false; 32];
    for i in 0..32 {
        bits[i] = n & 1 << 31 - i > 0;
    }

    //先找到第一个1
    let mut first = 31;
    while !bits[first] {
        first -= 1;
    }

    // 0 100000000
    if first == 1 {
        return -1;
    }

    //从first 开始往前找第一个为0的位置
    let mut zero = first;
    while bits[zero] {
        //将所遇到的1 都变成0
        bits[zero] = false;
        zero -= 1;
    }

    //0 11111111100
    if zero == 0 {
        return -1;
    }

    //将该位置置1
    bits[zero] = true;

    //将最后几位都变成1,
    for i in 0..first - zero - 1 {
        bits[31 - i] = true;
    }

    let mut res = 0;
    for i in 1..32 {
        if bits[i] {
            res += 1 << 31 - i;
        }
    }
    res
}

//6. Z 字形变换
pub fn convert2(s: String, num_rows: i32) -> String {
    if num_rows == 1 {
        return s;
    }

    let mut matrix = vec![vec![]; num_rows as usize];
    let mut row = 0;
    let mut is_down = true;
    for c in s.chars() {
        matrix[row].push(c);
        if is_down {
            if row == num_rows as usize - 1 {
                is_down = false;
                row -= 1;
            } else {
                row += 1;
            }
        } else {
            if row == 0 {
                is_down = true;
                row += 1;
            } else {
                row -= 1;
            }
        }
    }
    matrix
        .iter()
        .flat_map(|chars| chars.iter())
        .collect()
}

pub fn convert(s: String, num_rows: i32) -> String {
    if num_rows == 1 {
        return s;
    }

    let mut matrix = vec![vec![]; num_rows as usize];
    let mut row = 0;
    let mut is_down = true;
    for c in s.chars() {
        matrix[row].push(c);
        if is_down {
            if row + 1 == num_rows as usize {
                is_down = false;
                row -= 1;
            } else {
                row += 1;
            }
        } else {
            if row == 0 {
                is_down = true;
                row += 1;
            } else {
                row -= 1;
            }
        }
    }
    let mut res = String::with_capacity(s.len());
    for chars in matrix.iter() {
        for char in chars.iter() {
            res.push(*char);
        }
    }
    res
}