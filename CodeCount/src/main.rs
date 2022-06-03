use std::io::{BufReader, BufRead, stdin};
use std::fs::File;
use std::path::Path;

fn main() {
    let mut lang_vec = vec![Info::new("rust", ".rs"), Info::new("python", ".py"),
                            Info::new("java", ".java"), Info::new("golang", ".go")];

    lang_vec.sort_by_key(|info| usize::MAX - info.line_count);
    let (file_count, code_count) = lang_vec.iter().fold((0, 0), |pre, e| (pre.0 + e.file_count, pre.1 + e.line_count));
    lang_vec.iter().filter(|e| e.line_count > 0).for_each(|e| e.print(code_count));
    println!("  \t  总数：{}\t  总行数：{}\t  总：100%", file_count, code_count);
    stdin().read_line(&mut String::new());
}

struct Info {
    file_count: usize,
    line_count: usize,
    name: &'static str,
    extent_name: &'static str,
}

impl Info {
    pub fn new(name: &'static str, extent_name: &'static str) -> Self {
        let mut res = Self { file_count: 0, line_count: 0, extent_name, name };
        res.read_dir(Path::new("./"));
        res
    }

    pub fn print(&self, all_count: usize) {
        let num = self.line_count * 1000 / all_count;
        println!(" {}\t文件数：{}\t代码行数：{}\t占比：{}.{}%", self.name, self.file_count, self.line_count, num / 10, num % 10);
    }

    fn read_dir(&mut self, dir: &Path) {
        if !dir.ends_with("venv") && !dir.ends_with(".git") {
            for entry in dir.read_dir().unwrap() {
                let entry = entry.unwrap();
                if entry.file_type().unwrap().is_dir() {
                    self.read_dir(entry.path().as_path());
                } else if entry.file_name().to_str().unwrap().ends_with(self.extent_name) {
                    self.file_count += 1;
                    self.line_count += Self::file_line_count(entry.path().as_path());
                }
            }
        }
    }

    fn file_line_count(path: &Path) -> usize {
        let reader = BufReader::new(File::open(path).unwrap());
        reader.lines().map(|s| s.unwrap().trim().len() > 0).filter(|l| *l).count()
    }
}