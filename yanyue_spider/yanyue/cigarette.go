package yanyue

type Product struct {
	Id    int32  `json:"productid"`
	Name  string `json:"productname"`
	Brand Brand  `json:"brand"`
	// 每条价格
	BarPrice float32 `json:"barprice"`
	// 每包的价格
	PackPrice float32 `json:"packprice"`
	// 评论数
	CommentNum int32 `json:"commentnum"`
	// 评分数
	PingNum int32 `json:"pingnum"`
	// 综合得分
	ComSore float32 `json:"comscore"`
	// 口味得分
	ScoreWei float32 `json:"scorewei"`
	// 外观得分
	ScoreBao float32 `json:"scorebao"`
	// 性价比得分
	ScoreJia float32 `json:"scorejia"`
}

type Brand struct {
	Id   string `json:"brandid"`
	Name string `json:"brandname"`
}
