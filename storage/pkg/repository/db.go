package repository

import (
	"fmt"
	"github.com/Yurovskikh/news/storage/config"

	"github.com/Yurovskikh/news/storage/pkg/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"time"
)

func NewDB(cfg *config.Config) *gorm.DB {
	dls := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDBName,
	)
	db, err := gorm.Open("postgres", dls)
	if err != nil {
		log.Fatal(err)
	}
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Minute * 10)
	err = db.DB().Ping()
	if err != nil {
		log.Fatal(err)
	}

	for _, domain := range domains {
		err := db.AutoMigrate(domain).Error
		if err != nil {
			log.Fatal(err)
		}
	}
	return db
}

var domains = []interface{}{
	new(model.News),
}
