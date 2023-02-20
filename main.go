package main

import "github.com/rickyhidayatt/delivery"

func main() {
	delivery.Run()

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// // get environment variables
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// dbDriver := os.Getenv("DB_DRIVER")

	// fmt.Println("DB_HOST:", dbHost)
	// fmt.Println("DB_PORT:", dbPort)
	// fmt.Println("DB_USER:", dbUser)
	// fmt.Println("DB_PASSWORD:", dbPassword)
	// fmt.Println("DB_NAME:", dbName)
	// fmt.Println("DB_DRIVER:", dbDriver)

}
