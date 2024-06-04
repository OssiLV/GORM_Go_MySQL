package configs

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)

func InitDb(configs Config) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       configs.DB.ConnectionString, // data source name
		DefaultStringSize:         256,                                                                        // default size for string fields
		DisableDatetimePrecision:  true,                                                                       // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                       // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                       // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                      // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to MySQL DB")
	}

	log.Printf("Connect to MySQL DB success!")
	return db
}