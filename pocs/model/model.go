package model

// FileRecord 结构体表示数据库中的表结构
type FileRecord struct {
	ID      uint   `gorm:"primaryKey"`        // ID 为主键
	Name    string `gorm:"type:varchar(255)"` // 文件名
	From    string `gorm:"type:varchar(255)"` // 来源目录
	Content string `gorm:"type:text"`         // 文件内容
}
