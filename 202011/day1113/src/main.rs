fn main() {
    println!("Hello, world!");
}

impl Solution {

    //153. 寻找旋转排序数组中的最小值
    pub fn find_min(nums: Vec<i32>) -> i32 {
        //数组大小为1，直接返回即可
        if nums.len() == 1 {
            return nums[0];
        }
        //最后一个比第一个大，说明是有序的，返回第一个数即可
        if nums[nums.len() - 1] > nums[0] {
            return nums[0];
        }

        let middle = nums[(nums.len() - 1) / 2];
        //中间的数比前面都要大，说明前面都是比较大的了。忽略掉即可
        if middle > nums[0] {
            Solution::find_min(nums[(nums.len() - 1) / 2 + 1..].to_vec())
        } else if middle == nums[0] {
            //中间的数与第一个数相同，说明此时的长度为2，直接返回第二个数即可。
            nums[1]
        } else {
            //中间的数比第一个数小，说明，最小的数在前面的范围里，该范围包括中间的数
            Solution::find_min(nums[1..(nums.len() - 1) / 2 + 1].to_vec())
        }
    }

    //150. 逆波兰表达式求值
    pub fn eval_rpn(tokens: Vec<String>) -> i32 {
        use std::str::FromStr;
        //遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中
        let mut stack = vec![];
        for i in tokens {
            match i.as_str() {
                "+" => {
                    let num2 = stack.pop().unwrap();
                    let num1 = stack.pop().unwrap();
                    stack.push(num1 + num2);
                }
                "-" => {
                    let num2 = stack.pop().unwrap();
                    let num1 = stack.pop().unwrap();
                    stack.push(num1 - num2);
                }
                "*" => {
                    let num2 = stack.pop().unwrap();
                    let num1 = stack.pop().unwrap();
                    stack.push(num1 * num2);
                }
                "/" => {
                    let num2 = stack.pop().unwrap();
                    let num1 = stack.pop().unwrap();
                    stack.push(num1 / num2);
                }
                num @ _ => {
                    stack.push(i32::from_str(num).unwrap());
                }
            }
        }
        stack.pop().unwrap()
    }

    //187. 重复的DNA序列
    pub fn find_repeated_dna_sequences(s: String) -> Vec<String> {
        //用hashmap暴力解决
        let mut res = vec![];
        if s.len() <= 10 {
            return res;
        }
        let mut m = std::collections::HashMap::new();
        let mut i = 10;
        let s = s.as_bytes();
        while i <= s.len() {
            //截取一个字符串
            let sub_s = String::from_utf8(s[i - 10..i].to_vec()).unwrap();
            if let Some(count) = m.get(&sub_s) {
                if *count == 1 {
                    res.push(sub_s.clone());
                }
                m.insert(sub_s, 2);
            } else {
                m.insert(sub_s, 1);
            }
            i += 1;
        }
        res
    }


    //152. 乘积最大子数组
    pub fn max_product(nums: Vec<i32>) -> i32 {

        //思路：用一个二维数组来记录中间结果
        let mut matrix = vec![];
        let mut res = nums2[0];

        //去除掉多余的1以及去除掉-1，
        let mut nums2 = vec![];
        for i in 0..nums.len(){
            if nums[i] != 1{
                nums2.push(nums[i]);
            }else{
                if res < 1{
                    res = 1;
                }
            }
        }

        let lenth = nums2.len();
        for i in &nums2 {
            matrix.push(vec![0; lenth]);
            if *i > res {
                res = *i;
            }
        }

        for i in (0..lenth).rev() {
            matrix[i][i] = nums2[i];
            let mut j = i + 1;
            while j < nums2.len() {
                matrix[i][j] = matrix[i][j - 1] * nums2[j];
                if matrix[i][j] > res {
                    res = matrix[i][j];
                }
                j += 1;
            }
        }
        res
    }

    pub fn odd_even_list(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut hleft_p = None;
        let (mut llast_pp, mut rlast_pp) = (&mut hleft_p, &mut head);

        while (*llast_pp).is_some() || (*rlast_pp).is_some() {

            //把llast_pp指向新的插入过来的奇节点
            *llast_pp = (*rlast_pp).take();
            //下一个偶数节点，即新奇节点的下一个节点
            llast_pp = &mut (*llast_pp).as_mut().unwrap().next;

            //下一个偶数点为空的话
            if (*llast_pp).is_none() {
                break;
            }
            //把这个偶数节点取出来
            *rlast_pp = (*llast_pp).take();

            //得到它后面的下一个奇数节点
            rlast_pp = &mut (*rlast_pp).as_mut().unwrap().next;
        }

        *llast_pp = head;
        hleft_p
    }
}

#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>
}
struct Solution{}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode {
            next: None,
            val
        }
    }
}