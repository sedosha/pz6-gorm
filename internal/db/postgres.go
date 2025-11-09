package db

import (
    "log"
    "os"
    "time"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
    dsn := os.Getenv("DB_DSN")
    if dsn == "" {
        log.Fatal("DB_DSN environment variable is not set")
    }

    // ОТЛАДОЧНАЯ ИНФОРМАЦИЯ
    log.Printf("DEBUG: DSN value: %s", dsn)
    
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        log.Printf("DEBUG: Connection error: %v", err)
        log.Fatal("connect db:", err)
    }

    sqlDB, err := db.DB()
    if err != nil {
        log.Fatal(err)
    }
    sqlDB.SetMaxOpenConns(25)
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetConnMaxLifetime(2 * time.Hour)

    log.Println("Successfully connected to database")
    return db
}
