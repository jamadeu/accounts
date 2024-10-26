package config

import (
	"github.com/jamadeu/accounts/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() (*gorm.DB, error) {
	// logger := GetLogger("InitializeDb")
	// Connect DB
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	// fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
	// os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// fmt.Errorf("Error to connect database: %v", err)
		return nil, err
	}

	// Migrate the schema
	if err = db.AutoMigrate(&schemas.User{}, &schemas.Transaction{}); err != nil {
		// fmt.Errorf("Automigratoin error: %v", err)
		return nil, err
	}
	if err = db.AutoMigrate(&schemas.Account{}); err != nil {
		// fmt.Errorf("Automigratoin error: %v", err)
		return nil, err
	}
	return db, nil
}
