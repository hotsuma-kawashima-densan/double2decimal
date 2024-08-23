package main

import (
	"context"
	"fmt"
	"project+test/pkg/db/model"
	"project+test/pkg/db/query"

	"github.com/shopspring/decimal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:13306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("接続失敗:", err)
		return
	}

	qtx := query.Use(db)
	data, err := qtx.Temp.WithContext(context.Background()).Find()

	if err != nil {
		fmt.Println("データ取得失敗:", err)
		return
	}

	for i, data := range data {
		fmt.Printf("id: %d, input: %s, A: %f, B: %s, B0: %s, B1: %s, B2: %s, B3: %s, B4: %s\n", i, data.Input, data.A, data.B, data.B0, data.B1, data.B2, data.B3, data.B4)
	}

	// id: 0,  input: 1E-1,     A: 0.100000,   B: 0,   B0: 0,   B1: 0.1,   B2: 0.1,    B3: 0.1,    B4: 0.1
	// id: 1,  input: 1E-2,     A: 0.010000,   B: 0,   B0: 0,   B1: 0,     B2: 0.01,   B3: 0.01,   B4: 0.01
	// id: 2,  input: 1E-3,     A: 0.001000,   B: 0,   B0: 0,   B1: 0,     B2: 0,      B3: 0.001,  B4: 0.001
	// id: 3,  input: 1E-4,     A: 0.000100,   B: 0,   B0: 0,   B1: 0,     B2: 0,      B3: 0,      B4: 0.0001
	// id: 4,  input: 1E-5,     A: 0.000010,   B: 0,   B0: 0,   B1: 0,     B2: 0,      B3: 0,      B4: 0
	// id: 5,  input: 5E-5,     A: 0.000050,   B: 0,   B0: 0,   B1: 0,     B2: 0,      B3: 0,      B4: 0.0001
	// id: 6,  input: 0.0006,   A: 0.000600,   B: 0,   B0: 0,   B1: 0,     B2: 0,      B3: 0.001,  B4: 0.0006
	// id: 7,  input: 0.0004,   A: 0.000400,   B: 0,   B0: 0,   B1: 0,     B2: 0,      B3: 0,      B4: 0.0004
	// id: 8,  input: 0.006,    A: 0.006000,   B: 0,   B0: 0,   B1: 0,     B2: 0.01,   B3: 0.006,  B4: 0.006
	// id: 9,  input: 0.004,    A: 0.004000,   B: 0,   B0: 0,   B1: 0,     B2: 0,      B3: 0.004,  B4: 0.004
	// id: 10, input: 1.0,      A: 1.000000,   B: 1,   B0: 1,   B1: 1,     B2: 1,      B3: 1,      B4: 1
	// id: 11, input: 281.12,   A: 281.120000, B: 281, B0: 281, B1: 281.1, B2: 281.12, B3: 281.12, B4: 281.12
	// id: 12, input: 0.938,    A: 0.938000,   B: 1,   B0: 1,   B1: 0.9,   B2: 0.94,   B3: 0.938,  B4: 0.938
	// id: 13, input: 15.315,   A: 15.315000,  B: 15,  B0: 15,  B1: 15.3,  B2: 15.32,  B3: 15.315, B4: 15.315
	// id: 14, input: 1.234567, A: 1.234567,   B: 1,   B0: 1,   B1: 1.2,   B2: 1.23,   B3: 1.235,  B4: 1.2346

	var floatVal float64 = 9.45
	var decimalVal decimal.Decimal = decimal.NewFromFloat(9.45)

	_data15 := model.Temp{
		ID:    15,
		Input: "9.45",
		A:     floatVal,
		B:     decimalVal,
		B0:    decimalVal,
		B1:    decimalVal,
		B2:    decimalVal,
		B3:    decimalVal,
		B4:    decimalVal,
	}

	err = qtx.Temp.WithContext(context.Background()).Create(&_data15)

	if err != nil {
		fmt.Println("データ登録失敗:", err)
		return
	}

	data15, err := qtx.Temp.WithContext(context.Background()).Last()

	if err != nil {
		fmt.Println("データ取得失敗:", err)
		return
	}

	fmt.Printf("id: %d, input: %s, A: %.16f, B: %s, B0: %s, B1: %s, B2: %s, B3: %s, B4: %s\n", data15.ID, data15.Input, data15.A, data15.B, data15.B0, data15.B1, data15.B2, data15.B3, data15.B4)
	// id: 15, input: 9.45, A: 9.4499999999999993, B: 9, B0: 9, B1: 9.5, B2: 9.45, B3: 9.45, B4: 9.45
}
