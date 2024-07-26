package server

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newRDB(cfg *Config) (*gorm.DB, error) {
	if strings.ToLower(cfg.Database.Kind) == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
			cfg.Database.User,
			cfg.Database.Secret,
			cfg.Database.Host,
			cfg.Database.Schema,
		)
		if db, err := gorm.Open(mysql.New(mysql.Config{
			DSN:                       dsn,   // data source name
			DefaultStringSize:         256,   // default size for string fields
			DisableDatetimePrecision:  false, // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex:    false, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn:   false, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
		}), &gorm.Config{}); err != nil {
			return nil, err
		} else {
			db.Set("gorm:table_options", "ENGINE=InnoDB")
			return db, nil
		}
	}
	return nil, errors.New("not supported database")
}
