fn main() {
    println!("Hello, world!");
}

//1189. “气球” 的最大数量
pub fn max_number_of_balloons(text: String) -> i32 {
    //balloon
    let bytes = text.as_bytes();
    let (mut a, mut b, mut n, mut l, mut o) = (0, 0, 0, 0, 0);
    for &byte in bytes {
        match byte {
            b'a' => a += 1,
            b'b' => b += 1,
            b'n' => n += 1,
            b'l' => l += 1,
            b'o' => o += 1,
            _ => {}
        }
    }
    a.min(b).min(n).min(l/2).min(o/2)
}
