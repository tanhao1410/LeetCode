fn main() {
    println!("Hello, world!");
}

//692. 前K个高频单词
pub fn top_k_frequent(words: Vec<String>, k: i32) -> Vec<String> {
    // let mut dic = std::collections::HashMap::new();
    // words.iter().for_each(|&w| *dic.entry(w).or_insert(0)+= 1);
    let mut dic = words.iter()
        .fold(&mut std::collections::HashMap::new(), |mut m, w| {
            *m.entry(w).or_insert(0) += 1;
            m
        })
        .iter()
        .map(|(&k, &v)| (k, v))
        .collect::<Vec<_>>();
    dic.sort_by(|pre, pro|
        match pro.1.partial_cmp(&pre.1).unwrap() {
            std::cmp::Ordering::Equal => pre.0.partial_cmp(&pro.0).unwrap(),
            _ => pro.1.partial_cmp(&pre.1).unwrap()
        }
    );
    dic.iter()
        .take(k as usize)
        .map(|wc| wc.0.clone())
        .collect::<Vec<_>>()
}