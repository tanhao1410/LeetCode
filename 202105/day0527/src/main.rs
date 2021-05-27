fn main() {
    println!("Hello, world!");
}

//384. 打乱数组
struct Solution {
    nums: Vec<i32>
}

impl Solution {
    fn new(nums: Vec<i32>) -> Self {
        Solution {
            nums
        }
    }

    fn reset(&self) -> Vec<i32> {
        self.nums.clone()
    }

    fn rand(&self, max: usize) -> usize {
        //std::time::SystemTime::now().duration_since(std::time::UNIX_EPOCH).ok().unwrap().as_millis() % usize
        let mut rand_num = 0;
        unsafe {
            std::arch::x86_64::_rdrand16_step(&mut rand_num);
        }
        rand_num as usize % max
    }

    fn shuffle(&self) -> Vec<i32> {
        //随机进行
        let mut res = vec![];
        let mut temp = self.nums.clone();
        for i in (1..self.nums.len() + 1).rev() {
            let index = self.rand(i);
            // let last_num = temp[i - 1];
            // res.push(std::mem::replace(&mut temp[index],last_num));
            unsafe {
                std::ptr::swap(&mut temp[i - 1], &mut temp[index]);
            }
            res.push(temp[i - 1]);
        }
        res
    }
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