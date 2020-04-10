package model

// ProductAttrCategory 产品属性分类
type ProductAttrCategory struct {
	ID         int64
	Name       string
	AttrCount  int // 属性数量
	ParamCount int // 参数数量
}

// ProductAttr 商品属性参数
type ProductAttr struct {
	ID                    int64
	ProductAttrCategoryID int64 // ref->ProductAttrCategory
	Name                  string
	SelectType            int    // 属性选择类型：0->唯一；1->单选；2->多选
	InputType             int    // 属性录入方式：0->手工录入；1->从列表中选取
	InputList             string // 可选值列表，以逗号隔开
	Sort                  int    // 排序字段：最高的可以单独上传图片
	FilterType            int    // 分类筛选样式：1->普通；1->颜色
	SearchType            int    // 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索
	RelatedStatus         int    // 相同属性产品是否关联；0->不关联；1->关联
	HandAddStatus         int    // 是否支持手动新增；0->不支持；1->支持
	Kind                  int    // 属性的类型；0->规格；1->参数
}

// ProductAttrValue 存储产品参数信息
type ProductAttrValue struct {
	ID            int64
	ProductID     int64  // ref->Product
	ProductAttrID int64  // ref->ProductAttr
	Value         string // 手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开
}
