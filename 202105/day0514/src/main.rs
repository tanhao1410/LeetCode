fn main() {
    println!("Hello, world!");
}

//12. 整数转罗马数字 - 借鉴思路：贪心算法
pub fn int_to_roman2(num: i32) -> String {
    [(1000, "M"), (900, "CM"), (500, "D"), (400, "CD"),
        (100, "C"), (90, "XC"), (50, "L"), (40, "XL"),
        (10, "X"), (9, "IX"), (5, "V"), (4, "IV"),
        (1, "I")].into_iter()
        .fold((String::with_capacity(20), num), |(mut s, mut num), (base, unit)|
            (s + &unit.repeat((num / base) as usize), num % base))
        .0
}

//12. 整数转罗马数字
pub fn int_to_roman(num: i32) -> String {
    //num <= 3999
    //I             1
    // V             5
    // X             10
    // L             50
    // C             100
    // D             500
    // M             1000
    //思路：先求千，再求百
    let mut num = num;
    let mut res = Self::into_roman_part(&mut num, 1000, 'M', 'o', 'o');
    res.push_str(Self::into_roman_part(&mut num, 100, 'C', 'D', 'M').as_str());
    res.push_str(Self::into_roman_part(&mut num, 10, 'X', 'L', 'C').as_str());
    res.push_str(Self::into_roman_part(&mut num, 1, 'I', 'V', 'X').as_str());
    res
}
// 罗马数字转换工具，1000一下，992 ， 100，D,C => DCD num =>92 900 也是特殊的
pub fn into_roman_part(num: &mut i32, base: i32, base_id: char, five_id: char, ten_id: char) -> String {
    let mut res = String::new();

    // 处理 900， 90 等
    if *num >= base * 9 {
        res.push(base_id);
        res.push(ten_id);
        *num -= base * 9;
    }

    //处理大于500的
    if *num >= 5 * base {
        res.push(five_id);
        *num -= 5 * base;
    }

    //处理 400,40
    if *num >= base * 4 {
        res.push(base_id);
        res.push(five_id);
        *num -= 4 * base;
    }

    for i in 0..*num / base {
        res.push(base_id);
    }
    *num %= base;
    res
}
