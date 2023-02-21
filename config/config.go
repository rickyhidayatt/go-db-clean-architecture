package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	Db *sqlx.DB
}

func (c *Config) inittDb() {
	// ============= panggil file env ===============
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbDriver := os.Getenv("DB_DRIVER")

	// ============= panggil file env ===============

	//konek  ke DB
	connectDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	//Buka koneksi
	db, err := sqlx.Open(dbDriver, connectDB)
	// defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	if err := db.Ping(); err != nil {
		fmt.Println(err.Error())
	}

	c.Db = db
}

func (c *Config) DbConnect() *sqlx.DB { // method untuk mendapatkan koneksi
	return c.Db
}

func NewConfig() Config { // biar bisa di panggil initDB dke publik
	config := Config{}
	config.inittDb()
	return config // mengirimkan Objek config yang didalamnya terdapat attribut dari func initDb
}
