package configuration

import (
	"log"
	"os"
	"time"

	"github.com/donnyirianto/go-be-fiber/exception"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(config Config) *gorm.DB {
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
	exception.PanicLogging(err)

	sqlDB, err := db.DB()
	exception.PanicLogging(err)

	sqlDB.SetMaxOpenConns(maxPoolOpen)
	sqlDB.SetMaxIdleConns(maxPoolIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(maxPollLifeTime) * time.Second)

	// Uncomment and customize the auto-migrate part if needed
	// err = db.AutoMigrate(&entity.Product{})
	// err = db.AutoMigrate(&entity.Transaction{})
	// err = db.AutoMigrate(&entity.TransactionDetail{})
	// err = db.AutoMigrate(&entity.User{})
	// err = db.AutoMigrate(&entity.UserRole{})
	// exception.PanicLogging(err)

	return db
}
