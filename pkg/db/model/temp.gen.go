// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/shopspring/decimal"

const TableNameTemp = "temp"

// Temp mapped from table <temp>
type Temp struct {
	ID    int32           `gorm:"column:id;not null" json:"id"`
	Input string          `gorm:"column:input;not null" json:"input"`
	A     float64         `gorm:"column:A;not null" json:"A"`
	B     decimal.Decimal `gorm:"column:B;not null" json:"B"`
	B0    decimal.Decimal `gorm:"column:B0;not null" json:"B0"`
	B1    decimal.Decimal `gorm:"column:B1;not null" json:"B1"`
	B2    decimal.Decimal `gorm:"column:B2;not null" json:"B2"`
	B3    decimal.Decimal `gorm:"column:B3;not null" json:"B3"`
	B4    decimal.Decimal `gorm:"column:B4;not null" json:"B4"`
}

// TableName Temp's table name
func (*Temp) TableName() string {
	return TableNameTemp
}
