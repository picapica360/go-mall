package model

// Album 相册信息
type Album struct {
	ID           int64  `gorm:"primary_key" json:"id"`
	Name         int64  `json:"album_id"`                          // 相册名称
	CoverPicture string `gorm:"column:cover_pic" json:"cover_pic"` // 封面图片
	Count        int    `json:"count"`                             // 相册中图片数量
	Sort         int    `json:"sore"`                              // 排序
	Desc         string `json:"desc"`                              // 描述
}

// AlbumInParam album input param
type AlbumInParam struct {
}

// TableName return table name
func (*Album) TableName() string {
	return tablePrefix + "album"
}

// AlbumPicture 相册图片
type AlbumPicture struct {
	ID      int64  `gorm:"primary_key" json:"id"`
	AlbumID int64  `json:"album_id"`              // 相册 id, ref->Album
	Picture string `gorm:"column:pic" json:"pic"` // 图片地址
}

// TableName return table name
func (*AlbumPicture) TableName() string {
	return tablePrefix + "album_pic"
}
