package configs

import (
	
	"reflect"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	color "name/enums"
	"name/extensions"
	"name/models"

)

func InitDb(configs Config, isDropTables bool) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       configs.DB.ConnectionString, // data source name
		DefaultStringSize:         256,                         // default size for string fields
		DisableDatetimePrecision:  true,                        // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                        // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                        // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                       // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		extensions.Logger(color.Green, "Cannot connect to MySQL DB!")
	}

	extensions.Logger(color.Green, "Connect to MySQL DB success!")

	if isDropTables {
		// Drop Tables
		dropTable(*db, &models.User{})
		dropTable(*db, &models.Todo{})
	} else {
		// Init Tables
		createTable(*db, &models.User{})
		createTable(*db, &models.Todo{})
	}
	
	return db
}

func createTable(db gorm.DB, table interface{})  {
	if exist := db.Migrator().HasTable(table); exist {
		if error := db.Migrator().DropTable(table); error != nil {
			// log.Printf("Cannot drop [%v] table", reflect.TypeOf(table).Elem().Name())
			extensions.Logger(color.Green, "Cannot drop [%v] table", reflect.TypeOf(table).Elem().Name())
		}
		if error := db.Migrator().CreateTable(table); error != nil {
			extensions.Logger(color.Green, "Cannot create [%v] table", reflect.TypeOf(table).Elem().Name())
		}
	} else {
		if error := db.Migrator().CreateTable(table); error != nil {
			extensions.Logger(color.Green, "Cannot create [%v] table", reflect.TypeOf(table).Elem().Name())
		}
	}
	extensions.Logger(color.Green, "Create table [%v] Success!", reflect.TypeOf(table).Elem().Name())
}

func dropTable(db gorm.DB, table interface{})  {
	if error := db.Migrator().DropTable(table); error != nil {
		extensions.Logger(color.Green, "Cannot drop [%v] table", reflect.TypeOf(table).Elem().Name())
	}
	extensions.Logger(color.Green, "Drop table [%v] Success!", reflect.TypeOf(table).Elem().Name())
}