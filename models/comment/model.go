package comment

type Comment struct {
	ID        int64
	Text      string
	ReplyID   int64
	VideoID   int64
	UserID    int64
	IsDeleted bool
}
