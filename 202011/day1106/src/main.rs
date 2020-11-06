fn main() {
    println!("Hello, world!");

}

struct Solution {}

impl Solution {

    //剑指 Offer 15. 二进制中1的个数 191. 位1的个数 9 1001 2
    //n&(n−1) 解析： 二进制数字 nn 最右边的 11 变成 00 ，其余不变。
    pub fn hammingWeight(mut n: u32) -> i32 {
        //思路：&1，然后看等于即
        let mut res = 0;
        while n != 0 {
            res += n & 1;
            n>>=1;
        }
        res as i32
    }

    //1356. 根据数字二进制下 1 的数目排序
    pub fn sort_by_bits(mut arr: Vec<i32>) -> Vec<i32> {
        fn count_binary1(mut num :i32)->i32{
            let mut res = 0;
            while num != 0{
                res += num & 1;
                num >>= 1;
            }
            res
        }
        fn compare(num1: i32, num2: i32) -> bool {
            //先比较1的数目，再比较大小
            if count_binary1(num1) > count_binary1(num2){
                return true;
            }else if count_binary1(num1) == count_binary1(num2){
                return num1 > num2;
            }
            false
        }

        for i in 0..arr.len() -1{
            for j in i..arr.len(){
                if compare(arr[i],arr[j]){
                    let temp = arr[i];
                    arr[i] = arr[j];
                    arr[j] = temp;
                }
            }
        }
        arr
    }
}
