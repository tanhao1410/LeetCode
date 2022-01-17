fn main() {
    println!("Hello, world!");
}

pub fn count_vowel_permutation(n: i32) -> i32 {

    let MOD = 1000000007;
    //a,e,i,o,u
    //if a e a(n)=e(n-1)
    //if e a/i e(n) = a(n-1) + i(n - 1)
    //if i a/e/o/u i(n) = a(n-1) + e(n-1) + o(n - 1)+u(n - 1)
    //if o i/u o o(n) = i(n-1) + u(n - 1)
    //if u a u(n) = a(n - 1)
    //采用递归算法的话，应该要超时
    //a(n) + e(n) + i(n) + o(n) + u(n)
    let mut a = vec![1;n as usize];
    let mut e = vec![1;n as usize];
    let mut i_ = vec![1;n as usize];
    let mut o = vec![1;n as usize];
    let mut u = vec![1;n as usize];

    //从a能推断出u
    for i in 1..n as usize{
        a[i] = e[i - 1];
        e[i] = (a[i - 1] + i_[i -  1]) % MOD;
        i_[i] = ((a[i - 1] + e[i - 1])% MOD + (o[i - 1] + u[i - 1])% MOD)% MOD;
        o[i] = (i_[i - 1] + u[i - 1])% MOD;
        u[i] = a[i - 1];
    }

    (((a[n as usize - 1] + e[n as usize - 1] )%MOD+ (i_[n as usize - 1]

        + o[n as usize - 1])%MOD)%MOD + u[n as usize - 1])%MOD

    //let mut m: i64 = 1000000007;
    //         ((1..n).into_iter().fold((1, 1, 1, 1, 1, 5), |(a, e, i, o, u, sum), _| {
    //             let (_a, _e, _i, _o, _u) = ((e + i + u) % m, (a + i) % m, (e + o) % m, i, (i + o) % m);
    //             (_a, _e, _i, _o, _u, (_a + _e + _i + _o + _u))
    //         }).5 % m) as i32
    //
    // 作者：kyushu
    // 链接：https://leetcode-cn.com/problems/count-vowels-permutation/solution/rustgolangjava-dp-100-by-kyushu-zgmh/
    // 来源：力扣（LeetCode）
    // 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

}