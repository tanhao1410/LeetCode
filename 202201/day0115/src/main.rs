fn main(){

}
//1716. 计算力扣银行的钱
pub fn total_money(n: i32) -> i32 {
    //每天应拿的钱 n / 7 + n % 7
    (1..=n).map(|n| (n - 1) / 7 + match n % 7 {
        0 => 7 ,
        mod_ => mod_,
    }).sum()
}

//6. Z 字形变换
pub fn convert2(s: String, num_rows: i32) -> String {
    if num_rows == 1{
        return s;
    }

    let mut matrix = vec![vec![];num_rows as usize];
    let mut row = 0;
    let mut is_down = true;
    for c in s.chars(){
        matrix[row].push(c);
        if is_down{
            if row == num_rows as usize - 1{
                is_down = false;
                row -= 1;
            }else{
                row += 1;
            }
        }else{
            if row == 0{
                is_down = true;
                row += 1;
            }else{
                row -= 1;
            }
        }
    }
    matrix
        .iter()
        .flat_map(|chars|chars.iter())
        .collect()
}

pub fn convert(s: String, num_rows: i32) -> String {
    if num_rows == 1{
        return s;
    }

    let mut matrix = vec![vec![];num_rows as usize];
    let mut row = 0;
    let mut is_down = true;
    for c in s.chars(){
        matrix[row].push(c);
        if is_down{
            if row + 1 == num_rows as usize{
                is_down = false;
                row -= 1;
            }else{
                row += 1;
            }
        }else{
            if row == 0{
                is_down = true;
                row += 1;
            }else{
                row -= 1;
            }
        }
    }
    let mut res = String::with_capacity(s.len());
    for chars in matrix.iter(){
        for char in chars.iter(){
            res.push(*char);
        }
    }
    res
}