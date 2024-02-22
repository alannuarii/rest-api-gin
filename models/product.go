package models

type Product struct {
	Id int64 `gorm:"primaryKey" json:"id"`
	NamaProduct string `gorm:"type:varchar(50)" json:"nama_product"`
	Jumlah int64 `gorm:"type:int" json:"jumlah"`
	Harga int64 `gorm:"type:int" json:"harga"`
}