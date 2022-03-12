fn main() {
    println!("Hello, world!");
}


impl Solution {
    //860. 柠檬水找零
    pub fn lemonade_change(bills: Vec<i32>) -> bool {
        //优先找零10元的。
        let mut five = 0;
        let mut ten = 0;
        for i in bills {
            match i {
                5 => five += 1,
                10 => {
                    five -= 1;
                    ten += 1;
                }
                _ => {
                    //需要找零15元
                    if ten > 0 {
                        ten -= 1;
                        five -= 1;
                    } else {
                        five -= 3;
                    }
                }
            }
            if five < 0 {
                return false;
            }
        }
        true
    }
}

use std::collections::BinaryHeap;

//1845. 座位预约管理系统
struct SeatManager {
    queue: BinaryHeap<i32>,
    n: i32,
}

impl SeatManager {
    fn new(n: i32) -> Self {
        let mut queue = BinaryHeap::new();
        for i in 1..=n {
            queue.push(i);
        }
        Self { queue, n }
    }

    fn reserve(&mut self) -> i32 {
        self.n - self.queue.pop().unwrap() + 1
    }

    fn unreserve(&mut self, seat_number: i32) {
        self.queue.push(self.n + 1 - seat_number);
    }
}