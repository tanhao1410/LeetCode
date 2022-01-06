use std::path::PathBuf;

fn main() {
    println!("Hello, world!");

    let string = Solution::simplify_path("///ab/..///.//".to_string());
    println!("{}",string);
}

struct Solution{}

impl Solution {
    ///71. 简化路径
    pub fn simplify_path( path: String) -> String {

        // let mut stk = vec![];
        // path.split('/').for_each(|s| match s {
        //     "." | "" => (),
        //     ".." => {
        //         stk.pop();
        //     }
        //     _ => stk.push(s),
        // });
        // "/".to_string() + &stk.join("/")

        let splits = path.split("/");
        let mut paths = vec![];
        for split in splits{
            match split {
                ".."=>{
                    paths.pop();
                },
                "." | "" =>{},
                _=>paths.push(split),
            }
        }
        "/".to_string() + & paths.join("/")
    }
}
