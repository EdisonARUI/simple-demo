package model

type Video struct {
	VideoID   uint   `gorm:"primarykey;column:video_id;"`
	UserID    uint   `gorm:"column:user_id;not null"`
	Title     string `gorm:"column:title"`
	PlayUrl   string `gorm:"column:play_url"`
	CoverUrl  string `gorm:"column:cover_url"`
	CreatedAt int64
	Comments  []*Comment
	Likes     []*User `gorm:"many2many:like;joinForeignKey:video_id;joinReferences:user_id;"`
}

func (v *Video) TableName() string {
	return "video"
}
