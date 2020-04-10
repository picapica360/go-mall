package model

// Brand 品牌信息
type Brand struct {
}

// BrandInParam brand input param
type BrandInParam struct {
}

// TableName return table name
func (*Brand) TableName() string {
	return tablePrefix + "brand"
}
