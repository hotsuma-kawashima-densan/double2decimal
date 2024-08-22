package main

import (
	"context"
	"fmt"
	"project+test/pkg/db/query"

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
		fmt.Printf("id: %d, input: %d, A: %f, B: %f, B0: %f, B1: %f, B2: %f, B3: %f, B4: %f\n", i, data.Input, data.A, *data.B, *data.B0, *data.B1, *data.B2, *data.B3, *data.B4)
	}

	// id: 0, input: %!d(string=1E-1),      A: 0.100000,   B: 0.000000,   B0: 0.000000,   B1: 0.100000,   B2: 0.100000,   B3: 0.100000,   B4: 0.100000
	// id: 1, input: %!d(string=1E-2),      A: 0.010000,   B: 0.000000,   B0: 0.000000,   B1: 0.000000,   B2: 0.010000,   B3: 0.010000,   B4: 0.010000
	// id: 2, input: %!d(string=1E-3),      A: 0.001000,   B: 0.000000,   B0: 0.000000,   B1: 0.000000,   B2: 0.000000,   B3: 0.001000,   B4: 0.001000
	// id: 3, input: %!d(string=1E-4),      A: 0.000100,   B: 0.000000,   B0: 0.000000,   B1: 0.000000,   B2: 0.000000,   B3: 0.000000,   B4: 0.000100
	// id: 4, input: %!d(string=1E-5),      A: 0.000010,   B: 0.000000,   B0: 0.000000,   B1: 0.000000,   B2: 0.000000,   B3: 0.000000,   B4: 0.000000
	// id: 5, input: %!d(string=5E-5),      A: 0.000050,   B: 0.000000,   B0: 0.000000,   B1: 0.000000,   B2: 0.000000,   B3: 0.000000,   B4: 0.000100
	// id: 6, input: %!d(string=0.0006),    A: 0.000600,   B: 0.000000,   B0: 0.000000,   B1: 0.000000,   B2: 0.000000,   B3: 0.001000,   B4: 0.000600
	// id: 7, input: %!d(string=0.0004),    A: 0.000400,   B: 0.000000,   B0: 0.000000,   B1: 0.000000,   B2: 0.000000,   B3: 0.000000,   B4: 0.000400
	// id: 8, input: %!d(string=0.006),     A: 0.006000,   B: 0.000000,   B0: 0.000000,   B1: 0.000000,   B2: 0.010000,   B3: 0.006000,   B4: 0.006000
	// id: 9, input: %!d(string=0.004),     A: 0.004000,   B: 0.000000,   B0: 0.000000,   B1: 0.000000,   B2: 0.000000,   B3: 0.004000,   B4: 0.004000
	// id: 10, input: %!d(string=1.0),      A: 1.000000,   B: 1.000000,   B0: 1.000000,   B1: 1.000000,   B2: 1.000000,   B3: 1.000000,   B4: 1.000000
	// id: 11, input: %!d(string=281.12),   A: 281.120000, B: 281.000000, B0: 281.000000, B1: 281.100000, B2: 281.120000, B3: 281.120000, B4: 281.120000
	// id: 12, input: %!d(string=0.938),    A: 0.938000,   B: 1.000000,   B0: 1.000000,   B1: 0.900000,   B2: 0.940000,   B3: 0.938000,   B4: 0.938000
	// id: 13, input: %!d(string=15.315),   A: 15.315000,  B: 15.000000,  B0: 15.000000,  B1: 15.300000,  B2: 15.320000,  B3: 15.315000,  B4: 15.315000
	// id: 14, input: %!d(string=1.234567), A: 1.234567,   B: 1.000000,   B0: 1.000000,   B1: 1.200000,   B2: 1.230000,   B3: 1.235000,   B4: 1.234600

}
