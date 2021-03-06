fn main() {
    println!("Hello, world!");
}

impl Solution {

    //134. 加油站
    pub fn can_complete_circuit(gas: Vec<i32>, cost: Vec<i32>) -> i32 {

        //先找一个可以开始的起点，即cost[i]<=gas[i]
        for i in 0..gas.len(){
            //可以开始
            if cost[i]  <= gas[i]{
                let mut all = gas[i] - cost[i];
                let mut next = (i + 1) % gas.len();
                //如果油不够了，或完成了一圈退出循环
                while all >= 0 && next != i{
                    all += gas[next];
                    all -= cost[next];
                    next = (next+1) % gas.len()
                }
                if all >= 0 {
                    return i as i32;
                }
            }
        }
        -1
    }

    //223. 矩形面积
    pub fn compute_area(a: i32, b: i32, c: i32, d: i32, e: i32, f: i32, g: i32, h: i32) -> i32 {
        let all = (c - a).abs() * (d - b).abs() + (g - e).abs() * (h - f).abs();
        let width_start = if a > e { a } else { e };
        let width_end = if c > g { g } else { c };
        let height_start = if b > f { b } else { f };
        let height_end = if h > d { d } else { h };
        if width_end <= width_start || height_end <= height_start {
            return all;
        }
        all - (width_end - width_start) * (height_end - height_start)
    }

    //922. 按奇偶排序数组 II
    pub fn sort_array_by_parity_ii(mut a: Vec<i32>) -> Vec<i32> {
        //思路：两指针，每个走两格，遇到不符合的停下来，和另外一个不符合的交换
        let (mut even, mut odd) = (0, 1);
        while odd < a.len() && even < a.len() {
            while even < a.len() && a[even] % 2 == 0 {
                even += 2;
            }
            while odd < a.len() && a[even] % 2 == 1 {
                odd += 2;
            }
            if even < a.len() && odd < a.len() {
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

struct Solution {}