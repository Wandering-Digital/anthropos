package conn

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Wandering-Digital/anthropos/internal/config"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/lib/pq"
)

var db *DB

type DB struct {
	RawSQL *sql.DB
	GormDB *gorm.DB
}

func ConnectDB() error {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		config.DB().Host,
		config.DB().Username,
		config.DB().Password,
		config.DB().Name,
		config.DB().Port)

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(config.DB().MaxIdleConn)
	sqlDB.SetMaxOpenConns(config.DB().MaxOpenConn)
	sqlDB.SetConnMaxLifetime(config.DB().MaxConnLifetime)

	if err := sqlDB.Ping(); err != nil {
		return err
	}

	newLogger := logger.New(
		logrus.New(),
		logger.Config{
			SlowThreshold:             time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,      // Log level
			IgnoreRecordNotFoundError: true,             // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,             // Disable color
		},
	)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger, CreateBatchSize: config.DB().BatchSize})

	if err != nil {
		return err
	}

	newDB, _ := gormDB.DB()

	newDB.SetMaxIdleConns(config.DB().MaxIdleConn)
	newDB.SetMaxOpenConns(config.DB().MaxOpenConn)
	newDB.SetConnMaxLifetime(config.DB().MaxConnLifetime)

	db = &DB{
		RawSQL: sqlDB,
		GormDB: gormDB,
	}

	log.Printf("Database Connection Successful\nConnMaxLifetime: %v\nMaxIdleConn: %v\nMaxOpenConn: %v\n",
		config.DB().MaxConnLifetime,
		config.DB().MaxIdleConn,
		config.DB().MaxOpenConn)

	return nil
}

func GetDB() *DB {
	return db
}
