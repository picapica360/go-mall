package model

// Product 商品信息
type Product struct {
}

// TableName return table name
func (*Product) TableName() string {
	return tablePrefix + "product"
}

// ProductVertifyRecord 产品审核记录
type ProductVertifyRecord struct {
}

// TableName return table name
func (*ProductVertifyRecord) TableName() string {
	return tablePrefix + "product_vertify_record"
}
