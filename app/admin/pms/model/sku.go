package model

// SKUStock sku 库存
type SKUStock struct {
}

// TableName return table name
func (*SKUStock) TableName() string {
	return tablePrefix + "sku_stock"
}
