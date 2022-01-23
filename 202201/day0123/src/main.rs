#![feature(map_first_last)]


use std::collections::{HashMap, BTreeMap};

fn main() {
    println!("Hello, world!");

    assert_eq!(max_product(vec![0, 2]), 2);
    assert_eq!(max_product(vec![2, -5, -2, -4, 3]), 24);
    assert_eq!(max_product(vec![-2, 3, -4]), 24);
    assert_eq!(max_product(vec![-2, 0, -1]), 0);
    assert_eq!(max_product(vec![2, 3, -2, 4]), 6);

    assert_eq!(get_max_len(vec![0, 1, -2, -3, -4]), 3);
}

//2034. 股票价格波动
struct StockPrice {
    current: i32,
    prices: HashMap<i32, i32>,
    //用一个B树存储所有的价格，value 为该价格的数量
    values: BTreeMap<i32, i32>,
}

impl StockPrice {
    fn new() -> Self {
        Self{
            current:0,
            prices:HashMap::new(),
            values:BTreeMap::new(),
        }
    }

    fn update(&mut self, timestamp: i32, price: i32) {
        //看这个时间之前是否存在
        if let Some(value) = self.prices.get(&timestamp){
            //之前存在
            //在values中删除原来的
            let count = self.values.get_mut(&value).unwrap();
            *count -= 1;
            if *count == 0{
                self.values.remove(&price);
            }
        }
        self.prices.insert(timestamp,price);
        let entry = self.values.entry(price).or_insert(0);
        *entry += 1;
        if timestamp > self.current{
            self.current = timestamp;
        }
    }

    fn current(&self) -> i32 {
        *self.prices.get(&self.current).unwrap()
    }

    fn maximum(&self) -> i32 {
        *self.values.last_key_value().unwrap().0
    }

    fn minimum(&self) -> i32 {
        *self.values.iter().next().unwrap().0
    }
}

//1567. 乘积为正数的最长子数组长度
pub fn get_max_len(nums: Vec<i32>) -> i32 {
    //思路：用两个dp，一个记录以nums[i]结尾的最长乘积正数子数组长度，一个记录最长乘积负数子数组长度。
    let mut dp = vec![0; nums.len()];
    let mut dp2 = vec![0; nums.len()];
    if nums[0] > 0 {
        dp[0] = 1;
    } else if nums[0] < 0 {
        dp2[0] = 1;
    }
    let mut res = dp[0];
    for i in 1..nums.len() {
        if nums[i] == 0 {
            dp[i] = 0;
            dp2[i] = 0;
        } else if nums[i] > 0 {
            //
            dp[i] = 1 + dp[i - 1];
            //只有前面能形成负数的时候才会+1，否则不会+的
            if dp2[i - 1] > 0 {
                dp2[i] = 1 + dp2[i - 1];
            }
        } else {
            //如果当前数小于0，
            if dp2[i - 1] > 0 {
                dp[i] = 1 + dp2[i - 1];
            }
            dp2[i] = 1 + dp[i - 1];
        }
        res = res.max(dp[i]);
    }
    res
}

//2034. 股票价格波动-时间超时
struct StockPrice2 {
    //难点在于，每次更新操作后，会对系统的最大值和最小值产生影响。每次更新，删除一个值的时候
    //专门用一个结构来记录所有值，每次来插入一个，每次删除一个。olongN.
    current: i32,
    //最新时间
    values: Vec<i32>,
    //所有的股价
    map: HashMap<i32, i32>,//时间-股价
}

impl StockPrice2 {
    fn new() -> Self {
        Self {
            current: 0,
            values: vec![],
            map: HashMap::new(),
        }
    }

    //二分法定位，或插入位置
    fn location(&self, v: i32) -> usize {
        let values = &self.values;
        let mut i = 0;
        let mut j = values.len() - 1;
        let mut middle = (j - i) / 2 + i;
        while j >= i {
            if values[middle] == v {
                return middle;
            } else if values[middle] < v {
                i = middle + 1;
            } else {
                if middle == 0 {
                    break;
                }
                j = middle - 1;
            }
            middle = (j - i) / 2 + i;
        }
        i
    }

    fn update(&mut self, timestamp: i32, price: i32) {
        if let Some(value) = self.map.get(&timestamp) {
            //从values中删除该value
            let location = self.location(*value);
            self.values.remove(location);
        }
        self.map.insert(timestamp, price);
        //插入新的value
        self.values.insert(self.location(price), price);
        //更新最新时间
        if timestamp > self.current {
            self.current = timestamp;
        }
    }

    fn current(&self) -> i32 {
        *self.map.get(&self.current).unwrap()
    }

    fn maximum(&self) -> i32 {
        self.values[self.values.len() - 1]
    }

    fn minimum(&self) -> i32 {
        self.values[0]
    }
}

//167. 两数之和 II - 输入有序数组
pub fn two_sum(numbers: Vec<i32>, target: i32) -> Vec<i32> {
    let mut i = 0;
    let mut j = numbers.len() - 1;
    while numbers[i] + numbers[j] != target {
        if numbers[i] + numbers[j] > target {
            j -= 1;
        } else {
            i += 1;
        }
    }
    vec![i as i32 + 1, j as i32 + 1]
}

//283. 移动零
pub fn move_zeroes(nums: &mut Vec<i32>) {
    //双指针，i指向第一个为0的地方，j指向i + 1，j不停的往前走，碰到不为0的就将它放到i处，同时，i 放到j位置
    let mut i = 0;
    while i < nums.len() && nums[i] != 0 {
        i += 1;
    }
    let mut j = i + 1;
    while j < nums.len() {
        if nums[j] != 0 {
            nums[i] = nums[j];
            nums[j] = 0;
            i += 1;
        }
        j += 1;
    }
}

//152. 乘积最大子数组
pub fn max_product(nums: Vec<i32>) -> i32 {
    let mut dp_max = vec![0; nums.len()]; //以nums[i]结尾的最大子数组之积
    let mut dp_min = vec![0; nums.len()]; //以nums[i]结尾的负的最小值。

    dp_max[0] = nums[0];
    if nums[0] < 0 {
        dp_min[0] = nums[0];
    }

    for i in 1..nums.len() {
        if nums[i] >= 0 {
            //看它的最大值是否大于0，如果大于0，则dp_max[i] = nums[i] * dp[i - 1]
            //如果它的前一个最大值小于0，自己又大于0，所以，dp_max[i] = nums[i]
            //如果它的前一个等于0，自己又大于0，dp_max[i] = nums[i]
            //更新dp_min
            //dp_min[i] = nums[i] * dp_min[i - 1]，因为自己大于0，而前面的最大负，乘上正数，肯定更小
            if dp_max[i - 1] > 0 {
                dp_max[i] = nums[i] * dp_max[i - 1];
            } else {
                dp_max[i] = nums[i];
            }
            dp_min[i] = nums[i] * dp_min[i - 1];
        } else {
            //如果当前值小于0
            //如果dp_max[i - 1] > 0,dp_min[i - 1] < 0,=>dp_max[i] = nums[i] * dp_min[i - 1]
            //如果dp_max[i - 1] > 0,_ =>dp_max[i] = nums[i],
            //
            if dp_max[i - 1] > 0 {
                if dp_min[i - 1] < 0 {
                    //更新最大值
                    dp_max[i] = nums[i] * dp_min[i - 1];
                } else {
                    dp_max[i] = nums[i];
                }
                //更新负的最小值
                dp_min[i] = nums[i] * dp_max[i - 1];
            } else {
                dp_max[i] = dp_max[i - 1].min(dp_min[i - 1]) * nums[i];
                dp_min[i] = nums[i];
            }
        }
    }
    dp_max.into_iter().max().unwrap()
}