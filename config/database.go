package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		AppConfig.DBHost,
		AppConfig.DBUser,
		AppConfig.DBPassword,
		AppConfig.DBName,
		AppConfig.DBPort,
		AppConfig.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
	return nil
}

func RunMigrations() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database: %v", err)
	}

	_, err = sqlDB.Exec(`CREATE TABLE IF NOT EXISTS migrations (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL UNIQUE,
		applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %v", err)
	}

	rows, err := sqlDB.Query("SELECT name FROM migrations")
	if err != nil {
		return fmt.Errorf("failed to query migrations: %v", err)
	}
	defer rows.Close()

	appliedMigrations := make(map[string]bool)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return fmt.Errorf("failed to scan migration: %v", err)
		}
		appliedMigrations[name] = true
	}

	migrationFiles, err := filepath.Glob("db/migrations/*_*.up.sql")
	if err != nil {
		return fmt.Errorf("failed to read migrations: %v", err)
	}

	sort.Strings(migrationFiles)

	for _, file := range migrationFiles {
		baseName := filepath.Base(file)
		if appliedMigrations[baseName] {
			log.Printf("Migration %s already applied", baseName)
			continue
		}

		content, err := ioutil.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration %s: %v", file, err)
		}

		log.Printf("Applying migration %s", baseName)
		if _, err := sqlDB.Exec(string(content)); err != nil {
			return fmt.Errorf("failed to apply migration %s: %v", file, err)
		}

		if _, err := sqlDB.Exec("INSERT INTO migrations (name) VALUES ($1)", baseName); err != nil {
			return fmt.Errorf("failed to record migration %s: %v", file, err)
		}
	}

	return nil
}

func SeedDatabase() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database: %v", err)
	}

	seedOrder := []string{
		"products.sql",
		"orders.sql",
		"discounts.sql",
	}

	for _, seedName := range seedOrder {
		filePath := filepath.Join("db/seeds", seedName)

		if _, err := filepath.Glob(filePath); err != nil {
			return fmt.Errorf("seed file not found: %s", filePath)
		}

		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read seed %s: %v", filePath, err)
		}

		log.Printf("Applying seed %s", seedName)
		if _, err := sqlDB.Exec(string(content)); err != nil {
			return fmt.Errorf("failed to apply seed %s: %v", filePath, err)
		}
	}

	return nil
}
