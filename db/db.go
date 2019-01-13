package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/reivaj05/GoLogger"
)

type dbClientInterface interface {
	First(interface{}, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Close() error
}

type dbObj struct {
	Client dbClientInterface
}

var DB *dbObj

func StartDB(models ...interface{}) {
	client, err := gorm.Open("postgres", getDBOptions())
	if err != nil {
		GoLogger.LogPanic("Error connecting to DB", map[string]interface{}{"error": err.Error()})
	}
	DB = &dbObj{Client: client}
	client.AutoMigrate(models...)
}

func getDBOptions() string {
	data := "host=%s port=%s user=%s password=%s dbname=%s"
	host := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	sslMode := "disable"
	connectionQuery := fmt.Sprintf(data, host, port, user, password, name)
	if sslMode != "" {
		connectionQuery += fmt.Sprintf(" sslmode=%s", sslMode)
	}
	return connectionQuery
}

func (d *dbObj) Close() {
	if d.Client != nil {
		d.Client.Close()
	}
}
