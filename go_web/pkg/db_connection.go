package dao

import (
	"git_truongvudinh/go_web/internal/domain/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(dsn string) (*gorm.DB, error) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal("failed to connect database")
	}

	err = db.AutoMigrate(&entity.User{}, &entity.Project{})
	if err != nil {
		logrus.Fatal("failed to migrate database:", err)
	}

	logrus.Info("Database migrated successfully")
	return db, err
}
