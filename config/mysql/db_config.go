package mysql

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"golangredis/models/migrate"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/subosito/gotenv"
)

type Db struct {
	Database *gorm.DB
}

func AutoMigrate(db *gorm.DB) {
	if db == nil {
		log.Println("Database connection is nil.")
		return
	}

	model := &migrate.Redisz{}
	if !db.HasTable(model) {
		if err := db.AutoMigrate(model); err != nil {
			log.Println("Error auto migrating:", err)
			return
		}
		log.Println("Successfully migrated the database.")
	} else {
		log.Println("Table already exists.")
	}
}

func Initialize(dbName string) *gorm.DB {
	gotenv.Load()

	dbUsername := os.Getenv("DB_USER")
	dbPassowrd := os.Getenv("DB_PASS")
	dbIp := os.Getenv("DB_IP")
	dbType := os.Getenv("DB_TYPE")

	connSetting := "charset=utf8mb4&parseTime=true&checkConnLiveness=true"
	connString := fmt.Sprintf("%s:%s@(%s)/%s?%s", dbUsername, dbPassowrd, dbIp, dbName, connSetting)

	db, err := gorm.Open(dbType, connString)
	if err != nil {
		log.Printf("Cant open database - %v", err)
		return nil
	}
	log.Println("Open database successfull..")

	db.DB().SetMaxOpenConns(25)
	db.DB().SetMaxIdleConns(25)
	db.DB().SetConnMaxLifetime(5 * time.Minute)

	active, err := strconv.ParseBool(os.Getenv("DB_DEBUG"))
	if err != nil {
		active = false
	}
	db.LogMode(active)

	err = db.Exec("SET SESSION sql_mode=(SELECT REPLACE(@@sql_mode,'ONLY_FULL_GROUP_BY',''));").Error
	if err != nil {
		log.Panic(err)
	}

	AutoMigrate(db)
	return db
}

// Connection Database ...
func (m *Db) DatabaseConnection() {
	if os.Getenv("DB_NAME") != "" {
		m.Database = Initialize(os.Getenv("DB_NAME"))
		return
	}
	m.Database = Initialize("learning")
}

func URLRewriter(router *mux.Router, baseURLPath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, baseURLPath)
		router.ServeHTTP(w, r)
	}
}
