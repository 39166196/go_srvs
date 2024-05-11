package model

type Inventory struct {
	BaseModel
	Goods   int32 `gorm:"type:int;index"`
	Stocks  int32 `gorm:"type:int"`
	Version int32 `gorm:"type:int"`
}
type InventoryNew struct {
	BaseModel
	Goods   int32 `gorm:"type:int;index"`
	Stocks  int32 `gorm:"type:int"`
	Version int32 `gorm:"type:int"`
	Freeze  int32 `gorm:"type:int"`
}

//	type Stock struct {
//		BaseModel
//		Name    string
//		Address string
//	}
//type InventoryHistory struct {
//	user   int32
//	inventory  int32
//	nums   int32
//	order  int32
//	status int32
//}
