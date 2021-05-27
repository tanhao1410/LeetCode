fn main() {
    println!("Hello, world!");
}

//371. 两整数之和
pub fn get_sum(a: i32, b: i32) -> i32 {
    //从最低位开始，&，^,
    let (mut a, mut b, mut res, mut flag) = (a, b, 0, 0);
    for i in 0..32 {
        let bit_num = match ((a & 1) ^ (b & 1), (a & 1) & (b & 1), flag) {
            (0, 0, 1) => {
                flag = 0;
                1
            }
            (0, 1, 0) => {
                flag = 1;
                0
            }
            (0, 1, 1) | (1, 0, 0) => 1,
            _ => 0
        };
        a >>= 1;
        b >>= 1;
        res |= bit_num << i
    }
    res
}

//461. 汉明距离
pub fn hamming_distance(x: i32, y: i32) -> i32 {
    (x ^ y).count_ones() as i32
}