package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	connectDatabase()
	//err := DB.AutoMigrate(&Link{}) // TODO: add table structs here
	//if err != nil {
	//	logrus.Fatal(err)
	//}
	//err1 := DB.Migrator().CreateTable(&User{})
	//if err1 != nil {
	//	logrus.Fatal(err)
	//}
}

func connectDatabase() {
	viper.SetConfigName("conf")
	viper.AddConfigPath("./model")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic("viper read error")
	}

	loginInfo := viper.GetStringMapString("key")

	dbArgs := loginInfo["username"] + ":" + loginInfo["password"] +
		"@tcp(localhost:3306)/" + loginInfo["db_name"] + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dbArgs)
	var err error
	DB, err = gorm.Open(mysql.Open(dbArgs), &gorm.Config{})
	if err != nil {
		logrus.Panic(err)
	}
}
