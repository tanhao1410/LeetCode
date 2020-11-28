fn main() {
    println!("Hello, world!");
}

struct Solution {}

impl Solution {
    //每日一题：976. 三角形的最大周长
    pub fn largest_perimeter(a: Vec<i32>) -> i32 {
        //排序后，从后往前遍历，看最小的两条边之和是否大于第三边，如果大于，说明是三角形。
        //不大于的话，说明最大的这个肯定不符合了
        let mut a = a;
        a.sort_unstable();
        let mut res = 0;
        for num in (2..a.len()).rev() {
            if a[num] < a[num - 1] + a[num - 2] {
                res = a[num] + a[num - 1] + a[num - 2];
                break;
            }
        }
        res
    }
}
