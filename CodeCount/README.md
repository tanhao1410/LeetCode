###减小rust编译后程序大小

[1].Cargo配置
```toml
[profile.release]
opt-level = 'z'
lto = true
```
[2].命令行参数
```shell script
cargo build --release
strip -s target/release/CodeCount.exe
```