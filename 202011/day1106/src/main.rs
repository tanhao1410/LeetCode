fn main() {
    println!("Hello, world!");
}

struct Solution {}

impl Solution {
    //剑指 Offer 61. 扑克牌中的顺子
    pub fn is_straight(mut nums: Vec<i32>) -> bool {
        //思路2：用一个数组来记录所有的数
        let mut arrays = [0; 14];
        for i in nums {
            arrays[i as usize] += 1
        }
        let mut zero_count = arrays[0];
        //从第一个不是0的地方开始访问
        let mut i = 1;
        while arrays[i] == 0 {
            i += 1;
        }
        for j in 0..5 {
            if i + j == 14 {
                return zero_count == 5 - j;
            }
            if arrays[i + j] > 1 {
                return false;
            } else if arrays[i + j] == 0 {
                zero_count -= 1;
            }
        }
        zero_count == 0
    }

    //剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
    pub fn exchange(mut nums: Vec<i32>) -> Vec<i32> {
        //思路：双指针
        if nums.len() == 0 {
            return nums;
        }
        let (mut head, mut tail) = (0, nums.len() - 1);
        while head < tail {
            if nums[head] % 2 == 1 {
                //奇数
                head += 1;
            }
            if nums[tail] % 2 == 0 {
                tail -= 1;
            }
            if head < tail && nums[head] % 2 == 0 && nums[tail] % 2 == 1 {
                let temp = nums[head];
                nums[head] = nums[tail];
                nums[tail] = temp;
            }
        }
        nums
    }

    //剑指 Offer 15. 二进制中1的个数 191. 位1的个数 9 1001 2
    //n&(n−1) 解析： 二进制数字 nn 最右边的 11 变成 00 ，其余不变。
    pub fn hammingWeight(mut n: u32) -> i32 {
        //思路：&1，然后看等于即
        let mut res = 0;
        while n != 0 {
            res += n & 1;
            n >>= 1;
        }
        res as i32
    }

    //191. 位1的个数
    pub fn hammingWeight2(mut u: u32) -> i32 {
        let mut res = 0;
        while n != 0 {
            res += 1;
            n &= n - 1;
        }
        res
    }

    //1356. 根据数字二进制下 1 的数目排序
    pub fn sort_by_bits(mut arr: Vec<i32>) -> Vec<i32> {
        fn count_binary1(mut num: i32) -> i32 {
            let mut res = 0;
            while num != 0 {
                res += num & 1;
                num >>= 1;
            }
            res
        }
        fn compare(num1: i32, num2: i32) -> bool {
            //先比较1的数目，再比较大小
            if count_binary1(num1) > count_binary1(num2) {
                return true;
            } else if count_binary1(num1) == count_binary1(num2) {
                return num1 > num2;
            }
            false
        }

        for i in 0..arr.len() - 1 {
            for j in i..arr.len() {
                if compare(arr[i], arr[j]) {
                    let temp = arr[i];
                    arr[i] = arr[j];
                    arr[j] = temp;
                }
            }
        }
        arr
    }
}
