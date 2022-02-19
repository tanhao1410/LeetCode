fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //969. 煎饼排序
    pub fn pancake_sort(mut arr: Vec<i32>) -> Vec<i32> {
        let mut res = vec![];
        let reverse_arr = |arr: &mut Vec<i32>, mut start: usize, mut end: usize| {
            while start < end {
                let temp = arr[start];
                arr[start] = arr[end];
                arr[end] = temp;
                start += 1;
                end -= 1;
            }
        };
        for i in (1..=arr.len()).rev() {
            if arr[i - 1] != i as i32 {
                //需要翻转
                for j in 0..i {
                    if arr[j] == i as i32 {
                        if j != 0 {
                            //先翻转一次，将这个数翻转到开头
                            res.push(j as i32 + 1);
                            reverse_arr(&mut arr, 0, j);
                        }
                        res.push(i as i32);
                        reverse_arr(&mut arr, 0, i - 1);
                        break;
                    }
                }
            }
        }
        res
    }
}