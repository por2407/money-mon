package infrastructure

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDB() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// --- Configure Connection Pool ---
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns ตั้งค่าจำนวนการเชื่อมต่อสูงสุดใน idle pool
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns ตั้งค่าจำนวนการเชื่อมต่อสูงสุดที่เปิดไปยังฐานข้อมูล
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime ตั้งค่าระยะเวลาสูงสุดที่สามารถนำการเชื่อมต่อกลับมาใช้ใหม่ได้
	sqlDB.SetConnMaxLifetime(time.Hour)

	// โชว์ stats ครั้งแรกที่เชื่อมต่อ
	go func() {
		ticker := time.NewTicker(3 * time.Minute)
		for range ticker.C {
			stats := sqlDB.Stats()
			fmt.Printf("[DB STATS] Open: %d, InUse: %d, Idle: %d, WaitCount: %d, WaitDuration: %v\n",
				stats.OpenConnections,
				stats.InUse,
				stats.Idle,
				stats.WaitCount,
				stats.WaitDuration,
			)
		}
	}()

	return db, nil
}
