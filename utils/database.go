package utils

import (
	"fmt"
	"javan-inventory-barang/migrations"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// OpenPostgres opens a PostgreSQL connection with GORM and runs AutoMigrate for registered models.
func OpenPostgres() (*gorm.DB, error) {
	config := gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   viper.GetString("DB_TABLE_PREFIX"),
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	db, err := gorm.Open(postgres.Open(PostgresDSN()), &config)
	if err != nil {
		return nil, fmt.Errorf("gorm open: %w", err)
	}

	if viper.GetBool("DB_AUTO_MIGRATE") {
		if err := db.AutoMigrate(migrations.ModelMigrations...); err != nil {
			return nil, fmt.Errorf("auto migrate: %w", err)
		}
	}

	if nil != db {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(1)
		sqlDB.SetConnMaxLifetime(time.Second * 5)
	}

	return db, nil
}

// PostgresDSN builds a libpq connection string from viper (DB_* keys).
func PostgresDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
	)
}
