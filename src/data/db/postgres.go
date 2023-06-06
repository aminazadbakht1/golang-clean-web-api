package db

import (
	"fmt"
	"time"

	"github.com/aminazadbakht1/golang-clean-web-api/config"
	"github.com/aminazadbakht1/golang-clean-web-api/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

var dbClient *gorm.DB

func InitDb(cfg *config.Config) error {
	var err error
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password,
		cfg.Postgres.DbName, cfg.Postgres.SSLMode)

	dbClient, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})

	if err != nil {
		return err
	}

	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()

	if err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	logger.Info(logging.Postgres, logging.Startup, "Db connection established", nil)
	return nil
}

func GetDb() *gorm.DB{
	return dbClient
}

func CloseDb(){
	cnn, _ := dbClient.DB()
	cnn.Close()
}