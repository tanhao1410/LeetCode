package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9, 3, 2, 8, 4, 9}, 4))
}

//394. 字符串解码
func decodeString(s string) string {
	res := []byte{}

	for i := 0; i < len(s); i++ {
		if s[i] >= '1' && s[i] <= '9' {
			//需要处理数字了
			j := i + 1
			for ; j < len(s) && s[j] >= '1' && s[j] <= '9'; j++ {
			}
			//得到倍数
			count, _ := strconv.Atoi(s[i:j])
			//得到括号里面的内容
			left := 1
			k := j + 1
			for ; k < len(s); k++ {
				if s[k] == '[' {
					left++
				} else if s[k] == ']' {
					left--
				}
				if left == 0 {
					break
				}
			}
			//通过递归得到解码后的内部内容
			innerContent := decodeString(s[j+1 : k])
			for i := 0; i < count; i++ {
				res = append(res, innerContent...)
			}
			//跳过已解码的内容
			i = k
		} else {
			res = append(res, s[i])
		}
	}

	return string(res)
}

//355. 设计推特
type Twitter struct {
	//用户关注的列表
	UserFollows map[int][]int
	//用户的推文
	UserTweets map[int][]*Tweet
	//生成推文的序号
	Order int
}

type Tweet struct {
	TweetId int
	Order   int
}

func Constructor() Twitter {
	return Twitter{
		UserFollows: make(map[int][]int),
		UserTweets:  make(map[int][]*Tweet),
		Order:       0,
	}
}

//创建一条新的推文
func (this *Twitter) PostTweet(userId int, tweetId int) {
	//检查是否已经发表过,

	for k, v := range this.UserTweets[userId] {
		if v.TweetId == tweetId {
			//删除重发
			this.UserTweets[userId] = append(this.UserTweets[userId][:k], this.UserTweets[userId][k+1:]...)
			break
		}
	}

	tweet := Tweet{
		TweetId: tweetId,
		Order:   this.Order,
	}
	this.Order++
	if _, ok := this.UserTweets[userId]; ok {
		this.UserTweets[userId] = append(this.UserTweets[userId], &tweet)

		//排序推文
		sort.Slice(this.UserTweets[userId], func(i, j int) bool {
			return this.UserTweets[userId][i].Order > this.UserTweets[userId][j].Order
		})

	} else {
		this.UserTweets[userId] = []*Tweet{&tweet}
	}
}

//检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。推文必须按照时间顺序由最近的开始排序。
func (this *Twitter) GetNewsFeed(userId int) []int {
	//获取关注的人
	follows := this.UserFollows[userId]
	containSelf := false
	for _, v := range follows {
		if v == userId {
			containSelf = true
		}
	}
	if !containSelf {
		follows = append(follows, userId)
	}
	feeds := []*Tweet{}
	for _, v := range follows {
		//得到所有的推文,每个人的最多就十个,这里的排序有问题？拿出所有的推文，或每个人的前十条。
		for i := 0; i < 10 && i < len(this.UserTweets[v]); i++ {
			feeds = append(feeds, this.UserTweets[v][i])
		}
	}

	//排序推文
	sort.Slice(feeds, func(i, j int) bool {
		return feeds[i].Order > feeds[j].Order
	})

	res := []int{}
	for i := 0; i < 10 && i < len(feeds); i++ {
		res = append(res, feeds[i].TweetId)
	}
	return res
}

//关注一个用户
func (this *Twitter) Follow(followerId int, followeeId int) {

	//检查是否已关注过
	for _, v := range this.UserFollows[followerId] {
		if v == followeeId {
			return
		}
	}

	if _, ok := this.UserFollows[followerId]; ok {
		this.UserFollows[followerId] = append(this.UserFollows[followerId], followeeId)
	} else {
		this.UserFollows[followerId] = []int{followeeId}
	}
}

//取消关注一个用户
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := this.UserFollows[followerId]; ok {
		i := 0
		for ; i < len(this.UserFollows[followerId]); i++ {
			if this.UserFollows[followerId][i] == followeeId {
				break
			}
		}
		if i < len(this.UserFollows[followerId]) {
			this.UserFollows[followerId] = append(this.UserFollows[followerId][:i], this.UserFollows[followerId][i+1:]...)
		}
	}
}

//每日一题：714. 买卖股票的最佳时机含手续费
func maxProfit(prices []int, fee int) int {
	//思路：采用动态规划表，dp[i]代表price[i:]能挣的最多的钱数
	//l[i]代表，dp[i]购买的起点价
	dp, l := make([]int, len(prices)), make([]int, len(prices))
	max := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		//即有钱赚了
		if l[i+1] != 0 { // 这个if 判断不用加也可以的。
			if prices[i] < l[i+1] || prices[i]+fee < max {
				//比它后面购买时还要小，或比最大值要小
				//还要看max
				if max-prices[i]-fee > l[i+1]-prices[i] {
					dp[i] = dp[i+1] + max - prices[i] - fee
				} else {
					dp[i] = dp[i+1] - prices[i] + l[i+1]
				}
				l[i] = prices[i]
				max = 0
			} else {

				if prices[i] > max {
					max = prices[i]
				}
				dp[i] = dp[i+1]
				l[i] = l[i+1]
			}
		} else {
			//还没有赚钱时
			if prices[i] >= max {
				max = prices[i]
			} else {
				//比最大值要小
				if max-prices[i] > fee {
					dp[i] = max - prices[i] - fee
					l[i] = prices[i]
					max = 0
				}
			}
		}
	}
	return dp[0]
}
