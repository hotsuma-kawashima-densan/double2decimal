package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Years struct {
	ID                   int       `gorm:"primaryKey"`
	JapaneseYear         string    `gorm:"type:varchar(32);column:japanese_year"`
	JapaneseYearWithRuby string    `gorm:"type:varchar(1024);column:japanese_year_with_ruby"`
	RegisteredAt         time.Time `gorm:"type:timestamp;column:registered_at"`
	UpdatedAt            time.Time `gorm:"type:timestamp;column:updated_at"`
}

func main() {
	dsn := "backend_api:backend_api@tcp(127.0.0.1:13306)/myassess?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("接続できない:", err)
		return
	}

	fmt.Println("接続できます。")

	var years []Years
	result := db.Find(&years)
	if result.Error != nil {
		fmt.Println("エラー:", result.Error)
		return
	}

	fmt.Println("取得したデータ:")
	for _, year := range years {
		fmt.Printf("ID: %d, 和暦: %s, ルビ付き和暦: %s, 登録日時: %s, 更新日時: %s\n",
			year.ID, year.JapaneseYear, year.JapaneseYearWithRuby,
			year.RegisteredAt.Format(time.RFC3339), year.UpdatedAt.Format(time.RFC3339))
	}
}
