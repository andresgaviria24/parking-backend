package persistence

import (
	"os"
	"ws_parking/infrastructure/persistence"
	"ws_parking/infrastructure/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DbHelper struct {
	ParkingRepository repository.ParkingRepository
	db                *gorm.DB
}

func InitDbHelper() (*DbHelper, error) {

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.AutoMigrate()
	sqlDB := db.DB()
	sqlDB.SetMaxIdleConns(600)
	sqlDB.SetMaxOpenConns(0)
	return &DbHelper{
		ParkingRepository: persistence.InitParkingRepositoryImpl(db),
		db:                db,
	}, nil
}

//closes the  database connection
func (s *DbHelper) Close() error {
	return s.db.Close()
}
