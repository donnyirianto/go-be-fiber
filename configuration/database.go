package configuration

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/donnyirianto/go-be-fiber/exception"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(config Config) (*gorm.DB, error) {
	// Validate required configuration parameters
	requiredParams := []string{"DATASOURCE_USERNAME", "DATASOURCE_PASSWORD", "DATASOURCE_HOST", "DATASOURCE_PORT", "DATASOURCE_DB_NAME"}
	for _, param := range requiredParams {
		if config.GetString(param) == "" {
			err := errors.New("Missing required configuration parameter: " + param)
			exception.PanicLogging(err)
			return nil, err
		}
	}

	username := config.GetString("DATASOURCE_USERNAME")
	password := config.GetString("DATASOURCE_PASSWORD")
	host := config.GetString("DATASOURCE_HOST")
	port := config.GetString("DATASOURCE_PORT")
	dbName := config.GetString("DATASOURCE_DB_NAME")
	maxPoolOpen := config.GetInt("DATASOURCE_POOL_MAX_CONN")
	maxPoolIdle := config.GetInt("DATASOURCE_POOL_IDLE_CONN")
	maxPollLifeTime := config.GetInt("DATASOURCE_POOL_LIFE_TIME")

	loggerDb := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: loggerDb,
	})

	exception.PanicLogging(err) // Use your exception.PanicLogging function here

	sqlDB, err := db.DB()
	exception.PanicLogging(err)

	sqlDB.SetMaxOpenConns(maxPoolOpen)
	sqlDB.SetMaxIdleConns(maxPoolIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(maxPollLifeTime) * time.Second)

	return db, nil
}
