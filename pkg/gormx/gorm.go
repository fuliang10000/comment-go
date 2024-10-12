package gormx

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	ormLogger "gorm.io/gorm/logger"
)

func Init(c *Config) (*gorm.DB, error) {
	fmt.Println(c.MysqlDNS)
	var gormConf = &gorm.Config{
		Logger: ormLogger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), ormLogger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  ormLogger.Warn,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		}),
	}
	if c.IsDebug {
		gormConf.Logger.LogMode(ormLogger.Info)
	}

	if db, err := gorm.Open(mysql.Open(c.MysqlDNS), gormConf); err != nil {
		return nil, err
	} else {
		sqlDB, err := db.DB()
		if err != nil {
			return nil, err
		}
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(10)
		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(1000)
		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(time.Hour)

		return db, nil
	}
}

func IsNotfoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
