fn main() {
    println!("Hello, world!");
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