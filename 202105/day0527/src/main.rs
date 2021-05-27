fn main() {
    println!("Hello, world!");
}

//461. 汉明距离
pub fn hamming_distance(x: i32, y: i32) -> i32 {
    (x ^ y).count_ones() as i32
}