package main

import (
	"fmt"

	"github.com/shopspring/decimal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Result struct {
	ID int32           `gorm:"column:id;not null" json:"id"`
	B4 decimal.Decimal `gorm:"column:B4;not null" json:"B4"`
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:13306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("接続失敗:", err)
		return
	}

	// バインド変数
	bindParams := map[string]any{}

	sql := ""
	sql += " SELECT"
	sql += "     ID,"
	sql += "     B4"
	sql += " FROM temp"
	sql += " WHERE"
	sql += "     B4 = @B4"
	sql += " ;"

	bindParams["B4"] = decimal.NewFromFloat(0.2625)

	var result []Result
	if err := db.Debug().Raw(sql, bindParams).Scan(&result).Error; err != nil {
		fmt.Println("データ取得エラー:", err)
		return
	}
	// 2024/08/23 16:11:35 /Users/hotsumakawashima/task/20240822_decimal/DSK_PJ/main.go:40
	// [1.635ms] [rows:1]  SELECT     ID,     B4 FROM temp WHERE     B4 = '0.2625' ;

	for _, data := range result {
		fmt.Printf("id: %d, B4: %s\n", data.ID, data.B4)
		// id: 15, B4: 0.2625
	}
}
