fn main() {
    println!("Hello, world!");
}

impl Solution {
    //922. 按奇偶排序数组 II
    pub fn sort_array_by_parity_ii(mut a: Vec<i32>) -> Vec<i32> {
        //思路：两指针，每个走两格，遇到不符合的停下来，和另外一个不符合的交换
        let (mut even,mut odd)= (0,1);
        while odd < a.len() && even < a.len(){
            while even < a.len() && a[even] % 2 == 0{
                even += 2;
            }
            while odd < a.len() && a[even] % 2 == 1{
                odd += 2;
            }
            if even < a.len() && odd < a.len(){
                let temp = a[even];
                a[even] = a[odd];
                a[odd] = temp;
                even += 2;
                odd += 2;
            }
        }
        a
    }
}

struct Solution{}