fn main() {
    println!("Hello, world!");
}

//844. 比较含退格的字符串
pub fn backspace_compare(s: String, t: String) -> bool {
    let process_str = |s: String| -> String{
        let bytes = s.as_bytes();
        let mut res_bytes = vec![];
        for i in 0..bytes.len() {
            if bytes[i] == b'#' && res_bytes.len() > 0 {
                if res_bytes.len() > 0 {
                    res_bytes.remove(res_bytes.len() - 1);
                }
            } else {
                res_bytes.push(bytes[i]);
            }
        }
        String::from_utf8(res_bytes).unwrap()
    };
    process_str(s) == process_str(t)
}

//1984. 学生分数的最小差值
pub fn minimum_difference(nums: Vec<i32>, k: i32) -> i32 {
    let mut nums = nums;
    nums.sort_unstable();
    let mut res = i32::MAX;
    for i in 0..nums.len() - k as usize + 1 as usize {
        res = res.min(nums[i + k as usize - 1 as usize] - nums[i]);
    }
    res
}