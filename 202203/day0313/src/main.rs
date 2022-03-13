fn main() {
    println!("Hello, world!");
}

//341. 扁平化嵌套列表迭代器
#[derive(Debug, PartialEq, Eq)]
pub enum NestedInteger {
    Int(i32),
    List(Vec<NestedInteger>),
}

struct NestedIterator {
    data: Vec<i32>,
    index: usize,
}

impl NestedIterator {
    fn new(nestedList: Vec<NestedInteger>) -> Self {
        let mut data = vec![];
        for l in nestedList {
            Self::read_nested(l, &mut data);
        }
        Self { data, index: 0 }
    }

    fn read_nested(nest: NestedInteger, v: &mut Vec<i32>) {
        match nest {
            NestedInteger::Int(val) => v.push(val),
            NestedInteger::List(list) => {
                for l in list {
                    Self::read_nested(l, v);
                }
            }
        }
    }

    fn next(&mut self) -> i32 {
        self.index += 1;
        self.data[self.index - 1]
    }

    fn has_next(&self) -> bool {
        self.index < self.data.len()
    }
}


//155. 最小栈
struct MinStack {
    //思路：用两个栈，一个只用来存储最小值
    min: Vec<i32>,
    data: Vec<i32>,

}

impl MinStack {
    fn new() -> Self {
        Self {
            min: vec![],
            data: vec![],
        }
    }

    fn push(&mut self, val: i32) {
        self.data.push(val);
        if self.min.is_empty() || self.min[self.min.len() - 1] >= val {
            self.min.push(val);
        }
    }

    fn pop(&mut self) {
        let res = self.data.pop().unwrap();
        if self.min[self.min.len() - 1] == res {
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
