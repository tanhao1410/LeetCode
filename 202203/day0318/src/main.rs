fn main() {
    println!("Hello, world!");
}

//2043. 简易银行系统
struct Bank {
    balance: Vec<i64>,
}

impl Bank {
    fn new(balance: Vec<i64>) -> Self {
        Self { balance }
    }

    fn transfer(&mut self, account1: i32, account2: i32, money: i64) -> bool {
        if account1 as usize > self.balance.len() || account2 as usize > self.balance.len() || money > self.balance[account1 as usize - 1] {
            return false;
        }
        self.balance[account1 as usize - 1] -= money;
        self.balance[account2 as usize - 1] += money;
        true
    }

    fn deposit(&mut self, account: i32, money: i64) -> bool {
        if account as usize > self.balance.len() {
            return false;
        }
        self.balance[account as usize - 1] += money;
        true
    }

    fn withdraw(&mut self, account: i32, money: i64) -> bool {
        if account as usize > self.balance.len() || money > self.balance[account as usize - 1] {
            return false;
        }
        self.balance[account as usize - 1] -= money;
        true
    }
}
