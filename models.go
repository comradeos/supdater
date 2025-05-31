package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

//////////////////////////////////////////////////////////////////////////////////////

type File struct {
	gorm.Model
	Name          string `gorm:"column:name"`           // ім’я
	Size          int64  `gorm:"column:size"`           // розмір
	SHA256        string `gorm:"column:sha256"`         // контрольна сума
	UnpackedSize  int64  `gorm:"column:unpacked_size"`  // розпакований розмір
	InstallerPath string `gorm:"column:installer_path"` // шлях до EXE інсталятора
	SpecialParams string `gorm:"column:special_params"` // спеціальні параметри
}

func GetFileByName(db *gorm.DB, name string) (*File, error) {
	var file File
	if err := db.Where("name = ?", name).First(&file).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

//////////////////////////////////////////////////////////////////////////////////////

type Token struct {
	gorm.Model
	Token string `gorm:"uniqueIndex"` // токен должен быть уникальным
}

//////////////////////////////////////////////////////////////////////////////////////

func InitialiseDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе:", err)
	}

	if err := db.AutoMigrate(&File{}, &Token{}); err != nil {
		log.Fatal("Ошибка миграции:", err)
	}

	log.Println("База и таблица готовы.")

	return db
}
