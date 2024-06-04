package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	envAppEnums "name/enums"
	color "name/enums"
	"name/extensions"
)

type DBConfig struct {
	UserName			string
	Password			string
	DBName				string
	ConnectionString 	string
}

type Config struct {
	DB		DBConfig
	EnvApp	string
}

func New() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		extensions.Logger(color.Red, "Some error occured. Err: %s", err)
	}

	dbUserName 			:= getEnvDBUserName("DB_USER_NAME", "root")
	dbPassword 			:= getEnvDBPassword("DB_PASSWORD", "OssiLV711")
	dbName 				:= getEnvDBName("DB_NAME", "")
	dbConnectionString 	:= getEnvDBConnectionString(
		"DB_CONNECTION_STRING",
		dbUserName,
		dbPassword,
		dbName,
	)

	return &Config{
		DB: DBConfig{
			UserName: dbUserName,
			Password: dbPassword,
			DBName: dbName,
			ConnectionString: dbConnectionString,
		},
		EnvApp: getEnvApp("APP_ENV", envAppEnums.Develope),
	}
}

func getEnvApp(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	
	return defaultValue
}
func getEnvDBUserName(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	
	return defaultValue
}
func getEnvDBPassword(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	
	return defaultValue
}
func getEnvDBName(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	
	return defaultValue
}
func getEnvDBConnectionString(key, dbUserName, dbPassword, dbName string ) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fmt.Sprintf(
		"%v:%v@tcp(127.0.0.1:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		dbUserName,
		dbPassword,
		dbName,
	)
}