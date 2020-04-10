package model

// Comment 商品评价
type Comment struct {
}

// TableName return table name
func (*Comment) TableName() string {
	return tablePrefix + "comment"
}

// CommentReplay 评价回复
type CommentReplay struct {
}

// TableName return table name
func (*CommentReplay) TableName() string {
	return tablePrefix + "comment_replay"
}
