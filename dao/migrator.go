package dao

import "fmt"

func migrator() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate()
	if err != nil {
		fmt.Println("err:", err)
	}
	return
}
