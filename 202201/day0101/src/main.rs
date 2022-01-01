use std::str::FromStr;

impl Solution {
    fn diff_way_to_compute_vec(vec1: &[OperaOrNum]) -> Vec<i32> {
        let mut res = vec![];
        //大小为1 说明里面就一个数字
        if vec1.len() == 1 {
            res.push(match vec1[0] {
                OperaOrNum::NUM(i) => i,
                _ => unreachable!()
            });
        } else {
            //index为运算符下标
            for index in (1..vec1.len()).step_by(2) {
                //左边不包括运算符
                let left = Self::diff_way_to_compute_vec(&vec1[0..index]);
                //右边也不包括运算符
                let right = Self::diff_way_to_compute_vec(&vec1[index + 1..]);
                //组合下
                for num1 in left {
                    for num2 in &right {
                        res.push(match vec1[index] {
                            OperaOrNum::Oper(b'+') => num1 + *num2,
                            OperaOrNum::Oper(b'-') => num1 - *num2,
                            OperaOrNum::Oper(b'*') => num1 * *num2,
                            _ => unreachable!(),
                        })
                    }
                }
            }
        }
        res
    }

    //241. 为运算表达式设计优先级
    pub fn diff_ways_to_compute(expression: String) -> Vec<i32> {
        //思路：两个数的时候，就一种
        //三个数的时候，中间的数先和左边结合，1种，中间的数和右边结合，1中，共2种
        //四个数的时候，可以拆分为 左边一个数与后面三个数的组合，2种；左边两个数与后面两个数的组合，1种；左边三个数与右边一个数的组合，2种，供5种
        //先把表达式 分解为 运算符
        // 采用递归的方式，结束条件为只有两个数
        //先将表达式拆分，分为一个集合，里面有数字，运算符等
        let vec1 = OperaOrNum::parse_express(expression);
        Self::diff_way_to_compute_vec(&vec1)
    }

    ///2022. 将一维数组转变成二维数组
    pub fn construct2_d_array(original: Vec<i32>, m: i32, n: i32) -> Vec<Vec<i32>> {
        if m * n != original.len() as i32 {
            return vec![];
        }
        //original.chunks(n as usize).map(|v|)
        //original.chunks(n as usize).map(|x| x.to_vec()).collect()
        original.into_iter().fold(vec![], |mut res, v| {
            if res.last().is_none() || res.last().unwrap().len() == n as usize {
                res.push(vec![]);
            }
            res.last_mut().unwrap().push(v);
            res
        })
    }
}

enum OperaOrNum {
    Oper(u8),
    NUM(i32),
}

impl OperaOrNum {
    //将表达式拆分成
    fn parse_express(express: String) -> Vec<OperaOrNum> {
        let mut res = vec![];
        let mut start = 0;
        let bytes = express.as_bytes();
        for i in 0..bytes.len() {
            match bytes[i] {
                v @ b'+' | v @ b'-' |v @  b'*' => {
                    //将前面的数形成一个数放进结果集
                    let num = String::from_utf8_lossy(&bytes[start..i]);
                    let num = i32::from_str(num.as_ref()).unwrap();
                    res.push(OperaOrNum::NUM(num));
                    res.push(OperaOrNum::Oper(v));
                    start = i + 1;
                }
                _ => {}
            }
        }
        //最后一个数
        let num = String::from_utf8_lossy(&bytes[start..]);
        let num = i32::from_str(num.as_ref()).unwrap();
        res.push(OperaOrNum::NUM(num));
        res
    }
}

struct Solution {}


fn main() {
    println!("Hello, world!");
    let vec = Solution::diff_ways_to_compute(String::from("2*3-4*5"));
    println!("{:?}",vec);
}
