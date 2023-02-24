package message

var UserTimeMap map[string]int64

// 用来记录用户的最新聊天信息时间
func init() {
	UserTimeMap = make(map[string]int64)
}
