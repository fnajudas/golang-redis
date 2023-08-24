package setredis

import (
	"errors"
	dtredis "golangredis/models/dtRedis"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type setMysql struct {
	db *gorm.DB
}

func NewSetMysql(db *gorm.DB) *setMysql {
	return &setMysql{
		db: db,
	}
}

func (s *setMysql) SaveMysql(req dtredis.DataSet) error {
	query := `
		INSERT INTO
			rediszs
				(value, created_on)
		VALUES
			(?, ?)
	`

	trx := s.db.Begin()
	defer trx.Rollback()

	if err := trx.Exec(query, req.Value, time.Now()).Error; err != nil {
		log.Println("Error insert values redis to database")
		return errors.New("Error insert redis to database")
	}

	if err := trx.Commit().Error; err != nil {
		log.Println("Error commit")
		return errors.New("Error commit")
	}

	return nil
}
