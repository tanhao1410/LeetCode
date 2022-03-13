fn main() {
    println!("Hello, world!");
}

//155. 最小栈
struct MinStack {
    //思路：用两个栈，一个只用来存储最小值
    min:Vec<i32>,
    data:Vec<i32>,

}
impl MinStack {

    fn new() -> Self {
        Self{
            min:vec![],
            data:vec![],
        }
    }

    fn push(&mut self, val: i32) {
        self.data.push(val);
        if self.min.is_empty() || self.min[self.min.len() - 1] >= val{
            self.min.push(val);
        }
    }

    fn pop(&mut self) {
        let res = self.data.pop().unwrap();
        if self.min[self.min.len() - 1] == res{
            self.min.pop();
        }
    }

    fn top(&self) -> i32 {
        self.data[self.data.len() - 1]
    }

    fn get_min(&self) -> i32 {
        self.min[self.min.len() - 1]
    }
}
