

fn main() {
    println!("Hello, world!");
}

struct Solution{

}

impl Solution {
    //每日一题：1207.独一无二的出现次数
    pub fn unique_occurrences(arr: Vec<i32>) -> bool {
        //思路；用map来记录数字出现的次数（也可以直接用数组来记录，因为数的大小是固定的）
        use std::collections::HashMap;
        use std::collections::HashSet;
        let mut m = HashMap::new();
        for i in arr.iter(){
            if m.contains_key(i){
                m.insert(*i,m.get(i).unwrap() + 1);
            }else{
                m.insert(*i,1);
            }
        }
        let mut s = HashSet::new();
        for i in m.iter(){
            if s.contains(i.1){
                return false;
            }else{
                s.insert(*i.1);
            }
        }
        true
    }
}