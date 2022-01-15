fn main(){

}
//1716. 计算力扣银行的钱
pub fn total_money(n: i32) -> i32 {
    //每天应拿的钱 n / 7 + n % 7
    (1..=n).map(|n| (n - 1) / 7 + match n % 7 {
        0 => 7 ,
        mod_ => mod_,
    }).sum()
}