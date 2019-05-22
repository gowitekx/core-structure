package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/infinity-framework/backend/configs"
)

//DBMigrate -used for database migration
func DBMigrate() {
	db, error := sql.Open("mysql", configs.Config.Username+":"+configs.Config.Password+"@/"+configs.Config.DatabaseName+"?multiStatements=true")
	if error != nil {
		configs.Ld.Logger(context.Background(), configs.ERROR, "Migration Error:Failed to connect database!:", error)
	}

	driver, error := mysql.WithInstance(db, &mysql.Config{})
	if error != nil {
		configs.Ld.Logger(context.Background(), configs.ERROR, "Migration Error:", error)
	}

	m, error := migrate.NewWithDatabaseInstance(
		"file:/home/anant/go/src/github.com/infinity-framework/backend/database/migrations",
		"mysql",
		driver,
	)
	if error != nil {
		configs.Ld.Logger(context.Background(), configs.ERROR, "Migration Error:", error)
	}
	error1 := m.Up()
	if error1 != nil {
		configs.Ld.Logger(context.Background(), configs.WARN, "Migration Error:", error1)
		log.Println("Migration Status:", error1)
	}
	log.Println("Setup Done!")
}
