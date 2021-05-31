fn main() {
    println!("Hello, world!");
}

//717. 1比特与2比特字符
pub fn is_one_bit_character(bits: Vec<i32>) -> bool {
    bits.iter().fold(0, |p, q| {
        match p {
            1 => *q + 2,
            _ => *q
        }
    }) == 0
}

