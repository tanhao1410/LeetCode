fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //1089. 复写零
    pub fn duplicate_zeros(arr: &mut Vec<i32>) {
        let arr2 = arr.clone();
        let mut index = 0;
        let mut index2 = 0;
        while index < arr.len() {
            if arr2[index2] == 0 {
                arr[index] = 0;
                if index + 1 < arr.len() {
                    arr[index + 1] = 0;
                }
                index += 2;
            } else {
                arr[index] = arr2[index2];
                index += 1;
            }
            index2 += 1;
        }
    }
}